package services

import (
	"context"
	"setad/api/models"
	"time"

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
