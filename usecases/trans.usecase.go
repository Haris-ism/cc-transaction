package usecase

import (
	"cc-transaction/constants"
	"cc-transaction/hosts/merchant/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (uc *usecase)InquiryItems()([]models.InquiryItems,error){
	result:=models.ResponseMerchant{}
	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Content-Type", "application/json")
	res,data,err:=uc.host.Merchant().Send("InquiryItems","",header)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_DB)
	}
	if res.StatusCode!=200{
		return result.Data, errors.New(constants.ERROR_INQUIRY)
	}
	fmt.Println("ieu status:",res.StatusCode)
	fmt.Println("ieu byte:",data)
	err=json.Unmarshal(data,&result)
	if err!=nil{
		fmt.Println("failed unmarshal:",err)
		return result.Data, errors.New(constants.ERROR_INQUIRY)
	}
	fmt.Println("ieu result:",&result)
	return result.Data, nil
}