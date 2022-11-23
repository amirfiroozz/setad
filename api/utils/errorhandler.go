package utils

import "github.com/gin-gonic/gin"

type Error struct {
	Msg        string `json:"msg"`
	Code       int    `json:"code"`
	StatusCode int    `json:"statusCode"`
}

func sendError(c *gin.Context, err Error) {
	c.JSON(err.StatusCode, gin.H{
		"error": err,
	})
}

//TODO: handle personal code response
func newError(msg string, statusCode int) Error {
	return Error{Msg: msg, Code: 0, StatusCode: statusCode}
}

func CheckErrorNotNil(c *gin.Context, err error, statusCode int) bool {
	if err != nil {
		sendError(c, newError(err.Error(), statusCode))
		return true
	}
	return false
}
