package services

import (
	"context"
	"setad/api/configs"
	"setad/api/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func init() {
	userCollection = configs.GetCollection("user")
}

func Signup(user models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return nil, insertErr
	}
	return resultInsertionNumber, nil

}
