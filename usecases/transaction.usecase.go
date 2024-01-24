package usecase

import (
	"cc-transaction/constants"
	cModels "cc-transaction/controllers/models"
	dbModels "cc-transaction/databases/postgresql/models"
	"cc-transaction/hosts/callback/models"
	"cc-transaction/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func (uc *usecase)TransItem(req models.DecTransactionItems,headers cModels.ReqHeader)(string,error){
	result:=models.ResponseTransactionItems{}

	decryptedReq,err:=utils.DecryptTransItem(req)
	if err!=nil{
		return result.Data, err
	}
	
	itemID,err:=strconv.Atoi(decryptedReq.ItemID)
	if err!=nil{
		return result.Data, err
	}
	prices,err:=strconv.Atoi(decryptedReq.Price)
	if err!=nil{
		return result.Data, err
	}
	qtys,err:=strconv.Atoi(decryptedReq.Quantity)
	if err!=nil{
		return result.Data, err
	}
	percentages,_:=strconv.Atoi(decryptedReq.Percentage)
	if err!=nil{
		return result.Data, err
	}

	cc,err:=uc.postgre.GetCC(decryptedReq)
	if err!=nil{
		return result.Data,err
	}

	if decryptedReq.CVV!=cc.CVV{
		return result.Data,errors.New("Invalid Credit Cards")
	}

	discount:=float64(percentages)/100
	price:=float64(prices)
	qty:=float64(qtys)
	reqTotalPrice:=int((price-(price*discount))*qty)

	if reqTotalPrice>cc.Balance{
		return result.Data,errors.New("Insuficient Balance")
	}

	cc.Balance -=reqTotalPrice
	
	err=uc.postgre.DeductCC(cc)
	if err!=nil{
		return result.Data,err
	}

	reqDB:=dbModels.Order{}
	reqDB.ItemID=itemID
	reqDB.Name=decryptedReq.Name
	reqDB.Type=decryptedReq.Type
	reqDB.Price=prices
	reqDB.TotalPrice=reqTotalPrice
	reqDB.Quantity=qtys
	reqDB.CC=decryptedReq.CCNumber
	reqDB.Discount=decryptedReq.Discount
	reqDB.Status=constants.STATUS_PENDING

	resDB,err:=uc.postgre.OrderTransItem(reqDB)
	if err!=nil{
		cc.Balance +=reqTotalPrice
		err:=uc.postgre.DeductCC(cc)
		if err!=nil{
			return result.Data,err
		}
		return result.Data,err
	}
	timeStamp:=time.Now().Format("15:04:05")
	
	reqHost:=models.ReqCallbackItems{}
	reqHost.Amount=strconv.Itoa(reqDB.TotalPrice)
	reqHost.CCNumber=decryptedReq.CCNumber
	reqHost.Discount=decryptedReq.Discount
	reqHost.ItemID=decryptedReq.ItemID
	reqHost.Quantity=decryptedReq.Quantity

	// fmt.Println("reqHost",reqHost)
	encryptedReqHost,err:=utils.EncryptTransItem(reqHost)
	if err!=nil{
		return result.Data, err
	}
	// fmt.Println("encryptedReqHost:",encryptedReqHost)
	bytes,err:=json.Marshal(encryptedReqHost)
	if err!=nil{
		return result.Data, err
	}
	signature:=utils.Signature(string(bytes),timeStamp)
	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Content-Type", "application/json")
	header.Add("TimeStamp", timeStamp)
	header.Add("Signature", signature)
	_,bytes,err=uc.host.Callback().Send(constants.TRANSACTION_ITEMS,encryptedReqHost,header)
	if err!=nil{
		err:=uc.RollbackTrans(cc,resDB,reqTotalPrice)
		if err!=nil{
			return result.Data,err
		}
		return result.Data, errors.New(constants.ERROR_HOST)
	}
	err=json.Unmarshal(bytes,&result)
	if err!=nil{
		err:=uc.RollbackTrans(cc,resDB,reqTotalPrice)
		if err!=nil{
			return result.Data,err
		}
		return result.Data, errors.New(constants.ERROR_HOST)
	}
	if result.Code!=200{
		err:=uc.RollbackTrans(cc,resDB,reqTotalPrice)
		if err!=nil{
			return result.Data,err
		}
		return result.Data,errors.New(result.Message)
	}
	resDB.Status=constants.STATUS_SUCCESS
	err=uc.postgre.UpdateTransItem(resDB)
	if err!=nil{
		return result.Data,err
	}
	return result.Data,nil
}

func (uc *usecase)RollbackTrans(cc dbModels.CreditCards, resDB dbModels.Order, reqTotalPrice int)error{
	// fmt.Println("rollback called")
	cc.Balance +=reqTotalPrice
	err:=uc.postgre.DeductCC(cc)
	if err!=nil{
		return err
	}
	resDB.Status=constants.STATUS_FAILED
	err=uc.postgre.UpdateTransItem(resDB)
	if err!=nil{
		return err
	}
	return nil
}