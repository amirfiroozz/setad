package utils

import (
	"setad/api/structures"

	"github.com/gin-gonic/gin"
)

func ValidateLoginRequest(c *gin.Context, loginReq structures.LoginRequest, statusCode int) bool {
	//TODO: imp login validation
	if passwordWrongValidation(loginReq.Password) {
		sendError(c, newError(ValidationError_Password.Error(), statusCode))
		return true
	}
	if phoneNumberWrongValidation(loginReq.PhoneNumber) {
		sendError(c, newError(ValidationError_PhoneNumber.Error(), statusCode))
		return true
	}
	return false
}
func ValidateSignupRequest(c *gin.Context, signup structures.SignupRequest, statusCode int) bool {
	//TODO: imp signup validation
	if passwordWrongValidation(signup.Password) {
		sendError(c, newError(ValidationError_Password.Error(), statusCode))
		return true
	}
	if phoneNumberWrongValidation(signup.PhoneNumber) {
		sendError(c, newError(ValidationError_PhoneNumber.Error(), statusCode))
		return true
	}
	return false
}
func passwordWrongValidation(password string) bool {
	return len(password) < 4
}
func phoneNumberWrongValidation(phoneNumber string) bool {
	return len(phoneNumber) != 11
}
