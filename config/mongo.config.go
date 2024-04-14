package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connects to the mongo database using mongo go driver
func ConnectMongoDatabase() *mongo.Client {

	// get environment variables
	MongoDbURI := os.Getenv("MONGO_DB_URI")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDbURI))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
