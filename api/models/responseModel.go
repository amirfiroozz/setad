package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	SignupResponse struct {
		Message string      `json:"message"`
		UserID  interface{} `json:"userId"`
		Code    int         `json:"code"`
	}
	LoginResponse struct {
		Message string `json:"message"`
		JWT     string `json:"jwt"`
		Code    int    `json:"code"`
	}
	UserResponse struct {
		ID          primitive.ObjectID  `bson:"_id"`
		FirstName   string              `json:"firstName"`
		LastName    string              `json:"lastName"`
		PhoneNumber string              `json:"phoneNumber"`
		Role        string              `json:"role"`
		ParentID    *primitive.ObjectID `bson:"parentId"`
		Depth       int                 `json:"depth"`
		CreatedAt   time.Time           `json:"createdAt"`
		UpdatedAt   time.Time           `json:"updatedAt"`
	}
)

func NewSignupResponse(message string, userID interface{}, code int) SignupResponse {
	return SignupResponse{message, userID, code}
}
func NewLoginResponse(message string, jwt string, code int) LoginResponse {
	return LoginResponse{message, jwt, code}
}

func NewUserResponse(user User) UserResponse {
	var userRes UserResponse
	userRes.ID = user.ID
	userRes.FirstName = user.FirstName
	userRes.LastName = user.LastName
	userRes.PhoneNumber = user.PhoneNumber
	userRes.Role = user.Role
	userRes.ParentID = user.ParentID
	userRes.Depth = user.Depth
	userRes.CreatedAt = user.CreatedAt
	userRes.UpdatedAt = user.UpdatedAt
	return userRes
}
