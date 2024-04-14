package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connects to the mongo database using mongo go driver
func ConnectMongoDatabase(connUri string) *mongo.Client {

	// get environment variables
	MongoDbURI := connUri

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDbURI))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// fmt.Println("Connected to MongoDB!")

	return client
}
