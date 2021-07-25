package Controllers

import (
	"fmt"
	"freshers-bootcamp/Day3/exercise2/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllRecords(context *gin.Context){
	var records [] Models.Student
	fmt.Println(records)

	err := Models.GetAllRecords(&records)

	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, records)
	}
}

func CreateStudent(context *gin.Context){
	var record Models.Student
	context.BindJSON(&record)
	err := Models.AddRecords(&record)

	if err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, record)
	}
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Student
	err := Models.GetRecordsByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.GetRecordsByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateRecord(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteRecord(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
