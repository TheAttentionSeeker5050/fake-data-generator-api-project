package repositories

import (
	"context"
	"os"

	"example.com/main/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// A mongo driver repository for endpoint model
type MongoEndpointRepository struct {
	Conn       *mongo.Client
	Database   string
	Collection string
}

// This function creates an object that represents the mongo endpoint repository
func NewMongoEndpointRepository(conn *mongo.Client) *MongoEndpointRepository {
	return &MongoEndpointRepository{
		Conn:       conn,
		Database:   os.Getenv("MONGO_DB_NAME"),
		Collection: os.Getenv("MONGO_ENDPOINT_COLLECTION"),
	}
}

// This function creates a new endpoint in the database using the mongo repository
func (r *MongoEndpointRepository) CreateEndpoint(endpoint *models.EndpointModel) error {
	// create the endpoint
	_, err := r.Conn.Database(r.Database).Collection(r.Collection).InsertOne(context.Background(), endpoint)
	return err
}

// This function returns an endpoint by its id using the mongo repository
func (r *MongoEndpointRepository) GetEndpointById(id string) (*models.EndpointModel, error) {
	// create a new empty endpoint model object
	endpoint := &models.EndpointModel{}

	// return the endpoint, save it in the endpoint object
	err := r.Conn.Database(r.Database).Collection(r.Collection).FindOne(context.Background(), map[string]string{"id": id}).Decode(endpoint)
	return endpoint, err
}

// This function returns all the endpoints using the mongo repository for the user id. This returns an array of endpoints
func (r *MongoEndpointRepository) GetEndpointsByUserId(userId string) ([]models.EndpointModel, error) {
	// create a new empty endpoint model object
	endpoints := []models.EndpointModel{}

	// return the endpoints, save them in a cursor object
	cursor, cursorErr := r.Conn.Database(r.Database).Collection(r.Collection).Find(context.Background(), map[string]string{"user_id": userId})
	if cursorErr != nil {
		return nil, cursorErr
	}

	// iterate over the cursor and append the endpoints
	for cursor.Next(context.Background()) {
		endpoint := models.EndpointModel{}
		decodeErr := cursor.Decode(&endpoint)
		if decodeErr != nil {
			return nil, decodeErr
		}
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}

// This function updates an endpoint by its id using the mongo repository
func (r *MongoEndpointRepository) UpdateEndpointById(id string, endpoint *models.EndpointModel) error {
	// update the endpoint
	_, err := r.Conn.Database(r.Database).Collection(r.Collection).UpdateOne(context.Background(), map[string]string{"id": id}, endpoint)

	// if there is an error, return it
	return err
}

// This function deletes an endpoint by its id using the mongo repository
func (r *MongoEndpointRepository) DeleteEndpointById(id string) error {
	// delete the endpoint
	_, err := r.Conn.Database(r.Database).Collection(r.Collection).DeleteOne(context.Background(), map[string]string{"id": id})

	// if there is an error, return
	return err
}
