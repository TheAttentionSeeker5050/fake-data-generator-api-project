package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model definition API endpoint entries for MongoDB
type EndpointModel struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" Usage:"required"`
	Path      string             `json:"path" Usage:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
