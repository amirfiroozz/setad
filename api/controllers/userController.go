package controllers

import (
	"net/http"
	"setad/api/models"
	"setad/api/services"
	"setad/api/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	loginReq := models.NewLoginResuest()
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
	loginRes := models.NewLoginResponse("login successful", jwt, 1)
	utils.SendResponse(c, loginRes, http.StatusOK)
}

func generateSignupRequest(c *gin.Context) (*models.SignupRequest, error) {
	signupReq := models.NewSignupResuest()
	bindingError := c.BindJSON(&signupReq)
	if utils.CheckErrorNotNil(c, bindingError, http.StatusInternalServerError) {
		return nil, bindingError
	}
	validationError := utils.ValidateSignupRequest(c, signupReq, http.StatusBadRequest)
	if utils.CheckErrorNotNil(c, validationError, http.StatusBadRequest) {
		return nil, validationError
	}
	return &signupReq, nil
}

func Signup(c *gin.Context) {
	signupReq, signupReqError := generateSignupRequest(c)
	if signupReqError != nil {
		return
	}
	_, noUserFounded := services.FindOneUserByPhoneNumber(signupReq.PhoneNumber)
	if utils.CheckErrorNil(c, noUserFounded, utils.UserAlreadyExists, http.StatusBadRequest) {
		return
	}
	var parentId *primitive.ObjectID
	//TODO: here search if this phone number is added as someone's child
	passwordHash, hashingError := utils.HashPassword(signupReq.Password)
	if utils.CheckErrorNotNil(c, hashingError, http.StatusInternalServerError) {
		return
	}
	signupReq.Password = passwordHash
	result, signupError := services.Signup(*signupReq, parentId)
	if utils.CheckErrorNotNil(c, signupError, http.StatusInternalServerError) {
		return
	}
	singupRes := models.NewSignupResponse("user signup done!", result.InsertedID, 1)
	utils.SendResponse(c, singupRes, http.StatusOK)
}

func ShowAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if utils.CheckErrorNotNil(c, err, http.StatusNotFound) {
		return
	}
	utils.SendResponse(c, users, http.StatusOK)
}
