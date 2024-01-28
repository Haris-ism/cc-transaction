package usecase_grpc

import (
	"cc-transaction/protogen/merchant"
	"log"

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
		TransItems(req *merchant.ReqTransItemsModel)(string, error)
	}
)

func InitUsecaseGrpc(postgre postgre.PostgreInterface, redis redis_db.RedisInterface, host grpc_client.GrpcInterface) UsecaseGrpcInterface {
	log.Println("init uc grpc")
	return &usecaseGrpc{
		postgre: postgre,
		redis:   redis,
		host: host,
	}
}
