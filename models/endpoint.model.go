package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model definition API endpoint entries for MongoDB
type EndpointModel struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" Usage:"required" bson:"name"`
	Path      string             `json:"path" Usage:"required" bson:"path"`
	UserID    string             `json:"user_id" Usage:"required" bson:"user_id"`
	Method    string             `json:"method" Usage:"required" bson:"method"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
