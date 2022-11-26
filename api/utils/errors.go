package utils

import "errors"

var (
	ValidationError                error
	ValidationError_Password       error
	ValidationError_PhoneNumber    error
	NoUserWithThisPhoneNumberError error
	WrongPasswordError             error
	UserAlreadyExists              error
	NoAuthHeaderError              error
	JWTParsingError                error
	JWTBodyDecodingError           error
)

func init() {
	buildErrors()
}

func buildErrors() {
	ValidationError = errors.New("wrong validation")
	ValidationError_Password = errors.New("wrong validation for password")
	ValidationError_PhoneNumber = errors.New(" wrong validation for phone number")
	NoUserWithThisPhoneNumberError = errors.New("no user with this phone number")
	WrongPasswordError = errors.New("wrong password")
	UserAlreadyExists = errors.New("user already exits with this phone number")
	NoAuthHeaderError = errors.New("header with name Token does not exits")
	JWTParsingError = errors.New("There was an error in parsing")
	JWTBodyDecodingError = errors.New("There was an error in decoding jwt body")
}
