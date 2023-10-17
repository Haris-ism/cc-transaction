package usecase

import (
	"cc-transaction/constants"
	dbModels "cc-transaction/databases/postgresql/models"
	"cc-transaction/hosts/callback/models"
	"encoding/json"
	"errors"
	"net/http"
)

func (uc *usecase)TransItem(req models.TransactionItems)(models.ResponseItems,error){
	result:=models.ResponseTransactionItems{}
	
	cc,err:=uc.postgre.GetCC(req)
	if err!=nil{
		return result.Data,err
	}

	if req.CVV!=cc.CVV{
		return result.Data,errors.New("Invalid Credit Cards")
	}

	discount:=float64(req.Percentage)/100
	price:=float64(req.Price)
	qty:=float64(req.Quantity)
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
	reqDB.ItemID=req.ItemID
	reqDB.Name=req.Name
	reqDB.Type=req.Type
	reqDB.Price=req.Price
	reqDB.TotalPrice=reqTotalPrice
	reqDB.Quantity=req.Quantity
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
	req.Amount=reqDB.TotalPrice
	res,bytes,err:=uc.host.Callback().Send(constants.TRANSACTION_ITEMS,req,header)
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