package utils

import (
	"errors"
	"net/http"
)

type Error struct {
	Msg        string `json:"msg"`
	Code       int    `json:"code"`
	StatusCode int    `json:"statusCode"`
}

var (
	ValidationError                    *Error
	ValidationError_Password           *Error
	ValidationError_PhoneNumber        *Error
	NoUserWithThisPhoneNumberError     *Error
	HashingPasswordError               *Error
	WrongPasswordError                 *Error
	UserAlreadyExistsError             *Error
	NoAuthHeaderError                  *Error
	JWTGeneratingError                 *Error
	JWTParsingError                    *Error
	JWTBodyDecodingError               *Error
	AlreadySignedup                    *Error
	AlreadyInUserNetworkError          *Error
	DBInsertionError                   *Error
	UserFindingError                   *Error
	UserNetworkFindingError            *Error
	NetworkFindingError                *Error
	UserCollectingError                *Error
	NetworkCollectingError             *Error
	BindingError                       *Error
	PhoneNumberNotExistsInNetworkError *Error
	UpdatingChildrenError              *Error
	AggregatingError                   *Error
	ReadingQueryParamError             *Error
	ServerError                        error
)

func init() {
	buildErrors()
}

func buildErrors() {
	ValidationError = newError("wrong validation", 0, http.StatusBadRequest)
	ValidationError_Password = newError("wrong validation for password", 0, http.StatusBadRequest)
	ValidationError_PhoneNumber = newError("wrong validation for phone number", 0, http.StatusBadRequest)
	NoUserWithThisPhoneNumberError = newError("no user with this phone number", 0, http.StatusNotFound)
	HashingPasswordError = newError("error while hashing password", 0, http.StatusInternalServerError)
	WrongPasswordError = newError("wrong password", 0, http.StatusForbidden)
	UserAlreadyExistsError = newError("user already exits with this phone number", 0, http.StatusConflict)
	NoAuthHeaderError = newError("header with name Token does not exits", 0, http.StatusBadRequest)
	JWTGeneratingError = newError("there was an error generating JWT token", 0, http.StatusInternalServerError)
	JWTParsingError = newError("there was an error in parsing", 0, http.StatusInternalServerError)
	JWTBodyDecodingError = newError("there was an error in decoding jwt body", 0, http.StatusInternalServerError)
	AlreadySignedup = newError("this phone number is already signed up.", 0, http.StatusConflict)
	AlreadyInUserNetworkError = newError("you already added this phone number to your network", 0, http.StatusConflict)
	DBInsertionError = newError("error while inserting record", 0, http.StatusInternalServerError)
	UserFindingError = newError("error while finding users from db", 0, http.StatusInternalServerError)
	UserNetworkFindingError = newError("error while finding networks of a user from db", 0, http.StatusInternalServerError)
	NetworkFindingError = newError("error while finding networks from db", 0, http.StatusInternalServerError)
	UserCollectingError = newError("error while collecting users from db result", 0, http.StatusInternalServerError)
	NetworkCollectingError = newError("error while collecting networks from db result", 0, http.StatusInternalServerError)
	BindingError = newError("error while binding body", 0, http.StatusInternalServerError)
	PhoneNumberNotExistsInNetworkError = newError("this phone number is not in network db", 0, http.StatusNotFound)
	UpdatingChildrenError = newError("failed updating children", 0, http.StatusInternalServerError)
	AggregatingError = newError("failed aggrigating", 0, http.StatusInternalServerError)
	ReadingQueryParamError = newError("wrong query params", 0, http.StatusBadRequest)
	ServerError = errors.New("error occurred")
}

//TODO: handle personal code response
func newError(msg string, code int, statusCode int) *Error {
	return &Error{Msg: msg, Code: code, StatusCode: statusCode}
}

func NewError(err error, code int) *Error {
	return &Error{Msg: err.Error(), Code: code, StatusCode: http.StatusInternalServerError}
}
