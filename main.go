package main

import (
	controller "cc-transaction/controllers"
	postgre "cc-transaction/databases/postgresql"
	redis_db "cc-transaction/databases/redis"
	host "cc-transaction/hosts"
	merchantHost "cc-transaction/hosts/merchant"
	router "cc-transaction/routers"
	usecase "cc-transaction/usecases"
)

func main() {

	merchants:=merchantHost.InitMerchant()
	host:=host.InitHost(merchants)
	postgre := postgre.InitPostgre()
	redis := redis_db.InitRedis()
	uc := usecase.InitUsecase(postgre, redis, host)
	con := controller.InitController(uc)

	router.MainRouter(con)

}

