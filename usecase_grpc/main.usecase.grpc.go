package usecase_grpc

import (
	"cc-transaction/protogen/merchant"
	"context"
	"fmt"

	postgre "cc-transaction/databases/postgresql"
	redis_db "cc-transaction/databases/redis"
	grpc_client "cc-transaction/grpc/client"
)

type (
	usecaseGrpc struct {
		postgre postgre.PostgreInterface
		redis   redis_db.RedisInterface
		host	grpc_client.GrpcInterface
	}
	UsecaseGrpcInterface interface {
		// InquiryItems()([]*merchant.InquiryItemsModel, error)
		// InquiryDiscounts()([]*merchant.InquiryDiscountsModel, error)
		CallbackTransItems(ctx context.Context, req *merchant.ReqCallbackItems)(*merchant.ResCallbackItems, error)
	}
)

func InitUsecaseGrpc(postgre postgre.PostgreInterface, redis redis_db.RedisInterface, host grpc_client.GrpcInterface) UsecaseGrpcInterface {
	fmt.Println("init uc grpc")
	return &usecaseGrpc{
		postgre: postgre,
		redis:   redis,
		host: host,
	}
}
