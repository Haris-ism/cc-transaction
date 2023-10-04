package models

type InquiryItems struct{
	Name			string		`json:"name" gorm:"column:name"`
	Type			string		`json:"type" gorm:"column:type"`
	Price			int			`json:"price" gorm:"column:price"`
	Quantity		int			`json:"quantity" gorm:"column:quantity"`
}

type ResponseMerchant struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    []InquiryItems
}