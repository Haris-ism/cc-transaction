package models

import "gorm.io/gorm"

type Order struct{
	gorm.Model
	ItemID			int			`json:"item_id" gorm:"column:item_id"`
	Name			string		`json:"name" gorm:"column:name"`
	Type			string		`json:"type" gorm:"column:type"`
	Discount		string		`json:"discount" gorm:"column:discount"`
	Price			int			`json:"price" gorm:"column:price"`
	TotalPrice		int			`json:"total_price" gorm:"column:total_price"`
	Quantity		int			`json:"quantity" gorm:"column:quantity"`
	CC				string		`json:"cc" gorm:"column:cc"`
	Status			string		`json:"status" gorm:"column:status"`
}

func (t *Order)TableName()string{
	return"orders"
}

type CreditCards struct{
	gorm.Model
	Bank			string			`gorm:"column:bank"`
	Limit			int				`gorm:"column:limit"`
	Balance			int				`gorm:"column:balance"`
	CC_Number		string			`gorm:"column:cc_number"`
	CVV				string			`gorm:"column:cvv"`
	CredsEmail		string			`gorm:"column:creds_email"`
}

func (t *CreditCards) TableName()string{
	return "credit_cards"
}

