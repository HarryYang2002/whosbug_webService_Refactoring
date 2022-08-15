package routes

import (
	"github.com/gin-gonic/gin"
	"webService_Refactoring/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("v1/users")
	{
		router.POST("/registration")
	}
	r.Run(utils.HttpPort)
}
