package controller

import (
	"cc-transaction/constants"
	"cc-transaction/controllers/models"
	cModels "cc-transaction/hosts/callback/models"
	"cc-transaction/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (c *controller)TransItem(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code:http.StatusOK,
	}
	reqBody:=cModels.TransactionItems{}
	if err:=ctx.BindJSON(&reqBody);err!=nil{
		logrus.Error(err)
		res.Message=err.Error()
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	reqHeader:=models.ReqHeader{}
	if err:=ctx.BindHeader(&reqHeader);err!=nil{
		logrus.Error(err)
		res.Message=err.Error()
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	err:=utils.SignatureValidation(reqHeader,reqBody)
	if err!=nil{
		logrus.Error(err)
		res.Message=err.Error()
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	data,err:=c.usecase.TransItem(reqBody,reqHeader)
	if err!=nil{
		logrus.Error(err)
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}
	
	res.Data=data
	ctx.JSON(res.Code,res)
}