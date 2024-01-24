package controller_grpc

import (
	"cc-transaction/protogen/merchant"
	"context"

	"cc-transaction/usecase_grpc"

	"google.golang.org/grpc"
)

type (
	ControllerGrpc struct {
		Config *grpc.Server
		merchant.TransServicesServer
		uc usecase_grpc.UsecaseGrpcInterface
	}
	ControllerGrpcInterface interface {
		CallbackTransItems(ctx context.Context, req *merchant.ReqCallbackItems) (*merchant.ResMerchantCallbackModel, error)
		// InquiryItems(context.Context, *emptypb.Empty) (*merchant.InquiryMerchantItemsModel, error)
		// InquiryDiscounts(context.Context, *emptypb.Empty) (*merchant.InquiryMerchantDiscountsModel, error)
	}
)

func InitControllerGrpc(uc usecase_grpc.UsecaseGrpcInterface) ControllerGrpc {
	grpcConn:=grpc.NewServer()

	return ControllerGrpc{
		uc: uc,
		Config:grpcConn,
	}
}
