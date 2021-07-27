package Models

import "gorm.io/gorm"

type Orders struct{
	gorm.Model
	//OrderId int `gorm:"primary_key"`
	OrderQuantity int
	OrderStatus string
	CustomerId int `json: -`
	Customer Customer `gorm:"foreignkey:customer_id"`
	ProductId int `json: -`
	Product Product `gorm:"foreignkey:product_id"`
}

func(order *Orders) TableName() string{
	return "orders"
}