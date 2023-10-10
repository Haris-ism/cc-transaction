package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Item			string		`gorm:"column:item"`
	Quantity		int			`gorm:"column:quantity"`
	Discount		string		`gorm:"column:discount"`
	InitPrice		int			`gorm:"column:init_price"`
	TotalPrice		int			`gorm:"column:total_price"`
	Status			string		`gorm:"column:status"`
	Buyer			string		`gorm:"column:buyer"`
}

func (t *Order)TableName()string{
	return"orders"
}