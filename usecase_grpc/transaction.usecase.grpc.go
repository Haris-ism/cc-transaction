package usecase_grpc

import (
	"cc-transaction/protogen/merchant"
	"context"
	"log"
)

func (uc *usecaseGrpc)CallbackTransItems(ctx context.Context, req *merchant.ReqCallbackItems)(*merchant.ResCallbackItems, error){
	log.Println("masuk uc")
	result:=merchant.ResCallbackItems{}
	// fmt.Println("trans req:",req)
	// req,err:=utils.DecryptTransItemGrpc(req)
	// if err!=nil{
	// 	return &result, err
	// }
	
	// itemID,err:=strconv.Atoi(req.ItemID)
	// if err!=nil{
	// 	return &result, err
	// }
	// prices,err:=strconv.Atoi(req.Price)
	// if err!=nil{
	// 	return result.Data, err
	// }
	// qtys,err:=strconv.Atoi(req.Quantity)
	// if err!=nil{
	// 	return result.Data, err
	// }
	// percentages,_:=strconv.Atoi(req.Percentage)
	// if err!=nil{
	// 	return result.Data, err
	// }

	// cc,err:=uc.postgre.GetCC(req)
	// if err!=nil{
	// 	return result.Data,err
	// }

	// if req.CVV!=cc.CVV{
	// 	return result.Data,errors.New("Invalid Credit Cards")
	// }

	// discount:=float64(percentages)/100
	// price:=float64(prices)
	// qty:=float64(qtys)
	// reqTotalPrice:=int((price-(price*discount))*qty)

	// if reqTotalPrice>cc.Balance{
	// 	return result.Data,errors.New("Insuficient Balance")
	// }

	// cc.Balance -=reqTotalPrice
	
	// err=uc.postgre.DeductCC(cc)
	// if err!=nil{
	// 	return result.Data,err
	// }
	// res,err:=uc.host.Merchant().InquiryItems()
	// if err!=nil{
	// 	log.Println("err:",err)
	// 	return res.Data, errors.New(constants.ERROR_DB)
	// }
	// if res.Code!=200{
	// 	log.Println("err res code:",res.Code)
	// 	return res.Data, errors.New(constants.ERROR_INQUIRY)
	// }
	// log.Println("grpc res:",res)
	return &result,nil
}


// func (uc *usecase)TransItem(req models.TransactionItems,headers cModels.ReqHeader)(models.ResponseItems,error){
// 	result:=models.ResponseTransactionItems{}

// 	req,err:=utils.DecryptTransItem(req)
// 	if err!=nil{
// 		return result.Data, err
// 	}
	
// 	itemID,err:=strconv.Atoi(req.ItemID)
// 	if err!=nil{
// 		return result.Data, err
// 	}
// 	prices,err:=strconv.Atoi(req.Price)
// 	if err!=nil{
// 		return result.Data, err
// 	}
// 	qtys,err:=strconv.Atoi(req.Quantity)
// 	if err!=nil{
// 		return result.Data, err
// 	}
// 	percentages,_:=strconv.Atoi(req.Percentage)
// 	if err!=nil{
// 		return result.Data, err
// 	}

// 	cc,err:=uc.postgre.GetCC(req)
// 	if err!=nil{
// 		return result.Data,err
// 	}

// 	if req.CVV!=cc.CVV{
// 		return result.Data,errors.New("Invalid Credit Cards")
// 	}

// 	discount:=float64(percentages)/100
// 	price:=float64(prices)
// 	qty:=float64(qtys)
// 	reqTotalPrice:=int((price-(price*discount))*qty)

// 	if reqTotalPrice>cc.Balance{
// 		return result.Data,errors.New("Insuficient Balance")
// 	}

// 	cc.Balance -=reqTotalPrice
	
// 	err=uc.postgre.DeductCC(cc)
// 	if err!=nil{
// 		return result.Data,err
// 	}

// 	reqDB:=dbModels.Order{}
// 	reqDB.ItemID=itemID
// 	reqDB.Name=req.Name
// 	reqDB.Type=req.Type
// 	reqDB.Price=prices
// 	reqDB.TotalPrice=reqTotalPrice
// 	reqDB.Quantity=qtys
// 	reqDB.CC=req.CCNumber
// 	reqDB.Discount=req.Discount
// 	reqDB.Status=constants.STATUS_PENDING

// 	resDB,err:=uc.postgre.OrderTransItem(reqDB)
// 	if err!=nil{
// 		cc.Balance +=reqTotalPrice
// 		err:=uc.postgre.DeductCC(cc)
// 		if err!=nil{
// 			return result.Data,err
// 		}
// 		return result.Data,err
// 	}
// 	timeStamp:=time.Now().Format("15:04:05")
	
// 	// req.Amount=strconv.Itoa(reqDB.TotalPrice)
// 	reqHost:=models.ReqCallbackItems{}
// 	reqHost.Amount=strconv.Itoa(reqDB.TotalPrice)
// 	reqHost.CCNumber=req.CCNumber
// 	reqHost.Discount=req.Discount
// 	reqHost.ItemID=req.ItemID
// 	reqHost.Quantity=req.Quantity

// 	reqHost,err=utils.EncryptTransItem(reqHost)
// 	if err!=nil{
// 		return result.Data, err
// 	}
// 	bytes,err:=json.Marshal(reqHost)
// 	if err!=nil{
// 		return result.Data, err
// 	}
// 	signature:=utils.Signature(string(bytes),timeStamp)
// 	header := make(http.Header)
// 	header.Add("Accept", "*/*")
// 	header.Add("Content-Type", "application/json")
// 	header.Add("TimeStamp", timeStamp)
// 	header.Add("Signature", signature)
// 	_,bytes,err=uc.host.Callback().Send(constants.TRANSACTION_ITEMS,reqHost,header)
// 	if err!=nil{
// 		err:=uc.RollbackTrans(cc,resDB,reqTotalPrice)
// 		if err!=nil{
// 			return result.Data,err
// 		}
// 		return result.Data, errors.New(constants.ERROR_HOST)
// 	}
// 	err=json.Unmarshal(bytes,&result)
// 	if err!=nil{
// 		err:=uc.RollbackTrans(cc,resDB,reqTotalPrice)
// 		if err!=nil{
// 			return result.Data,err
// 		}
// 		return result.Data, errors.New(constants.ERROR_HOST)
// 	}
// 	if result.Code!=200{
// 		err:=uc.RollbackTrans(cc,resDB,reqTotalPrice)
// 		if err!=nil{
// 			return result.Data,err
// 		}
// 		return result.Data,errors.New(result.Message)
// 	}
// 	resDB.Status=constants.STATUS_SUCCESS
// 	err=uc.postgre.UpdateTransItem(resDB)
// 	if err!=nil{
// 		return result.Data,err
// 	}
// 	return result.Data,nil
// }