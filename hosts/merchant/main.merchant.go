package merchant

import (
	"cc-transaction/constants"
	"cc-transaction/utils"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

type(
	merchant struct{
		merchantHost	string
	}
	MerchantInterface interface{
		Send(types string, body interface{},header http.Header)(gorequest.Response,[]byte,error)
	}
)

func (m *merchant)Send(types string, body interface{},header http.Header)(gorequest.Response,[]byte,error){
	var url string
	var method string
	switch types{
		case "InquiryItems":
			url=m.merchantHost+"/v2/inquiry/items"
			method=constants.HTTP_GET
	}
	res,data,err:=utils.HTTPRequest(url,method,body,header)
	
	if err!=nil{
		return res,data,err
	}
	return res,data,nil
}

func InitMerchant()MerchantInterface{
	return &merchant{
		merchantHost:utils.GetEnv("MERCHANT_HOST"),
	}
}