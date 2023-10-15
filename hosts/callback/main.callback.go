package callback

import (
	"cc-transaction/constants"
	"cc-transaction/utils"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

type(
	callback struct{
		callbackHost		string
		CallbackItems		string
	}
	CallbackInterface interface{
		Send(types string, body interface{},header http.Header)(gorequest.Response,[]byte,error)
	}
)

func (h *callback)Send(types string, body interface{},header http.Header)(gorequest.Response,[]byte,error){
	var url string
	var method string
	switch types{
		case constants.TRANSACTION_ITEMS:
			url=h.callbackHost+h.CallbackItems
			method=constants.HTTP_POST
	}
	res,data,err:=utils.HTTPRequest(url,method,body,header)
	
	if err!=nil{
		return res,data,err
	}
	return res,data,nil
}

func InitCallback()CallbackInterface{
	return &callback{
		callbackHost:utils.GetEnv("CALLBACK_HOST"),
		CallbackItems: utils.GetEnv("CALLBACK_ITEMS"),
	}
}