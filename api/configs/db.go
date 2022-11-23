package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_CLINET     *mongo.Client
	DB_NAME          string
	MONGODB_FULL_URL string
)

func init() {
	MONGO_CLINET = getMongoClient()
	DB_NAME = getDBName()
	MONGODB_FULL_URL = getMongoDBFullURL()
}

func getDBName() string {
	return os.Getenv("DB_NAME")
}
func getDBServerPort() string {
	return os.Getenv("MONGODB_SERVER_PORT")
}
func getDBServerIP() string {
	return os.Getenv("MONGODB_SERVER_IP")
}
func getMongoDBFullURL() string {
	return fmt.Sprintf("mongodb://%v:%v/%v", getDBServerIP(), getDBServerPort(), getDBName())
}

func getMongoClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGODB_FULL_URL))
	if err != nil {
		log.Fatal("Error creating new db client: ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connectionError := client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to db: ", connectionError)
	}
	fmt.Printf("connected to %v\n", MONGODB_FULL_URL)
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return MONGO_CLINET.Database(DB_NAME).Collection(collectionName)
}
