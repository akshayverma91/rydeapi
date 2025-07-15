package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const users = "users"
const follows = "follows"

var UserCollection *mongo.Collection
var FollowCollection *mongo.Collection

func InitializeMongoDB() {
	// Read the Mongo Db url from the environment variable
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set. Did you create a .env file?")
	}

	// Read the Database Name from the environment variable
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Println("DB_NAME not set, using default 'testdb'")
		dbName = "rydeapi"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

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
	db := client.Database(dbName)
	UserCollection = db.Collection(users)
	FollowCollection = db.Collection(follows)
}
