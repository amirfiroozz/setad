package controllers

import (
	"net/http"
	"setad/api/models"
	"setad/api/services"
	"setad/api/utils"
	"strconv"

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
	parentId, parentDepth, noNetworkFounded := findParentIdAndParentDepth(signupReq.PhoneNumber)
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
	result, signupError := services.Signup(*signupReq, parentId, parentDepth)
	if utils.CheckErrorNotNil(c, signupError) {
		return
	}
	singupRes := models.NewSignupResponse("user signup done!", result.InsertedID, 1)
	utils.SendResponse(c, singupRes, http.StatusOK)
}

func ShowAllNetworksOfUser(c *gin.Context) {
	stringUserId, _ := c.Get("_id")
	userId := utils.ToObjectID(stringUserId)
	maxDepth, queryErr := getMaxDepthFromQuery(c)
	if utils.CheckErrorNotNil(c, queryErr) {
		return
	}
	networks, err := services.GetNetworksOfUser(*userId, maxDepth)
	utils.SortUserNetworks(networks)
	if utils.CheckErrorNotNil(c, err) {
		return
	}
	utils.SendResponse(c, networks, http.StatusOK)
}

func ShowAllChildrenOfUser(c *gin.Context) {
	stringUserId, _ := c.Get("_id")
	userId := utils.ToObjectID(stringUserId)
	networks, err := services.GetNetworksOfUser(*userId, 0)
	utils.SortUserNetworks(networks)
	if utils.CheckErrorNotNil(c, err) {
		return
	}
	utils.SendResponse(c, networks, http.StatusOK)
}

func findParentIdAndParentDepth(phoneNumber string) (*primitive.ObjectID, int, *utils.Error) {
	networks, noNetworksFoundedErr := services.FindNetworksByPhoneNumber(phoneNumber)
	if noNetworksFoundedErr != nil {
		return nil, 0, utils.PhoneNumberNotExistsInNetworkError
	}
	//TODO: check correct parentId
	return networks[0].ParentID, networks[0].ParentDepth, nil
}

func findUserByPhoneNumber(phoneNumber string) (*models.User, *utils.Error) {
	return services.FindOneUserByPhoneNumber(phoneNumber)
}

func getMaxDepthFromQuery(c *gin.Context) (int, *utils.Error) {
	maxDepth, convertingErr := strconv.Atoi(c.Query("maxdepth"))
	if convertingErr != nil {
		return -1, utils.ReadingQueryParamError
	}
	return maxDepth, nil
}
