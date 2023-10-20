package middleware

import (
	"cc-transaction/constants"
	"cc-transaction/controllers/models"
	"cc-transaction/utils"
	"fmt"
	"net/http"

	hModels "cc-transaction/hosts/callback/models"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

func SignatureValidation(ctx *gin.Context){
	res:=models.GeneralResponse{}
	reqHeader:=models.ReqHeader{}
	err:=ctx.BindHeader(&reqHeader)
	if err!=nil{
		logrus.Error("err:",err)
		res.Message=constants.ERROR_TOKEN
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		ctx.Abort()
		return
	}
	fmt.Println("header:",reqHeader)
	reqBody:=hModels.TransactionItems{}
	if err:=ctx.BindJSON(&reqBody);err!=nil{
		logrus.Error("err:",err)
		res.Message=constants.INVALID_INPUT
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		ctx.Abort()
		return
	}
	body,err:=json.Marshal(reqBody)
	hash:=utils.Signature(string(body),reqHeader.TimeStamp)
	fmt.Println("hash:",hash)
	fmt.Println("sig:",reqHeader.Signature)
	if hash!=reqHeader.Signature{
		logrus.Error("err:",err)
		res.Message=constants.ERROR_TOKEN
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		ctx.Abort()
		return
	}
	ctx.Set("reqBody",reqBody)
	ctx.Next()
}

