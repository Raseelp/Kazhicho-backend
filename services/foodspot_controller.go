package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kazhicho-backend/config"
	"kazhicho-backend/models"
)

func RequestRegisterFoodSpots(c *gin.Context) {
	var newSpot models.FoodSpot

	if err := c.ShouldBindJSON(&newSpot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Request"})
		return
	}
	newSpot.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//insert into the food_spot_request collection
	_, err := config.DB.Collection("food_spot_requests").InsertOne(ctx, newSpot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save Food Spot Request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Food Spot Request Sent Successfully"})
}

func GetfoodSpotRequests(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := config.DB.Collection("food_spot_requests")
	//	find the requests

	curser, err := collection.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error Fetching Food spot requests"})
		return
	}
	defer curser.Close(ctx)

	var requests []models.FoodSpot
	if err = curser.All(ctx, &requests); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error decoding the food spots"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": requests})

}
