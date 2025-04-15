package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reel struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Caption        string               `bson:"caption"`
	VideoURL       string               `bson:"videoUrl"`
	TaggedFoodSpot []primitive.ObjectID `bson:"taggedFoodSpot,omitempty"`
	UploadedBy     primitive.ObjectID   `bson:"uploadedBy,omitempty"`
	Tags           []string             `bson:"tags"`
	likes          int                  `bson:"likes"`
	Comments       []Comment            `bson:"comments,omitempty"`
	CreatedAt      time.Time            `bson:"createdAt"`
}
type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    primitive.ObjectID `bson:"userId,omitempty"`
	Comment   string             `bson:"comment"`
	CreatedAt time.Time          `bson:"createdAt"`
}
