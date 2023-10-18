package usecase

import (
	"cc-transaction/constants"
	dbModels "cc-transaction/databases/postgresql/models"
	"cc-transaction/hosts/callback/models"
	"cc-transaction/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (uc *usecase)TransItem(req models.TransactionItems)(models.ResponseItems,error){
	result:=models.ResponseTransactionItems{}

	req,err:=utils.DecryptTransItem(req)
	if err!=nil{
		return result.Data, err
	}
	
	itemID,_:=strconv.Atoi(req.ItemID)
	// discounts,_:=strconv.Atoi(req.Discount)
	prices,_:=strconv.Atoi(req.Price)
	qtys,_:=strconv.Atoi(req.Quantity)
	percentages,_:=strconv.Atoi(req.Percentage)

	cc,err:=uc.postgre.GetCC(req)
	if err!=nil{
		return result.Data,err
	}

	if req.CVV!=cc.CVV{
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
	reqDB.Name=req.Name
	reqDB.Type=req.Type
	reqDB.Price=prices
	reqDB.TotalPrice=reqTotalPrice
	reqDB.Quantity=qtys
	reqDB.CC=req.CCNumber
	reqDB.Discount=req.Discount
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

	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Content-Type", "application/json")
	req.Amount=strconv.Itoa(reqDB.TotalPrice)
	reqHost,err:=utils.EncryptTransItem(req)
	if err!=nil{
		return result.Data, err
	}
	res,bytes,err:=uc.host.Callback().Send(constants.TRANSACTION_ITEMS,reqHost,header)
	if err!=nil{
		err:=uc.RollbackTrans(cc,resDB,reqTotalPrice)
		if err!=nil{
			return result.Data,err
		}
		return result.Data, errors.New(constants.ERROR_HOST)
	}
	if res.StatusCode!=200{
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
		return result.Data,errors.New(constants.ERROR_HOST)
	}
	resDB.Status=constants.STATUS_SUCCESS
	err=uc.postgre.UpdateTransItem(resDB)
	if err!=nil{
		return result.Data,err
	}
	return result.Data,nil
}

func (uc *usecase)RollbackTrans(cc dbModels.CreditCards, resDB dbModels.Order, reqTotalPrice int)error{
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