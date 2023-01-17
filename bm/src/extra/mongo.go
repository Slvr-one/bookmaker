package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseCollections struct {
	Bets *mongo.Collection
}

func initMongoDB() {
	//init mongo client
	url := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(url)

	client, error := mongo.Connect(context.TODO(), clientOptions)
	if error != nil {
		panic(error)
	}

	//check if MongoDB database has been found and connected
	if error := client.Ping(context.TODO(), readpref.Primary()); error != nil {
		panic(error)
	}

	bmdb := client.Database("bookmaker")
	betsCollection := bmdb.Collection("bets")

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	bets := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}, {"amount", 250}, {"profit", 500}},
		bson.D{{"fullName", "User 3"}, {"age", 20}, {"amount", 550}, {"profit", 1100}},
		bson.D{{"fullName", "User 4"}, {"age", 28}, {"amount", 420}, {"profit", -110}},
	}
	// insert the bson object slice using InsertMany()
	results, error := betsCollection.InsertMany(context.TODO(), bets)
	// check for errors in the insertion
	if error != nil {
		panic(error)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)
	// fmt.Println("Inserted a single document: ", result.InsertedID)
}
