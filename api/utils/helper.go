package utils

import (
	"setad/api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SendResponse(c *gin.Context, body interface{}, statusCode int) {
	c.JSON(statusCode, gin.H{
		"result": body,
	})
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func IsWrongPassword(actual string, expected string) error {
	err := bcrypt.CompareHashAndPassword([]byte(expected), []byte(actual))
	if err != nil {
		return WrongPasswordError
	}
	return nil
}

func GenerateJWT(user models.User) (string, error) {
	return "this is jwt token.", nil
}
