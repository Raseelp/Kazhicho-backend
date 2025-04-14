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

func AddReview(c *gin.Context) {
	foodSpotID := c.Param("foodspot_id")
	spotID, err := primitive.ObjectIDFromHex(foodSpotID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user id"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//parse the request Review and store the user id and foodSpot id
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Review Data"})
		return
	}

	//add it into the reviews collections
	reviewCollection := config.DB.Collection("reviews")
	foodSpotCollection := config.DB.Collection("foodspot")
	result, err := reviewCollection.InsertOne(ctx, &review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to add the review to review collection"})
		return
	}

	//get the reviewID and add it into the Reviews section in foodSpotCollection
	reviewID := result.InsertedID.(primitive.ObjectID)
	_, err = foodSpotCollection.UpdateOne(ctx, bson.M{"_id": spotID}, bson.M{"$push": bson.M{"reviews": reviewID}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to add the review to foodSpot collection"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review Added Successfully", "review_id": reviewID.Hex()})

}
