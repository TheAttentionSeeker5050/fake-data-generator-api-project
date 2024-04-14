package repositories

import (
	"context"
	"os"
	"testing"
	"time"

	"example.com/main/config"
	"example.com/main/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestCreateEndpoint tests the MongoEndpointRepository and its CRUD operations. It uses the MongoDB test database and the endpoints collection.
func TestEndpointsRepository(t *testing.T) {

	if os.Getenv("GIN_MODE") != "GITHUB_ACTIONS" {

		// Load the .env file in the current directory
		envVariableErr := godotenv.Load("../../.env")
		assert.Nil(t, envVariableErr, "Expected no error loading .env file")
		if envVariableErr != nil {
			return
		}
	}

	connUri := os.Getenv("MONGO_TEST_DB_URI")

	// call mongo config to connect to the test database
	client := config.ConnectMongoDatabase(connUri)
	defer client.Disconnect(context.Background())

	// Check if the client is nil
	assert.NotNil(t, client, "Expected client to be created")
	if client == nil {
		return
	}

	// Call the NewMongoEndpointRepository function with the real mongo.Client
	repo := NewMongoEndpointRepository(client)

	// Assert that the repository is created with the correct values
	assert.NotNil(t, repo, "Expected repository to be created")
	assert.Equal(t, os.Getenv("MONGO_DB_NAME"), repo.Database, "Expected database name to match environment variable")
	assert.Equal(t, os.Getenv("MONGO_ENDPOINT_COLLECTION"), repo.Collection, "Expected collection name to match environment variable")
	if repo == nil {
		return
	}

	// Create an endpoint model
	endpoint := &models.EndpointModel{
		ID:        primitive.NewObjectID(),
		UserID:    "123",
		Path:      "/api/v1/test",
		Name:      "Test Endpoint",
		Method:    "GET",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createErr := repo.CreateEndpoint(endpoint)
	assert.Nil(t, createErr, "Expected no error when creating an endpoint")
	if createErr != nil {
		return
	}
	// Get the endpoint by its ID
	endpointID := endpoint.ID.Hex()
	getEndpoint, getErr := repo.GetEndpointById(endpointID)

	assert.Nil(t, getErr, "Expected no error when getting an endpoint by ID")
	assert.Equal(t, endpointID, getEndpoint.ID.Hex(), "Expected the endpoint ID to match")
	if getErr != nil {
		return
	}

	// Add a second endpoint
	endpoint2 := &models.EndpointModel{
		ID:        primitive.NewObjectID(),
		UserID:    "123",
		Path:      "/api/v1/test2",
		Name:      "Test Endpoint 2",
		Method:    "POST",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createErr2 := repo.CreateEndpoint(endpoint2)
	assert.Nil(t, createErr2, "Expected no error when creating a second endpoint")
	if createErr2 != nil {
		return
	}

	// Get all the endpoints for the user
	endpoints, getErr2 := repo.GetEndpointsByUserId(endpoint2.UserID)
	assert.Nil(t, getErr2, "Expected no error when getting all endpoints by user ID")
	assert.Equal(t, 2, len(endpoints), "Expected two endpoints to be returned")
	if getErr2 != nil {
		return
	}

	// Delete all endpoints for all users using built-in database methods
	deleteCursor, deleteErr := client.Database(os.Getenv("MONGO_DB_NAME")).Collection(os.Getenv("MONGO_ENDPOINT_COLLECTION")).DeleteMany(context.Background(), primitive.M{})
	assert.Nil(t, deleteErr, "Expected no error when deleting all endpoints by user ID")
	assert.NotNil(t, deleteCursor, "Expected cursor to be created")
	if deleteErr != nil {
		return
	}

}
