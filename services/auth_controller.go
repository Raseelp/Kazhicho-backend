package services

import (
	"context"
	"kazhicho-backend/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"kazhicho-backend/models"
	"kazhicho-backend/utils"
)

var loginCollection *mongo.Collection
var userCollection *mongo.Collection

func Register(c *gin.Context) {
	var reqBody struct {
		Username string `bson:"username"`
		Email    string `bson:"email"`
		Password string `bson:"password"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Check if username already exists
	var existingUser models.Login
	err := loginCollection.FindOne(ctx, bson.M{"username": reqBody.Username}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Username already exists"})
		return
	}
	// Hash the password
	hashedPassword, err := utils.HashPassword(reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to hash the password"})
	}
	// Save to login collection
	login := models.Login{
		Username: reqBody.Username,
		Password: hashedPassword,
		Type:     "user",
	}

	_, err = loginCollection.InsertOne(ctx, login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save the login details"})
		return
	}

	user := models.User{
		Username: reqBody.Username,
		Email:    reqBody.Email,
	}
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save the user details"})
	}

	token, err := utils.GenarateJWT(reqBody.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error Genarating Token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Registered Successfully", "token": token})
}

func Login(c *gin.Context) {
	var loginData models.Login
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "invalid Request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Find the user in the login collection by username
	var storedLogin models.Login
	err := config.DB.Collection("login").FindOne(ctx, bson.M{"username": loginData.Username}).Decode(&storedLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Username or Password"})
		return
	}
	// Compare the password
	if !utils.CheckPasswordHash(loginData.Password, storedLogin.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Username or Password"})
		return
	}

	//Genarate auth token
	token, err := utils.GenarateJWT(loginData.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to genarate Token"})
		return
	}
	if storedLogin.Type == "admin" {
		c.JSON(http.StatusOK, gin.H{"Message": "Admin Login Successful", "token": token, "type": "admin"})
		return
	}
	if storedLogin.Type == "restaurant" {
		c.JSON(http.StatusOK, gin.H{"Message": "restaurant Login Successful", "token": token, "type": "restaurant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Login Successful", "token": token, "type": "user"})
}

func InitCollections(db *mongo.Database) {
	loginCollection = db.Collection("login")
	userCollection = db.Collection("user")
}
