package controllers

import (
	"net/http"
	"setad/api/models"
	"setad/api/services"
	"setad/api/structures"
	"setad/api/utils"

	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {
	query := c.Request.URL.Query()
	param := c.Param("id")
	type Adderss struct {
		City   string `json:"city"`
		Street string `json:"street"`
	}
	type User struct {
		Name    string  `json:"name"`
		Number  int     `json:"number"`
		Adderss Adderss `json:"address"`
	}
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"query": query,
		"param": param,
		"body":  user,
	})
}

func Login(c *gin.Context) {
	loginReq := structures.NewLoginResuest()
	err := c.BindJSON(&loginReq)
	if utils.CheckErrorNotNil(c, err, http.StatusInternalServerError) {
		return
	}
	if utils.ValidateLoginRequest(c, loginReq, http.StatusBadRequest) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": loginReq,
	})
}

func Signup(c *gin.Context) {
	signupReq := structures.NewSignupResuest()
	err := c.BindJSON(&signupReq)
	if utils.CheckErrorNotNil(c, err, http.StatusInternalServerError) {
		return
	}
	if utils.ValidateSignupRequest(c, signupReq, http.StatusBadRequest) {
		return
	}
	passwordHash, hashingError := utils.HashPassword(signupReq.Password)
	if utils.CheckErrorNotNil(c, hashingError, http.StatusInternalServerError) {
		return
	}
	signupReq.Password = passwordHash
	user := models.NewUser(signupReq)
	result, signupError := services.Signup(user)
	if utils.CheckErrorNotNil(c, signupError, http.StatusInternalServerError) {
		return
	}
	singupRes := structures.NewSignupResponse("user signup done!", result.InsertedID, 1)
	c.JSON(http.StatusOK, gin.H{
		"result": singupRes,
	})
}
