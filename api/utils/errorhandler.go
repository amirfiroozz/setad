package utils

import (
	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, err *Error) {
	c.JSON(err.StatusCode, gin.H{
		"error": err.Msg,
		"code":  err.Code,
	})
}

func CheckErrorNotNil(c *gin.Context, err *Error) bool {
	if err != nil {
		sendError(c, err)
		return true
	}
	return false
}
func CheckErrorNil(c *gin.Context, nilError, sendingError *Error) bool {
	if nilError == nil {
		sendError(c, sendingError)
		return true
	}
	return false
}
