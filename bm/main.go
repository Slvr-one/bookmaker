package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"src/main.go/helpers"
	"src/main.go/types"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Init horses var as a slice of Horse struct (slice: an array with no fixed Size or Type)
var (
	horses []types.Horse
	bets   []types.Bet

	// serverPort = 5000
	// mongoPort  = 27017
)

func initMongoDB() {

	// Connect to the MongoDB database using the username and password

	//get port env var (if provided by user, default / 5000)
	mongoPort := os.Getenv("mongoPort")
	if mongoPort == "" {
		mongoPort = "27017"
	}

	mongodbUrl := fmt.Sprintf("mongodb://localhost:%s", mongoPort)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//init mongo client
	clientOptions := options.Client().ApplyURI(mongodbUrl)
	// clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb:%v", mongoPath))

	client, connectionErr := mongo.Connect(ctx, clientOptions)
	// client, connectionErr := mongo.NewClient(clientOptions)

	// terminate app if errored on db connection
	helpers.HandleErr(connectionErr, "err on db connection.")

	//check if MongoDB database has been found and connected
	pingErr := client.Ping(ctx, readpref.Primary())
	helpers.HandleErr(pingErr, "err on db ping test.")

	//init a collection (table) of bets in bookmaker db
	bmDB := client.Database("bookmaker")
	betsCollection := bmDB.Collection("bets")

	// insert multiple documents into a collection
	// by createing a slice of bson.D objects
	bets := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}, {"amount", 250}, {"profit", 500}},
		bson.D{{"fullName", "User 3"}, {"age", 20}, {"amount", 550}, {"profit", 1100}},
		bson.D{{"fullName", "User 4"}, {"age", 28}, {"amount", 420}, {"profit", -110}},
	}

	result, insertErr := betsCollection.InsertMany(ctx, bets)
	helpers.HandleErr(insertErr, "err on db objects insertion.")

	// display the ids of the newly inserted objects
	fmt.Println("Inserted a single document: ", result.InsertedID)
}

// main
func main() {
	// init a router handler for the server client
	router := mux.NewRouter()
	// fs := http.FileServer(http.Dir("static"))

	//get port env var (if provided by user, default / 5000)
	serverPort := os.Getenv("serverPort")
	if serverPort == "" {
		serverPort = "5000"
	}

	listenAddr := flag.String("listenaddr", serverPort, "port to serve the app")

	// // checking for a local .env file containing vars - redundant as of now
	// envLoadErr := godotenv.Load()
	// helpers.HandleErr(envLoadErr, "err loading .env file.")

	// init a server client with custom spec - for listen & serve
	s := &http.Server{
		Addr:           ":" + *listenAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // "1 times 2, 20 times" or 1048576 - standard sizing :)
	}

	// Hardcoded data - @todo: add database
	horses = append(horses, types.Horse{Name: "Monahen boy", Color: "brown", Record: &types.Record{Wins: 8, Loses: 3}})
	horses = append(horses, types.Horse{Name: "Dangerous", Color: "brown:white", Record: &types.Record{Wins: 7, Loses: 1}})
	horses = append(horses, types.Horse{Name: "horse 3", Color: "black", Record: &types.Record{Wins: 4, Loses: 5}})

	// router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// router.Handle("/assets/", http.StripPrefix("/assets/", staticHandler()).Methods("GET")

	router.HandleFunc("/", Welcom).Methods("GET")
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/metrics", Monitor).Methods("GET")

	router.HandleFunc("/LH", GetHorses).Methods("GET")          //List all available horses
	router.HandleFunc("/GH/{name}", GetHorse).Methods("GET")    //Get a specific horse
	router.HandleFunc("/UH/{name}", updateHorse).Methods("PUT") //Update a specific horse
	router.HandleFunc("/invest/{horse}/{amount}", Invest).Methods("UPDATE")

	log.Printf("set to listen on port: %v", serverPort)

	// log.Fatal(http.ListenAndServe(":"+serverPort, router))
	// log.Fatal(http.ListenAndServe(":"+*listenAddr, router))
	log.Fatal(s.ListenAndServe())

}
