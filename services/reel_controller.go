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

func UploadReel(c *gin.Context) {

	//parsing the request data to the Reel struct
	var newReel models.Reel
	if err := c.ShouldBindJSON(&newReel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Reel Data"})
		return
	}
	newReel.ID = primitive.NewObjectID()
	newReel.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//Insert the Reel to the reels collection
	reelCollection := config.DB.Collection("reels")
	_, err := reelCollection.InsertOne(ctx, newReel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error inserting reel to the reels Collection"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "new Reel Uploaded Successfully", "Reel": newReel})
}
