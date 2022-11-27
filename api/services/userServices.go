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

var userCollection *mongo.Collection

func Signup(signupReq models.SignupRequest, parentId *primitive.ObjectID) (*mongo.InsertOneResult, *utils.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := models.NewUser(signupReq, parentId)
	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return nil, utils.DBInsertionError
	}
	return resultInsertionNumber, nil
}

func FindOneUserByPhoneNumber(phoneNumber string) (*models.User, *utils.Error) {
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

func GetAllUsers() ([]models.UserResponse, *utils.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	users := []models.UserResponse{}
	cur, findUsersErr := userCollection.Find(ctx, bson.M{})
	if findUsersErr != nil {
		return nil, utils.UserFindingError
	}
	collectingUsersErr := cur.All(ctx, &users)
	if collectingUsersErr != nil {
		return nil, utils.UserCollectingError
	}
	return users, nil
}
