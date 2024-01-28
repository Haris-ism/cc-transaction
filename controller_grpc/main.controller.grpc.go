package controller_grpc

import (
	"cc-transaction/protogen/merchant"
	"context"
	"log"

	"cc-transaction/usecase_grpc"

	"google.golang.org/grpc"
)

type (
	ControllerGrpc struct {
		Config *grpc.Server
		merchant.MerchantServicesServer
		uc usecase_grpc.UsecaseGrpcInterface
	}
	ControllerGrpcInterface interface {
		TransItems(context.Context, *merchant.ReqTransItemsModel) (*merchant.ResMerchantTransModel, error)
	}
)

func InitControllerGrpc(uc usecase_grpc.UsecaseGrpcInterface) ControllerGrpc {
	grpcConn:=grpc.NewServer()
	log.Println("init grpc controller")
	return ControllerGrpc{
		uc: uc,
		Config:grpcConn,
	}
}
