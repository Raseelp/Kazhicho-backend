package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kazhicho-backend/config"
	"kazhicho-backend/models"
)

func AddDeal(c *gin.Context) {
	foodSpotID := c.Param("foodspot_id")
	spotID, err := primitive.ObjectIDFromHex(foodSpotID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid food spot id"})
		return
	}

	var deal models.Deal
	if err := c.ShouldBindJSON(&deal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Deal Data"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dealsCollection := config.DB.Collection("deals")
	foodSpotCollection := config.DB.Collection("foodspot")
	//add to the deals collection
	result, err := dealsCollection.InsertOne(ctx, &deal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to add deals to The deals Collection"})
		return
	}

	//get the inserted deals id
	dealID := result.InsertedID.(primitive.ObjectID)

	//push the deal id to the toDaysDeal Section in the foodSpot collection
	_, err = foodSpotCollection.UpdateOne(ctx, bson.M{"_id": spotID}, bson.M{"$push": bson.M{"todays_deals": dealID}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to add the dealsID to The corresponding Foods pot collection"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deal Added Successfully", "deal_id": dealID.Hex()})
}
