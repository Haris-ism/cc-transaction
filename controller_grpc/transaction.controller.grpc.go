package controller_grpc

import (
	"cc-transaction/constants"
	"cc-transaction/protogen/merchant"
	"cc-transaction/utils"
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func (g *ControllerGrpc) TransItems(ctx context.Context,req *merchant.ReqTransItemsModel) (*merchant.ResMerchantTransModel, error){

	res:=&merchant.ResMerchantTransModel{
		Message:constants.SUCCESS,
		Code:http.StatusOK,
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return res,errors.New(constants.STATUS_FAILED)
	}
	err:=utils.SignatureValidationGrpc(md,req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		return res,err
	}
	resp,err:=g.uc.TransItems(req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		return res,err
	}
	res.Data=resp
	return res,nil
}
