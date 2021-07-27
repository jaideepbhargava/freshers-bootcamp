package Controllers

import (
	"freshers-bootcamp/Day4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCustomer(context *gin.Context){
	var customer Models.Customer
	context.BindJSON(&customer)
	err := Models.AddCustomer(&customer)

	if err != nil {
		context.AbortWithStatusJSON(200, gin.H{"status": false, "message": err.Error()})
	} else {
		context.JSON(http.StatusOK, customer)
	}
}