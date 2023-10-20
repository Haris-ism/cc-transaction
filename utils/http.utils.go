package utils

import (
	"cc-transaction/constants"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

func HTTPRequest(url string, method string, body interface{}, header http.Header)(gorequest.Response,[]byte,error){
	var(
		res gorequest.Response
		data []byte
		err error
	)
	switch method{
		case constants.HTTP_GET:
			res,data,err=HTTPGET(url, header)
		case constants.HTTP_POST:
			res,data,err=HTTPPOST(url, body, header)
		case constants.HTTP_PUT:
			res,data,err=HTTPPUT(url, body, header)
		case constants.HTTP_DELETE:
			res,data,err=HTTPDELETE(url, body, header)
	}
	return res,data,err
}

func HTTPGET(url string, header http.Header)(gorequest.Response,[]byte,error){
	request:=gorequest.New()
	request.SetDebug(true)
	reqagent := request.Get(url)
	reqagent.Header = header
	res,data,err:=request.End()
	if err!=nil{
		return res,[]byte(data),err[0]
	}
	return res,[]byte(data),nil
}
func HTTPPOST(url string,body interface{}, header http.Header)(gorequest.Response,[]byte,error){
	request:=gorequest.New()
	request.SetDebug(true)
	reqagent := request.Post(url)
	reqagent.Header = header
	res,data,err:=reqagent.Send(body).End()
	if err!=nil{
		return res,[]byte(data),err[0]
	}
	return res,[]byte(data),nil
}
func HTTPPUT(url string,body interface{}, header http.Header)(gorequest.Response,[]byte,error){
	return nil,nil,nil
}
func HTTPDELETE(url string,body interface{}, header http.Header)(gorequest.Response,[]byte,error){
	return nil,nil,nil
}