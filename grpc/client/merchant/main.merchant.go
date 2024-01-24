package grpc_merchant

import (
	"cc-transaction/protogen/merchant"
	"cc-transaction/utils"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type(
	merchantGrpc struct{
		merchantConn merchant.InquiryServicesClient
		// callbackHost		string
		// inquiryItems		string
		// inquiryDiscounts	string
	}
	MerchantInterface interface{
		InquiryItems()(*merchant.InquiryMerchantItemsModel,error)
		InquiryDiscounts()(*merchant.InquiryMerchantDiscountsModel,error)
	}
)

func (g *merchantGrpc)InquiryItems()(*merchant.InquiryMerchantItemsModel,error){
	res,err:=g.merchantConn.InquiryItems(context.Background(),&emptypb.Empty{})
	if err != nil {
		log.Println("Error on grpc merchant :", err)
		return res,err
	}

	return res,nil
}

func (g *merchantGrpc)InquiryDiscounts()(*merchant.InquiryMerchantDiscountsModel,error){
	res,err:=g.merchantConn.InquiryDiscounts(context.Background(),&emptypb.Empty{})
	if err != nil {
		log.Println("Error on grpc merchant :", err)
		return res,err
	}

	return res,nil
}

func InitGrpcMerchant()MerchantInterface{
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// fmt.Println("merchant server:",utils.GetEnv("MERCHANT_HOST_GRPC"))
	conn, err := grpc.Dial(utils.GetEnv("MERCHANT_HOST_GRPC"),opts...)
	if err!=nil{
		log.Println("failed to dial grpc merchant:",err)
	}
	
	merchantConn:=merchant.NewInquiryServicesClient(conn)
	log.Println("grpc merchant connected")
	return &merchantGrpc{
		merchantConn:merchantConn,
	}
}