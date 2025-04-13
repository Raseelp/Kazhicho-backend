package services

import (
	"context"
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
