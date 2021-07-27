package Models

import (
	"freshers-bootcamp/Day4/db"
	"sync"
)

var mutex = sync.Mutex{}

func GetAllProducts(product *[]Product) (err error){
	mutex.Lock()
	err = database.DB.Find(product).Error
	mutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error){
	if err = database.DB.Where("product_id = ?", id).First(product).Error; err != nil{
		return err
	}
	return nil
}

func GetAllOrders(order *[]Orders) (err error){
	if err = database.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func AddProduct(product *Product) (err error){

	mutex.Lock()
	err = database.DB.Create(product).Error
	mutex.Unlock()
	if err != nil{
		product.Message = "Product was not added"
		return err
	}
	product.Message = "product has been successfully added"
	return nil
}

func AddCustomer(customer *Customer) (err error){
	mutex.Lock()
	err = database.DB.Create(customer).Error
	mutex.Unlock()
	if err != nil{
		return err
	}

	return nil
}

func AddOrder(order *Orders) (err error){

	mutex.Lock()
	err = database.DB.Create(order).Error
	mutex.Unlock()
	if err != nil{
		order.OrderStatus = "order Failed"
		return err
	}
	order.OrderStatus = "Order Placed Successfully"
	return nil
}

func UpdateProduct(product *Product, id string) (err error){
	mutex.Lock()
	err = GetProductByID(product, id)
	mutex.Unlock()
	if err != nil {
		return err
	}
	database.DB.Save(product)
	return nil
}

func GetOrdersByCustomerID(order *[]Orders, id string) (err error){

	err = database.DB.Joins("customer").Find(order).Error

	if err != nil{
		return err
	}
	return nil

}