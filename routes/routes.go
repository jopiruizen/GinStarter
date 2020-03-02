package routes

import (
	"github.com/gin-gonic/gin"
	"go-restapi/controllers/users"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/user/register", users.Register)
	router.POST("/user/find", users.Find)
	router.POST("/user/find2", users.FindOnFile)
	return router
}
