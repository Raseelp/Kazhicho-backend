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

func AddList(c *gin.Context) {
	//parsing the request body to list model
	var newList models.List
	if err := c.ShouldBindJSON(&newList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid List data"})
		return
	}
	newList.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//add the newList into the lists collection
	listCollection := config.DB.Collection("lists")

	_, err := listCollection.InsertOne(ctx, newList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error inserting new Lists into the lists collection"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully added new lists", "List": newList})
}
