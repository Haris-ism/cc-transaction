package grpc_merchant

import (
	"cc-transaction/protogen/merchant"
	"cc-transaction/utils"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type(
	merchantGrpc struct{
		merchantConn merchant.MerchantServicesClient
	}
	MerchantInterface interface{
		TransItems(ctx context.Context, req *merchant.ReqTransItemsModel)(*merchant.ResMerchantTransModel, error)
	}
)

func (g *merchantGrpc)TransItems(ctx context.Context, req *merchant.ReqTransItemsModel)(*merchant.ResMerchantTransModel, error){
	res,err:=g.merchantConn.TransItems(ctx, req)
	if err != nil {
		log.Println("Error on grpc merchant :", err)
		return res,err
	}

	return res,nil
}


func InitGrpcMerchant()MerchantInterface{
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(utils.GetEnv("CALLBACK_HOST_GRPC"),opts...)
	if err!=nil{
		log.Println("failed to dial grpc merchant:",err)
	}
	
	merchantConn:=merchant.NewMerchantServicesClient(conn)
	log.Println("grpc merchant connected")
	return &merchantGrpc{
		merchantConn:merchantConn,
	}
}