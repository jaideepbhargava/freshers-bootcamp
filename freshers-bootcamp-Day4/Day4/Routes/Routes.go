package Routes

import (
	"freshers-bootcamp/Day4/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	returnValue :=  gin.Default()

	group := returnValue.Group("/")
	{
		group.GET("/retailer/products",Controllers.GetAllProducts)
		group.GET("/retailer/product/:product_id",Controllers.GetProductByID)
		group.GET("/retailer/orders",Controllers.GetAllOrders)
		group.GET("/order/customer/:customer_id", Controllers.GetOrdersByCustomerID)
		group.POST("/retailer/product",Controllers.AddProduct)
		group.POST("/customer", Controllers.AddCustomer)
		group.POST("/customer/order", Controllers.AddOrder)
		group.PUT("/product/:product_id", Controllers.UpdateProduct)
	}
	return returnValue
}