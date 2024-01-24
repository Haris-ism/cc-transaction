package controller_grpc

import (
	"cc-transaction/constants"
	"cc-transaction/protogen/merchant"
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func (g *ControllerGrpc) CallbackTransItems(ctx context.Context, req *merchant.ReqCallbackItems) (*merchant.ResMerchantCallbackModel, error){
	fmt.Println("masuk con")
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println("ieu md:",md.Get("tes")[0])
	fmt.Println("ieu ok:",ok)
	res:=merchant.ResMerchantCallbackModel{
		Message:constants.SUCCESS,
		Code:http.StatusOK,
	}
	result,err:=g.uc.CallbackTransItems(ctx,req)
	
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
	}
	res.Data=result
	return &res,nil
}

// func (g *ControllerGrpc) InquiryDiscounts(context.Context, *emptypb.Empty) (*merchant.InquiryMerchantDiscountsModel, error){
// 	fmt.Println("masuk con")
// 	res:=merchant.InquiryMerchantDiscountsModel{
// 		Message:constants.SUCCESS,
// 		Code:http.StatusOK,
// 	}
// 	result,err:=g.uc.InquiryDiscounts()
	
// 	if err!=nil{
// 		res.Message=err.Error()
// 		res.Code=http.StatusInternalServerError
// 	}
// 	res.Data=result
// 	return &res,nil
// }