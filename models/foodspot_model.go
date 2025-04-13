package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FoodSpot struct {
	ID               primitive.ObjectID   `bson:"_id,omitempty"`
	Name             string               `bson:"name"`
	Latitude         float64              `bson:"latitude"`
	Longitude        float64              `bson:"longitude"`
	Address          string               `bson:"address"`
	ShortDescription string               `bson:"short_description"`
	Taglines         []string             `bson:"taglines"`
	PhoneNumber      []string             `bson:"phone_number"`
	Image            string               `bson:"image"`
	Email            string               `bson:"email"`
	Website          string               `bson:"website"`
	Images           []string             `bson:"images"`
	Videos           []string             `bson:"videos"`
	IsOpen           bool                 `bson:"is_open"`
	OpensAt          string               `bson:"opens_at"`
	ClosesAt         string               `bson:"closes_at"`
	TodaysDeals      []primitive.ObjectID `bson:"todays_deals,omitempty"`
	Menu             []primitive.ObjectID `bson:"menu,omitempty"`
	Rating           float32              `bson:"rating"`
	Reviews          []primitive.ObjectID `bson:"reviews,omitempty"`
}
