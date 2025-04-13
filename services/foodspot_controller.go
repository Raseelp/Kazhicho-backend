package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kazhicho-backend/utils"
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

func ApproveFoodSpotRequest(c *gin.Context) {
	idParam := c.Param("id")
	requestID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Request ID"})
		return
	}
	//get username and password from the request body
	var loginData models.Login

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Login Data"})
		return
	}
	if loginData.Password == "" || loginData.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Username and password are required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	requestCollection := config.DB.Collection("food_spot_requests")
	approvedCollection := config.DB.Collection(("foodspot"))
	loginCollection := config.DB.Collection("login")
	// Find the request document
	var foodSpot models.FoodSpot
	err = requestCollection.FindOne(ctx, bson.M{"_id": requestID}).Decode(&foodSpot)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "FoodSpot Request not found"})
		return
	}
	//Hash password
	hashedPassword, err := utils.HashPassword(loginData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Unable to hash password"})
		return
	}
	//make a login entry
	loginEntry := models.Login{
		Username: loginData.Username,
		Password: hashedPassword,
		Type:     "foodspot",
	}
	//insert foodSpot login data to login collection
	_, err = loginCollection.InsertOne(ctx, loginEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save the login data"})
		return
	}
	// Insert into the Approved collection
	_, err = approvedCollection.InsertOne(ctx, foodSpot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to approve request"})
		return
	}

	// Delete from requests collection
	_, err = requestCollection.DeleteOne(ctx, bson.M{"_id": requestID})
	if err != nil {
		{
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete the food spot request"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "FoodSpot Request Approved Successfully"})
}

func RejectFoodSpotRequest(c *gin.Context) {
	idParam := c.Param("id")
	requestID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	requestsCollection := config.DB.Collection("food_spot_requests")

	// Just delete the request
	_, err = requestsCollection.DeleteOne(ctx, bson.M{"_id": requestID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reject request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food spot request rejected and removed"})
}
