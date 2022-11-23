package controllers

import (
	"net/http"
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
	validationError := utils.ValidateLoginRequest(loginReq, http.StatusBadRequest)
	if utils.CheckErrorNotNil(c, validationError, http.StatusBadRequest) {
		return
	}
	user, noUserFounded := services.FindOneUserByPhoneNumber(loginReq.PhoneNumber)
	if utils.CheckErrorNotNil(c, noUserFounded, http.StatusNotFound) {
		return
	}
	wrongPasswordError := utils.IsWrongPassword(loginReq.Password, user.Password)
	if utils.CheckErrorNotNil(c, wrongPasswordError, http.StatusUnauthorized) {
		return
	}
	jwt, jwtError := utils.GenerateJWT(*user)
	if utils.CheckErrorNotNil(c, jwtError, http.StatusInternalServerError) {
		return
	}
	loginRes := structures.NewLoginResponse("login successful", jwt, 1)
	utils.SendResponse(c, loginRes, http.StatusOK)
}

func Signup(c *gin.Context) {
	signupReq := structures.NewSignupResuest()
	err := c.BindJSON(&signupReq)
	if utils.CheckErrorNotNil(c, err, http.StatusInternalServerError) {
		return
	}
	validationError := utils.ValidateSignupRequest(c, signupReq, http.StatusBadRequest)
	if utils.CheckErrorNotNil(c, validationError, http.StatusBadRequest) {
		return
	}
	_, noUserFounded := services.FindOneUserByPhoneNumber(signupReq.PhoneNumber)
	if utils.CheckErrorNil(c, noUserFounded, utils.UserAlreadyExists, http.StatusBadRequest) {
		return
	}
	passwordHash, hashingError := utils.HashPassword(signupReq.Password)
	if utils.CheckErrorNotNil(c, hashingError, http.StatusInternalServerError) {
		return
	}
	signupReq.Password = passwordHash
	result, signupError := services.Signup(signupReq)
	if utils.CheckErrorNotNil(c, signupError, http.StatusInternalServerError) {
		return
	}
	singupRes := structures.NewSignupResponse("user signup done!", result.InsertedID, 1)
	utils.SendResponse(c, singupRes, http.StatusOK)
}
