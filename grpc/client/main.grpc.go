package grpc_client

import (
	grpc_merchant "cc-transaction/grpc/client/merchant"
)

type (
	host struct{
		merchant	grpc_merchant.MerchantInterface
		// transaction	transaction.TransactionInterface
		// conn		*grpc.ClientConn
	}
	GrpcInterface interface{
		Merchant() grpc_merchant.MerchantInterface
		// Transaction()transaction.TransactionInterface
	}
)

func InitGrpcClient(merchant grpc_merchant.MerchantInterface) GrpcInterface {
	
	return &host{
		merchant: merchant,
		// transaction: transaction,
		// conn:conn,
	}
}

func (g *host) Merchant()grpc_merchant.MerchantInterface{
	return g.merchant
}
// func (g *host) Transaction()transaction.TransactionInterface{
// 	return g.transaction
// }