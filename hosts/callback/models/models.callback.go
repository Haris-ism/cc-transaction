package models

type TransactionItems struct{
	ItemID			int			`json:"item_id"`
	Discount		string		`json:"discount"`
	Quantity		int			`json:"quantity"`
	CCNumber		string		`json:"cc_number"`
	CVV				string		`json:"cvv"`
	Amount			int			`json:"amount"`
	Price			int			`json:"price"`
	Name			string		`json:"name"`
	Type			string		`json:"type"`
	Percentage		int			`json:"percentage"`
}

type ResponseItems struct{
	ID			int			`json:"item_id"`
	Name		string		`json:"item_name"`
	Quantity	int			`json:"quantity"`
	CC			string		`json:"cc_number"`
	Code		[]string	`json:"code"`
}

type ResponseTransactionItems struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    ResponseItems
}