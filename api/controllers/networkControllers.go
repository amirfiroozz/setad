package controllers

import (
	"net/http"
	"setad/api/models"
	"setad/api/services"
	"setad/api/utils"

	"github.com/gin-gonic/gin"
)

func generateAddToNetworkRequest(c *gin.Context) (*models.AddToNetworkRequest, error) {
	addReq := models.NewAddToNetworkRequest()
	bindingError := c.BindJSON(&addReq)
	if utils.CheckErrorNotNil(c, bindingError, http.StatusInternalServerError) {
		return nil, bindingError
	}
	validationError := utils.ValidateAddToNetworkRequest(c, addReq, http.StatusBadRequest)
	if utils.CheckErrorNotNil(c, validationError, http.StatusBadRequest) {
		return nil, validationError
	}
	parentId, _ := c.Get("_id")
	parentDepth, _ := c.Get("depth")
	addReq.ParentID = utils.ToObjectID(parentId)
	addReq.ParentDepth = utils.ToInt(parentDepth)
	return &addReq, nil
}

func AddToNetwork(c *gin.Context) {
	addReq, addReqError := generateAddToNetworkRequest(c)
	if addReqError != nil {
		return
	}
	_, noUserFounded := services.FindOneUserByPhoneNumber(addReq.ChildPhoneNumber)
	if utils.CheckErrorNil(c, noUserFounded, utils.AlreadySignedup, http.StatusForbidden) {
		return
	}
	_, noNetworkFounded := services.FindOneNetworkByPhoneNumberAndParentId(*addReq.ParentID, addReq.ChildPhoneNumber)
	if utils.CheckErrorNil(c, noNetworkFounded, utils.AlreadyInUserNetwork, http.StatusForbidden) {
		return
	}
	result, addReqError := services.AddNetwork(*addReq)
	if utils.CheckErrorNotNil(c, addReqError, http.StatusInternalServerError) {
		return
	}
	utils.SendResponse(c, result, http.StatusOK)

}
