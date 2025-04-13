package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FoodItem struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	//FoodSpots        []primitive.ObjectID `bson:"foodSpots,omitempty"`
	ShortDescription string   `bson:"shortDescription"`
	Price            float64  `bson:"price"`
	Rating           float32  `bson:"rating"`
	IsAvailable      bool     `bson:"isAvailable"`
	Images           []string `bson:"images"`
	Category         []string `bson:"category"`
	Calories         float64  `bson:"calories"`
}
