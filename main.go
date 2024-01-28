package main

import (
	"cc-transaction/controller_grpc"
	controller "cc-transaction/controllers"
	postgre "cc-transaction/databases/postgresql"
	redis_db "cc-transaction/databases/redis"
	grpc_client "cc-transaction/grpc/client"
	grpc_merchant "cc-transaction/grpc/client/merchant"
	grpc_server "cc-transaction/grpc/server"
	host "cc-transaction/hosts"
	callbackHost "cc-transaction/hosts/callback"
	merchantHost "cc-transaction/hosts/merchant"
	router "cc-transaction/routers"
	"cc-transaction/usecase_grpc"
	usecase "cc-transaction/usecases"
)

func main() {
	callback:=callbackHost.InitCallback()
	merchants:=merchantHost.InitMerchant()
	host:=host.InitHost(merchants,callback)
	postgre := postgre.InitPostgre()
	redis := redis_db.InitRedis()
	uc := usecase.InitUsecase(postgre, redis, host)
	con := controller.InitController(uc)

	merchantGrpc:=grpc_merchant.InitGrpcMerchant()
	hostGrpc:=grpc_client.InitGrpcClient(merchantGrpc)
	ucGrpc:=usecase_grpc.InitUsecaseGrpc(postgre,redis,hostGrpc)
	conGrpc:=controller_grpc.InitControllerGrpc(ucGrpc)
	go func(){
		// res:=hostGrpc.Merchant().InquiryItems()
		// fmt.Println("res grpc merchant:",res)
		grpc_server.InitGrpcServer(conGrpc)
	}()
	router.MainRouter(con)

}

