package utils

import (
	"fmt"
	"setad/api/configs"
	"setad/api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SendResponse(c *gin.Context, body interface{}, statusCode int) {
	c.JSON(statusCode, gin.H{
		"result": body,
	})
}

func HashPassword(password string) (string, *Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	if err != nil {
		return "", HashingPasswordError
	}
	return string(hashedPassword), nil
}

func IsWrongPassword(actual string, expected string) *Error {
	err := bcrypt.CompareHashAndPassword([]byte(expected), []byte(actual))
	if err != nil {
		return WrongPasswordError
	}
	return nil
}

func GenerateJWT(user models.User) (string, *Error) {
	JWT_SECRET := configs.JWT_SECRET
	JWT_EXP := configs.JWT_EXP
	var mySigningKey = []byte(JWT_SECRET)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = user.ID
	claims["depth"] = user.Depth
	claims["phoneNumber"] = user.PhoneNumber
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(JWT_EXP)).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", JWTGeneratingError
	}
	return tokenString, nil
}

func ToString(inp interface{}) string {
	return fmt.Sprintf("%v", inp)
}

func ToInt(inp interface{}) int {
	res, err := strconv.Atoi(ToString(inp))
	if err != nil {
		panic(err)
	}
	return res
}

func ToObjectID(inp interface{}) *primitive.ObjectID {
	res, err := primitive.ObjectIDFromHex(ToString(inp))
	if err != nil {
		panic(err)
	}
	return &res
}

func BindJSON(c *gin.Context, x interface{}) *Error {
	err := c.BindJSON(x)
	if err != nil {
		return BindingError
	}
	return nil
}
