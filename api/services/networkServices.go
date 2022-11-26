package services

import (
	"context"
	"setad/api/models"
	"setad/api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var networkCollection *mongo.Collection

func AddNetwork(addReq models.AddToNetworkRequest) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	network := models.NewNetwork(addReq)
	resultInsertionNumber, insertErr := networkCollection.InsertOne(ctx, network)
	if insertErr != nil {
		return nil, insertErr
	}
	return resultInsertionNumber, nil
}
func FindOneNetworkByPhoneNumberAndParentId(parentId primitive.ObjectID, childPhoneNumber string) (*models.Network, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var network models.Network
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "childphonenumber", Value: childPhoneNumber}},
				bson.D{{Key: "parentId", Value: parentId}},
			},
		},
	}
	err := networkCollection.FindOne(ctx, filter).Decode(&network)
	if err != nil {
		return nil, utils.NoUserWithThisPhoneNumberError
	}
	return &network, nil
}
