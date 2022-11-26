package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Network struct {
	ID               primitive.ObjectID  `bson:"_id"`
	ParentID         *primitive.ObjectID `bson:"parentId"`
	ParentDepth      int                 `json:"parentDepth"`
	ChildPhoneNumber string              `json:"childPhoneNumber"`
	ChildFirstName   string              `json:"childFirstName"`
	ChildLastName    string              `json:"childLastName"`
	CreatedAt        time.Time           `json:"createdAt"`
	UpdatedAt        time.Time           `json:"updatedAt"`
}

func NewNetwork(addReq AddToNetworkRequest) Network {
	var network Network
	network.ID = primitive.NewObjectID()
	network.ParentID = addReq.ParentID
	network.ParentDepth = addReq.ParentDepth
	network.ChildPhoneNumber = addReq.ChildPhoneNumber
	network.ChildFirstName = addReq.ChildFirstName
	network.ChildLastName = addReq.ChildLastName
	network.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	network.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return network
}
