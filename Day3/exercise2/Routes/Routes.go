package Routes

import (
	"freshers-bootcamp/Day3/exercise2/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	returnValue := gin.Default()

	group := returnValue.Group("/student")
	{
		group.GET("",Controllers.GetAllRecords)
		group.POST("",Controllers.CreateStudent)
		group.GET(":id",Controllers.GetUserByID)
		group.PUT(":id", Controllers.UpdateUser)
		group.DELETE(":id",Controllers.DeleteUser)
	}

	return returnValue
}

