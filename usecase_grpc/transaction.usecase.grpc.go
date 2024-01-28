package usecase_grpc

import (
	"cc-transaction/constants"
	dbModels "cc-transaction/databases/postgresql/models"
	"cc-transaction/hosts/callback/models"
	"cc-transaction/protogen/merchant"
	"cc-transaction/utils"
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"log"

	"google.golang.org/grpc/metadata"
)

func (uc *usecaseGrpc)TransItems(req *merchant.ReqTransItemsModel)(string, error){
	result:=""

	decryptedReq,err:=utils.DecryptTransItemGrpc(req)
	if err!=nil{
		return result, err
	}
	
	itemID,err:=strconv.Atoi(decryptedReq.ItemID)
	if err!=nil{
		return result, err
	}
	prices,err:=strconv.Atoi(decryptedReq.Price)
	if err!=nil{
		return result, err
	}
	qtys,err:=strconv.Atoi(decryptedReq.Quantity)
	if err!=nil{
		return result, err
	}
	percentages,_:=strconv.Atoi(decryptedReq.Percentage)
	if err!=nil{
		return result, err
	}

	cc,err:=uc.postgre.GetCC(decryptedReq)
	if err!=nil{
		return result,err
	}

	if decryptedReq.CVV!=cc.CVV{
		return result,errors.New("Invalid Credit Cards")
	}

	discount:=float64(percentages)/100
	price:=float64(prices)
	qty:=float64(qtys)
	reqTotalPrice:=int((price-(price*discount))*qty)

	if reqTotalPrice>cc.Balance{
		return result,errors.New("Insuficient Balance")
	}

	cc.Balance -=reqTotalPrice
	
	err=uc.postgre.DeductCC(cc)
	if err!=nil{
		return result,err
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
			return result,err
		}
		return result,err
	}
	timeStamp:=time.Now().Format("15:04:05")
	
	reqHost:=models.ReqCallbackItems{}
	reqHost.Amount=strconv.Itoa(reqDB.TotalPrice)
	reqHost.CCNumber=decryptedReq.CCNumber
	reqHost.Discount=decryptedReq.Discount
	reqHost.ItemID=decryptedReq.ItemID
	reqHost.Quantity=decryptedReq.Quantity

	encryptedReqHost,err:=utils.EncryptTransItemGrpc(reqHost)
	if err!=nil{
		return result, err
	}
	bytes,err:=json.Marshal(encryptedReqHost)
	if err!=nil{
		return result, err
	}
	signature:=utils.Signature(string(bytes),timeStamp)
	meta:=map[string]string{
		"timestamp":timeStamp,
		"signature":signature,
	}
	md := metadata.New(meta)
	ctx:=metadata.NewOutgoingContext(context.Background(),md)
	resHost,err:=uc.host.Merchant().TransItems(ctx,encryptedReqHost)
	log.Println("resHost:",resHost)

	if err!=nil{
		err:=uc.RollbackTrans(cc,resDB,reqTotalPrice)
		if err!=nil{
			return result,err
		}
		return result, errors.New(constants.ERROR_HOST)
	}

	if resHost.Code!=200{
		err:=uc.RollbackTrans(cc,resDB,reqTotalPrice)
		if err!=nil{
			return result,err
		}
		return result,errors.New(constants.STATUS_FAILED)
	}
	resDB.Status=constants.STATUS_SUCCESS
	err=uc.postgre.UpdateTransItem(resDB)
	if err!=nil{
		return result,err
	}

	return resHost.Data,nil
}

func (uc *usecaseGrpc)RollbackTrans(cc dbModels.CreditCards, resDB dbModels.Order, reqTotalPrice int)error{
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