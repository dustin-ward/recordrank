package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var err error
var client *mongo.Client
var collection *mongo.Collection

func Connect() {
	// "mongodb://mongodb:27017/"
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("dbtest").Collection("users")
	fmt.Println("Connected to DB!")
}

func Disconnect() {
	client.Disconnect(ctx)
	fmt.Println("Disconnected from DB")
}

func GetCollection() *mongo.Collection {
	return collection
}
