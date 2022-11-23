package routers

import "github.com/gin-gonic/gin"

func CreateRoutes(router *gin.RouterGroup) {
	createUserRoutes(router.Group("/users"))
}
