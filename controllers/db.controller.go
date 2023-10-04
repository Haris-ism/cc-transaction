package controller

import (
	"net/http"
	"cc-transaction/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (c *controller) Ping(ctx *gin.Context) {
	res := models.GeneralResponse{
		Message: "pong euy",
		Code:    200,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *controller) WriteRedis(ctx *gin.Context) {
	res := models.GeneralResponse{
		Message: "success write euy",
		Code:    http.StatusOK,
	}
	req := models.RedisReq{}
	if err := ctx.BindJSON(&req); err != nil {
		logrus.Error("ieu error body:", err)
		res.Message = "Invalid Body Request"
		res.Code = http.StatusBadRequest
		ctx.JSON(res.Code, res)
		return
	}
	if err := c.usecase.WriteRedis(req); err != nil {
		logrus.Error("ieu error write redis:", err)
		res.Message = "Failed to Write Redis"
		res.Code = http.StatusInternalServerError
		ctx.JSON(res.Code, res)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *controller) ReadRedis(ctx *gin.Context) {
	res := models.GeneralResponse{
		Message: "success write euy",
		Code:    http.StatusOK,
	}
	req := models.RedisReq{}
	if err := ctx.BindJSON(&req); err != nil {
		logrus.Error("ieu error body:", err)
		res.Message = "Invalid Body Request"
		res.Code = http.StatusBadRequest
		ctx.JSON(res.Code, res)
		return
	}
	response, err := c.usecase.ReadRedis(req)
	if err != nil {
		res.Message = "Failed to Read Redis"
		res.Code = http.StatusInternalServerError
		ctx.JSON(res.Code, res)
		return
	}
	res.Message = response
	ctx.JSON(http.StatusOK, res)
}

func (c *controller) InsertPostgre(ctx *gin.Context) {
	res := models.GeneralResponse{
		Message: "success write euy",
		Code:    http.StatusOK,
	}
	req := models.ItemList{}
	if err := ctx.BindJSON(&req); err != nil {
		logrus.Error("ieu error body:", err)
		res.Message = "Invalid Body Request"
		res.Code = http.StatusBadRequest
		ctx.JSON(res.Code, res)
		return
	}
	if err := c.usecase.InsertDB(req); err != nil {
		res.Message = "Failed to Insert Postgre"
		res.Code = http.StatusInternalServerError
		ctx.JSON(res.Code, res)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
