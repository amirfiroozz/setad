package utils

import (
	"setad/api/structures"

	"github.com/gin-gonic/gin"
)

func ValidateLoginRequest(loginReq structures.LoginRequest, statusCode int) error {
	//TODO: imp login validation
	if phoneNumberWrongValidation(loginReq.PhoneNumber) {
		return ValidationError_PhoneNumber
	}
	return nil
}

func ValidateSignupRequest(c *gin.Context, signup structures.SignupRequest, statusCode int) error {
	//TODO: imp signup validation
	if passwordWrongValidation(signup.Password) {
		return ValidationError_Password
	}
	if phoneNumberWrongValidation(signup.PhoneNumber) {
		return ValidationError_PhoneNumber
	}
	return nil
}

func passwordWrongValidation(password string) bool {
	return len(password) < 4
}

func phoneNumberWrongValidation(phoneNumber string) bool {
	return len(phoneNumber) != 11
}
