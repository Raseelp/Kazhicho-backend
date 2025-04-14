package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Review struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userID,omitempty"`
	Rating    float32            `bson:"rating"`
	Comment   string             `bson:"comment"`
	CreatedAt time.Time          `bson:"createdAt"`
	Images    []string           `bson:"images"`
}
