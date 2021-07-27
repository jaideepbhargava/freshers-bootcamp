package Controllers

import (
	"fmt"
	"freshers-bootcamp/Day4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProducts(context *gin.Context) {
	var product []Models.Product

	err := Models.GetAllProducts(&product)

	if err != nil {
		context.AbortWithStatusJSON(200, gin.H{"status": false, "message": err.Error()})
	} else {
		context.JSON(http.StatusOK, product)
	}
}

func AddProduct(context *gin.Context){
	var product Models.Product
	context.BindJSON(&product)
	err := Models.AddProduct(&product)

	if err != nil {
		fmt.Println()
		context.AbortWithStatusJSON(200, gin.H{"status": false, "message": err.Error()})
	} else {
		context.JSON(http.StatusOK, product)
	}
}

func GetProductByID(context *gin.Context){
	var product Models.Product
	id := context.Params.ByName("product_id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		context.AbortWithStatusJSON(200, gin.H{"status": false, "message": err.Error()})
	} else {
		context.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(context *gin.Context){
	var product Models.Product
	id := context.Params.ByName("product_id")
	context.BindJSON(&product)
	err := Models.UpdateProduct(&product, id)
	if err != nil {
		context.AbortWithStatusJSON(200, gin.H{"status": false, "message": err.Error()})
	} else {
		context.JSON(http.StatusOK, product)
	}
}
