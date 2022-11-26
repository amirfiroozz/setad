package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	LoginRequest struct {
		PhoneNumber string `json:"phoneNumber"`
		Password    string `json:"password"`
	}
	SignupRequest struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Password    string `bson:"password" json:"password"`
		PhoneNumber string `json:"phoneNumber"`
	}
	AddToNetworkRequest struct {
		ParentID         *primitive.ObjectID `bson:"parentId"`
		ParentDepth      int                 `json:"parentDepth"`
		ChildPhoneNumber string              `json:"childPhoneNumber"`
		ChildFirstName   string              `json:"childFirstName"`
		ChildLastName    string              `json:"childLastName"`
	}
	JWT struct {
		ID          *primitive.ObjectID `bson:"_id"`
		Depth       int                 `json:"parentDepth"`
		PhoneNumber string              `json:"phoneNumber"`
	}
)

func NewSignupResuest() SignupRequest {
	return SignupRequest{}
}
func NewLoginResuest() LoginRequest {
	return LoginRequest{}
}
func NewAddToNetworkRequest() AddToNetworkRequest {
	return AddToNetworkRequest{}
}
func NewJWT() JWT {
	return JWT{}
}
