package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection
var FollowCollection *mongo.Collection

func InitializeMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Set a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	db := client.Database("ryde")
	UserCollection = db.Collection("users")
	FollowCollection = db.Collection("follows")
}
