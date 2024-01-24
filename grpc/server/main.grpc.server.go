package grpc_server

import (
	"cc-transaction/controller_grpc"
	"cc-transaction/protogen/merchant"
	"cc-transaction/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

type (
	GrpcServer struct {
		Config *grpc.Server
		merchant.TransServicesServer
		TCP	net.Listener
	}
	ControllerGrpcInterface interface {
		// Run()
	}
)

func InitGrpcServer(grpcCon controller_grpc.ControllerGrpc)  {

	listen,err:=net.Listen("tcp",utils.GetEnv("PORT_GRPC"))
	if err!=nil{
		log.Println("failed to listen tcp:",err)
	}

	log.Println("register start")
	merchant.RegisterTransServicesServer(grpcCon.Config,&grpcCon)
	// merchant.RegisterInquiryServicesServer(grpcCon.Config,&grpcCon)
	log.Println("register grpc server")

	err=grpcCon.Config.Serve(listen)

	if err!=nil{
		log.Println("failed to listen grpc:",err)
	}
}