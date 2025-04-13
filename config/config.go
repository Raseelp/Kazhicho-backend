package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitConfig() {
	err := godotenv.Load("kazhicho.env")
	if err != nil {
		log.Fatal("Error loading .env files", err)
	}
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error", err)
	}
	DB = client.Database(dbName)
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
