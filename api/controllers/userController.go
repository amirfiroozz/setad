package controllers

import (
	"fmt"
	"net/http"
	"setad/api/models"
	"setad/api/services"
	"setad/api/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(c *gin.Context) {
	loginReq := models.NewLoginResuest()
	bindingError := utils.BindJSON(c, &loginReq)
	if utils.CheckErrorNotNil(c, bindingError) {
		return
	}
	validationError := utils.ValidateLoginRequest(loginReq, http.StatusBadRequest)
	if utils.CheckErrorNotNil(c, validationError) {
		return
	}
	user, noUserFounded := services.FindOneUserByPhoneNumber(loginReq.PhoneNumber)
	if utils.CheckErrorNotNil(c, noUserFounded) {
		return
	}
	wrongPasswordError := utils.IsWrongPassword(loginReq.Password, user.Password)
	if utils.CheckErrorNotNil(c, wrongPasswordError) {
		return
	}
	jwt, jwtError := utils.GenerateJWT(*user)
	if utils.CheckErrorNotNil(c, jwtError) {
		return
	}
	loginRes := models.NewLoginResponse("login successful", jwt, 1)
	utils.SendResponse(c, loginRes, http.StatusOK)
}

func generateSignupRequest(c *gin.Context) (*models.SignupRequest, *utils.Error) {
	signupReq := models.NewSignupResuest()
	bindingError := utils.BindJSON(c, &signupReq)
	if utils.CheckErrorNotNil(c, bindingError) {
		return nil, bindingError
	}
	validationError := utils.ValidateSignupRequest(c, signupReq, http.StatusBadRequest)
	if utils.CheckErrorNotNil(c, validationError) {
		return nil, validationError
	}
	return &signupReq, nil
}

func Signup(c *gin.Context) {
	signupReq, signupReqError := generateSignupRequest(c)
	if signupReqError != nil {
		return
	}
	parentId, noNetworkFounded := findParentId(signupReq.PhoneNumber)
	if utils.CheckErrorNotNil(c, noNetworkFounded) {
		return
	}
	_, noUserFounded := findUserByPhoneNumber(signupReq.PhoneNumber)
	if utils.CheckErrorNil(c, noUserFounded, utils.UserAlreadyExistsError) {
		return
	}
	passwordHash, hashingError := utils.HashPassword(signupReq.Password)
	if utils.CheckErrorNotNil(c, hashingError) {
		return
	}
	signupReq.Password = passwordHash
	result, signupError := services.Signup(*signupReq, parentId)
	if utils.CheckErrorNotNil(c, signupError) {
		return
	}
	singupRes := models.NewSignupResponse("user signup done!", result.InsertedID, 1)
	utils.SendResponse(c, singupRes, http.StatusOK)
}

func ShowAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if utils.CheckErrorNotNil(c, err) {
		return
	}
	utils.SendResponse(c, users, http.StatusOK)
}

func findParentId(phoneNumber string) (*primitive.ObjectID, *utils.Error) {
	networks, noNetworksFoundedErr := services.FindNetworksByPhoneNumber(phoneNumber)
	if noNetworksFoundedErr != nil {
		return nil, utils.PhoneNumberNotExistsInNetworkError
	}
	fmt.Println(networks)
	return nil, utils.PhoneNumberNotExistsInNetworkError
}

func findUserByPhoneNumber(phoneNumber string) (*models.User, *utils.Error) {
	return services.FindOneUserByPhoneNumber(phoneNumber)
}
