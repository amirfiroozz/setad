package routers

import (
	"setad/api/controllers"

	"github.com/gin-gonic/gin"
)

func createUserRoutes(router *gin.RouterGroup) {
	router.GET("/test", controllers.Greet)
}
