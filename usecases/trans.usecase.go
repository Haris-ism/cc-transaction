package usecase

import (
	"cc-transaction/constants"
	"cc-transaction/hosts/merchant/models"
	"encoding/json"
	"errors"
	"net/http"
)

func (uc *usecase)InquiryItems()([]models.InquiryItems,error){
	result:=models.ResponseMerchantItems{}
	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Content-Type", "application/json")
	res,data,err:=uc.host.Merchant().Send(constants.INQUIRY_ITEMS,"",header)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_DB)
	}
	if res.StatusCode!=200{
		return result.Data, errors.New(constants.ERROR_INQUIRY)
	}
	err=json.Unmarshal(data,&result)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_INQUIRY)
	}
	return result.Data, nil
}
func (uc *usecase)InquiryDiscounts()([]models.InquiryDiscounts,error){
	result:=models.ResponseMerchantDiscounts{}
	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Content-Type", "application/json")
	res,data,err:=uc.host.Merchant().Send(constants.INQUIRY_DISCOUNTS,"",header)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_DB)
	}
	if res.StatusCode!=200{
		return result.Data, errors.New(constants.ERROR_INQUIRY)
	}
	err=json.Unmarshal(data,&result)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_INQUIRY)
	}
	return result.Data, nil
}