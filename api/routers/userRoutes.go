package routers

import (
	"setad/api/controllers"
	"setad/api/middlewares"

	"github.com/gin-gonic/gin"
)

func createUserRoutes(router *gin.RouterGroup) {
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.Signup)
	router.GET("/mynetwork", middlewares.IfLoggedIn(), controllers.ShowAllNetworksOfUser)
	router.GET("/mychildren", middlewares.IfLoggedIn(), controllers.ShowAllChildrenOfUser)
}
