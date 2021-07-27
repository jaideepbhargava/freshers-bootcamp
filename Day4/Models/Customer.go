package Models

type Customer struct{
	CustomerId int `gorm:"primary_key"`
	CustomerName string
	CustomerEmail string
}

func (customer *Customer) TableName() string{
	return "customer"
}