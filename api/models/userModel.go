package models

import (
	"setad/api/configs"
	"setad/api/structures"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID  `bson:"_id"`
	FirstName   string              `json:"firstName"`
	LastName    string              `json:"lastName"`
	PhoneNumber string              `json:"phoneNumber"`
	Role        string              `json:"role"`
	Password    string              `bson:"password" json:"password"`
	ParentID    *primitive.ObjectID `bson:"parentId"`
	Depth       int                 `json:"depth"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"`
}

func NewUser(singupReq structures.SignupRequest) User {
	var user User
	user.ID = primitive.NewObjectID()
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Role = configs.USER_ROLE
	user.FirstName = singupReq.FirstName
	user.LastName = singupReq.LastName
	user.Password = singupReq.Password
	user.PhoneNumber = singupReq.PhoneNumber
	return user
}
