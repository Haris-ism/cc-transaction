package controller

import (
	"cc-transaction/constants"
	"cc-transaction/controllers/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller)InquiryItems(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message:constants.SUCCESS,
		Code:http.StatusOK,
	}
	result,err:=c.usecase.InquiryItems()
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
	}
	res.Data=result
	ctx.JSON(http.StatusOK,res)
}
func (c *controller)InquiryDiscounts(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message:constants.SUCCESS,
		Code:http.StatusOK,
	}
	result,err:=c.usecase.InquiryDiscounts()
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
	}
	res.Data=result
	ctx.JSON(http.StatusOK,res)
}