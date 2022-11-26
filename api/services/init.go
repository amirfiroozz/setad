package services

import (
	"setad/api/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	userCollection = getUserCollection()
	networkCollection = getNetworkCollection()
}

func getUserCollection() *mongo.Collection {
	return configs.GetCollection("user")
}

func getNetworkCollection() *mongo.Collection {
	return configs.GetCollection("network")
}
