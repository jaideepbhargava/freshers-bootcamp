package Routes

import (
	"freshers-bootcamp/Day3/exercise1/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	returnValue := gin.Default()
	group1 := returnValue.Group("/user-api")
	{
		group1.GET("user", Controllers.GetUsers)
		group1.POST("user", Controllers.CreateUser)
		group1.GET("user/:id", Controllers.GetUserByID)
		group1.PUT("user/:id", Controllers.UpdateUser)
		group1.DELETE("user/:id", Controllers.DeleteUser)
	}
	return returnValue
}