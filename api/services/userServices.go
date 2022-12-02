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

func Signup(signupReq models.SignupRequest, parentId *primitive.ObjectID, parentDepth int) (*mongo.InsertOneResult, *utils.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := models.NewUser(signupReq, parentId, parentDepth)
	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return nil, utils.DBInsertionError
	}
	updatingChildrenOfParentErr := addChildToUser(parentId, &user.ID)
	if updatingChildrenOfParentErr != nil {
		return nil, updatingChildrenOfParentErr
	}
	return resultInsertionNumber, nil
}

func addChildToUser(parentId, childId *primitive.ObjectID) *utils.Error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": parentId}
	change := bson.M{"$push": bson.M{"children": childId}}

	_, updatingErr := userCollection.UpdateOne(ctx, filter, change)
	if updatingErr != nil {
		return utils.UpdatingChildrenError
	}
	return nil
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
	// unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$children"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "user"}, {Key: "localField", Value: "children"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "children"}}}}
	cur, err := userCollection.Aggregate(ctx, mongo.Pipeline{lookupStage})
	if err != nil {
		panic(err)
	}
	users := []models.UserResponse{}
	collectingUsersErr := cur.All(ctx, &users)
	if collectingUsersErr != nil {
		return nil, utils.UserCollectingError
	}
	return users, nil
}

func GetNetworksOfUser(userId primitive.ObjectID, maxDepth int) ([]*models.UserNetworkResponse, *utils.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conditions := generateConditions(userId, maxDepth)
	cur, aggregatingErr := userCollection.Aggregate(ctx, conditions)
	if aggregatingErr != nil {
		return nil, utils.AggregatingError
	}
	user := []models.UserResponse{}
	collectingUserNetworksErr := cur.All(ctx, &user)
	if collectingUserNetworksErr != nil {
		return nil, utils.UserNetworkFindingError
	}
	return user[0].Network, nil
}

func generateConditions(userId primitive.ObjectID, maxDepth int) []bson.M {
	if maxDepth == -1 {
		//to return all results
		maxDepth = 20
	}
	cond := make([]bson.M, 0)
	cond = append(cond, bson.M{"$match": bson.M{"_id": userId}})
	cond = append(cond, bson.M{
		"$graphLookup": bson.M{
			"from":             "user",
			"startWith":        "$children",
			"connectFromField": "children",
			"connectToField":   "_id",
			"as":               "network",
			"maxDepth":         maxDepth,
		}})
	cond = append(cond, bson.M{"$addFields": bson.M{"networkLength": bson.D{{Key: "$size", Value: "$network"}}}})
	return cond
}
