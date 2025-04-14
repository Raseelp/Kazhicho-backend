package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deal struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	Name                  string             `bson:"name"`
	FoodSpotID            primitive.ObjectID `bson:"foodSpotID"`
	Description           string             `bson:"description"`
	Price                 float64            `bson:"price"`
	Discount              float64            `bson:"discount"`
	Image                 string             `bson:"image"`
	IsAvailable           bool               `bson:"isAvailable"`
	claimLimit            uint16             `bson:"claimLimit"`
	category              []string           `bson:"category"`
	isExclusiveForPremium bool               `bson:"isExclusiveForPremium"`
	ClaimedBy             []ClaimEntry       `bson:"claimedBy"`
	StartDate             time.Time          `bson:"startDate"`
	EndDate               time.Time          `bson:"endDate"`
}

type ClaimEntry struct {
	UserID    primitive.ObjectID `bson:"userID,omitempty"`
	claimedAt time.Time          `bson:"claimedAt,omitempty"`
	isUsed    bool               `bson:"isUsed,omitempty"`
}
