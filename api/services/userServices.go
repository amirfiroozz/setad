package services

import (
	"context"
	"setad/api/configs"
	"setad/api/models"
	"setad/api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func init() {
	userCollection = configs.GetCollection("user")
}

func Signup(signupReq models.SignupRequest) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := models.NewUser(signupReq)
	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return nil, insertErr
	}
	return resultInsertionNumber, nil
}

func FindOneUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{
		"phonenumber": phoneNumber,
	}).Decode(&user)
	if err != nil {
		return nil, utils.NoUserWithThisPhoneNumberError
	}
	return &user, nil
}
