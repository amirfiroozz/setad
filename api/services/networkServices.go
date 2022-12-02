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

func AddNetwork(addReq models.AddToNetworkRequest) (*mongo.InsertOneResult, *utils.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	network := models.NewNetwork(addReq)
	resultInsertionNumber, insertErr := networkCollection.InsertOne(ctx, network)
	if insertErr != nil {
		return nil, utils.DBInsertionError
	}
	return resultInsertionNumber, nil
}
func FindOneNetworkByPhoneNumberAndParentId(parentId primitive.ObjectID, childPhoneNumber string) (*models.Network, *utils.Error) {
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

func FindNetworksByPhoneNumber(childPhoneNumber string) ([]*models.Network, *utils.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "childphonenumber", Value: childPhoneNumber}}
	var networks []*models.Network
	cur, _ := networkCollection.Find(ctx, filter)
	cur.All(ctx, &networks)
	return networks, nil
}

func GetAllNetworks() ([]models.Network, *utils.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	networks := []models.Network{}
	cur, findNetworksErr := networkCollection.Find(ctx, bson.M{})
	if findNetworksErr != nil {
		return nil, utils.NetworkFindingError
	}
	collectingNetworksErr := cur.All(ctx, &networks)
	if collectingNetworksErr != nil {
		return nil, utils.NetworkCollectingError
	}
	return networks, nil
}
