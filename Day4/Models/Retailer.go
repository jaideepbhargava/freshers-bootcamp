package Models

type Product struct{
	ProductId int `gorm:"primary_key"`
	ProductName string
	ProductQuantity int
	ProductPrice int
	Message string
}

func (product *Product) TableName() string{
	return "retailer"
}