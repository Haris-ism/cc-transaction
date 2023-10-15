package main

import (
	controller "cc-transaction/controllers"
	postgre "cc-transaction/databases/postgresql"
	redis_db "cc-transaction/databases/redis"
	host "cc-transaction/hosts"
	callbackHost "cc-transaction/hosts/callback"
	merchantHost "cc-transaction/hosts/merchant"
	router "cc-transaction/routers"
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

	router.MainRouter(con)

}

