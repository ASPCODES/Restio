package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


func databaseInstance() *mongo.Client {
	// loading .env file
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env File")
	}

	// Taking credentials from .env
	MongoDb := os.Getenv("MONGODB_URI")
	if MongoDb == "" {
		log.Fatal("MONGODB_URI not set in .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Connected succesfully!!")
	return client
}

var Client *mongo.Client = databaseInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection  {
	var collection *mongo.Collection = client.Database("restio").Collection(collectionName)
	return collection
}