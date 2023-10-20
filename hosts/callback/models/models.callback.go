package models

type TransactionItems struct{
	ItemID			string		`json:"item_id"`
	Discount		string		`json:"discount"`
	Quantity		string		`json:"quantity"`
	CCNumber		string		`json:"cc_number"`
	CVV				string		`json:"cvv"`
	Amount			string		`json:"amount"`
	Price			string		`json:"price"`
	Name			string		`json:"name"`
	Type			string		`json:"type"`
	Percentage		string		`json:"percentage"`
}
type ReqCallbackItems struct{
	ItemID			string		`json:"item_id"`
	Discount		string		`json:"discount"`
	Quantity		string		`json:"quantity"`
	CCNumber		string		`json:"cc_number"`
	Amount			string		`json:"amount"`
}

type ResponseItems struct{
	ID			string		`json:"item_id"`
	Name		string		`json:"item_name"`
	Quantity	string		`json:"quantity"`
	CC			string		`json:"cc_number"`
	Code		string		`json:"code"`
}

type ResponseTransactionItems struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    ResponseItems
}

type DecTransItem struct{
	ID			string		`json:"item_id"`
	Name		string		`json:"item_name"`
	Quantity	string		`json:"quantity"`
	CC			string		`json:"cc_number"`
	Code		[]string	`json:"code"`
}