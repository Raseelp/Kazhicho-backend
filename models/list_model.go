package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type List struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Title          string               `bson:"title"`
	IsFoodItemList bool                 `bson:"isFoodItemList"`
	IsPublic       bool                 `bson:"isPublic"`
	FoodItems      []primitive.ObjectID `bson:"foodItems,omitempty"`
	FoodSpots      []primitive.ObjectID `bson:"foodSpots,omitempty"`
	CreatedBy      string               `bson:"createdBy"`
	CreatedAt      time.Time            `bson:"createdAt"`
}
