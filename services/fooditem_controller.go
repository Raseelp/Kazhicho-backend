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

func AddFoodItemToFoodSpot(c *gin.Context) {
	foodSpotID := c.Param("foodspot_id")

	//parse incoming food item
	var foodItem models.FoodItem
	if err := c.ShouldBindJSON(&foodItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food item data"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	foodItemCollection := config.DB.Collection("fooditem")
	foodSpotCollection := config.DB.Collection("foodspot")

	//insert the food item into the foodItem collection
	result, err := foodItemCollection.InsertOne(ctx, foodItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to add food spot"})
		return
	}
	foodItemID := result.InsertedID.(primitive.ObjectID)

	//convert foodSpotID and update its menu
	spotID, err := primitive.ObjectIDFromHex(foodSpotID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid food spot id"})
		return
	}

	_, err = foodSpotCollection.UpdateOne(ctx, bson.M{"_id": spotID}, bson.M{"$push": bson.M{"menu": foodItemID}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "failed to Update the food spot menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food Item Added Successfully", "foodItemID": foodItemID.Hex()})
}

type EditFoodItemRequest struct {
	FoodItemID string                 `json:"foodItemID"`
	Updates    map[string]interface{} `json:"updates"`
}

func EditFoodItem(c *gin.Context) {
	//getting request and parsing it to id and updates
	var request EditFoodItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	//parsing the id to object id from string
	itemID, err := primitive.ObjectIDFromHex(request.FoodItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid food item body"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//remove the id field in the interface because I don't want to try to change an immutable field like _id
	delete(request.Updates, "_id")

	//set the new values to the old document
	update := bson.M{"$set": request.Updates}
	foodItemCollection := config.DB.Collection("fooditem")
	result, err := foodItemCollection.UpdateOne(ctx, bson.M{"_id": itemID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "failed to Update the food spot menu"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "food item does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "food item updated Successfully"})

}
