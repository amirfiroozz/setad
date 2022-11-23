package routers

import (
	"setad/api/controllers"

	"github.com/gin-gonic/gin"
)

func createUserRoutes(router *gin.RouterGroup) {
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.Signup)
}
