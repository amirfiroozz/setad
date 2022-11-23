package utils

import "errors"

var (
	ValidationError             error
	ValidationError_Password    error
	ValidationError_PhoneNumber error
)

func init() {
	buildErrors()
}

func buildErrors() {
	ValidationError = errors.New("wrong validation")
	ValidationError_Password = errors.New("password wrong")
	ValidationError_PhoneNumber = errors.New("phone number wrong")
}
