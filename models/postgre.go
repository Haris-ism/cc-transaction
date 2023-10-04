package models

import "gorm.io/gorm"

//table
type ItemList struct {
	gorm.Model
	Inquiry
}

type Inquiry struct {
	Item  string `json:"item" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
