package routers

import (
	"setad/api/controllers"
	"setad/api/middlewares"

	"github.com/gin-gonic/gin"
)

func createNetworkRoutes(router *gin.RouterGroup) {
	router.POST("/add", middlewares.IfLoggedIn(), controllers.AddToNetwork)
}
