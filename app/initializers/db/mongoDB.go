package db

import (
	"context"
	"fmt"
	"time"

	h "github.com/Slvr-one/bookmaker/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

func MongoDB(mongodbUri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//init mongo client
	clientOptions := options.Client().ApplyURI(mongodbUri)
	// clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb:%v", mongoPath))

	// conn = &Conn{}
	// conn.Client, err = mongo.Connect(ctx, clientOptions)
	client, connErr := mongo.Connect(ctx, clientOptions)
	// client, connErr := mongo.NewClient(clientOptions)
	// Check(connErr, "err on db client creation.")
	// connErr = client.Connect(ctx)
	h.Check(connErr, "err on db connection.")

	defer func() {
		connErr := client.Disconnect(ctx)
		h.Check(connErr, "db dissconnected")
	}()

	//check if MongoDB database has been found and connected
	pingErr := client.Ping(ctx, readpref.Primary())
	h.Check(pingErr, "err on db ping test")

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// Check(err, "err on listing databases")
	// fmt.Println(databases)

	//init a collection (table) of bets in bookmaker db
	bmDB := client.Database("bookmaker")
	betsColl := bmDB.Collection("bets")
	Conn.Collection = betsColl

	MainBoard.Bets[0].Investor.UserName,
		MainBoard.Bets[1].Investor.UserName,
		MainBoard.Bets[2].Investor.UserName,
		MainBoard.Bets[3].Investor.UserName =
		"User 1", "dviross", "root", "Fat Guy"

	MainBoard.Bets[0].Investor.Age, MainBoard.Bets[1].Investor.Age, MainBoard.Bets[2].Investor.Age, MainBoard.Bets[3].Investor.Age = 30, 99, 52, 71
	MainBoard.Bets[0].Amount, MainBoard.Bets[1].Amount, MainBoard.Bets[2].Amount, MainBoard.Bets[3].Amount = 500, 70, 9900, 15
	MainBoard.Bets[0].Profit, MainBoard.Bets[1].Profit, MainBoard.Bets[2].Profit, MainBoard.Bets[3].Profit = -100, 400, -350, 1110
	// insert multiple documents into a collection
	// by createing a slice of bson.D objects
	var testBoard = []interface{}{
		bson.M{"fullName": "User 1", "age": 30, "amount": 1110, "profit": -100},
		bson.M{"fullName": "dviross", "age": 76, "amount": 78, "profit": 6900},
		bson.M{"fullName": "root", "age": 13, "amount": 2480, "profit": 1100},
		bson.M{"fullName": "Fat Guy", "age": 84, "amount": 30500, "profit": -13000},
	}

	// insertMany returns a result object with the ids of the newly inserted objects
	result, insertErr := betsColl.InsertMany(ctx, testBoard)
	h.Check(insertErr, "err on db objects insertion.")

	// display the ids of the newly inserted objects
	// fmt.Println("Inserted a single document: ", result.InsertedID)
	fmt.Println("Inserted a single document: ", result)

	// title := "Back to the Future"

	// var result bson.M
	// err = betsColl.FindOne(ctx, bson.D{{"title", title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	msg := fmt.Sprintf("No document was found with the title %s\n", title)
	// 	Check(err, msg)
	// }

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	return client, connErr
}
