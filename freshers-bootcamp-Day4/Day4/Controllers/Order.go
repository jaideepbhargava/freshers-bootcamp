package Controllers

import (
	"fmt"
	"freshers-bootcamp/Day4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOrders(context *gin.Context){
	var orderDetails [] Models.Orders

	err := Models.GetAllOrders(&orderDetails)

	if err != nil {
		context.AbortWithStatus(4)
	} else {
		context.JSON(http.StatusOK, orderDetails)
	}
}

func AddOrder(context *gin.Context){
	var order Models.Orders
	context.BindJSON(&order)
	err := Models.AddOrder(&order)

	if err != nil {
		fmt.Println()
		context.AbortWithStatus(4)
	} else {
		context.JSON(http.StatusOK, order)
	}
}

func GetOrdersByCustomerID(context *gin.Context){
	var order []Models.Orders
	id := context.Params.ByName("id")
	err := Models.GetOrdersByCustomerID(&order, id)
	if err != nil {
		context.AbortWithStatusJSON(200, gin.H{"status": false, "message": err.Error()})
	} else {
		context.JSON(http.StatusOK, order)
	}
}