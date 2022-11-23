package routers

import "github.com/gin-gonic/gin"

func createUserRoutes(router *gin.RouterGroup) {
	router.GET("/test", test)
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": []string{"test", "test"},
	})
}
