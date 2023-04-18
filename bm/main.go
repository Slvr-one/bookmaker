package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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
	// if envLoadErr != nil {
	// 	log.Printf("err loading .env file: %s", envLoadErr)
	// }

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

// a func implementing a gambling option for a specific horse
func Invest(w http.ResponseWriter, r *http.Request) {
	// get the params from the url
	vars := mux.Vars(r)
	horseName := vars["horse"]
	amount := vars["amount"]

	// check if the horse exists
	horseExists := false
	for _, horse := range horses {
		if horse.Name == horseName {
			horseExists = true
			break
		}
	}
	if horseExists == false {
		fmt.Fprintf(w, "Horse does not exist")
		return

	}
	// check if the amount is a number
	amountIsNumber := false
	amountInt, err := fmt.Sscanf(amount, "%d")
	if err == nil {
		amountIsNumber = true
		if amountInt < 0 {
			fmt.Fprintf(w, "Amount must be a positive number")
			return
		}
		if amountInt > horse.Record.Wins {
			fmt.Fprintf(w, "Amount must be less than the wins")
			return
		}
		// check if the bet is already placed
		betExists := false
		for _, bet := range Bets {
			if bet.Horse == horseName {
				betExists = true
				break
			}
		}
		if betExists == true {
			fmt.Fprintf(w, "Bet already placed")
			return
		}
		// add the bet to the slice of bets
		bet := types.Bet{Horse: horseName, Amount: amountInt}
		Bets = append(Bets, bet)
		// update the wins and loses of the horse
		for i, horse := range horses {
			if horse.Name == horseName {
				horses[i].Record.Wins -= amountInt
				horses[i].Record.Loses += amountInt
				break
			}
		}
		// return the updated record
		fmt.Fprintf(w, "%v", horses)
		return
	}
	// if the amount is not a number
	fmt.Fprintf(w, "Amount must be a number")
	return
}

// a func to get all the available horses
func GetHorses(w http.ResponseWriter, r *http.Request) {
	// return the list of horses
	fmt.Fprintf(w, "%v", horses)
	return
}

// a func to get a specific horse
func GetHorse(w http.ResponseWriter, r *http.Request) {
	// get the params from the url
	vars := mux.Vars(r)
	horseName := vars["name"]
	// check if the horse exists
	horseExists := false
	for _, horse := range horses {
		if horse.Name == horseName {
			horseExists = true
			break
		}
	}
	if horseExists == false {
		fmt.Fprintf(w, "Horse does not exist")
		return
	}
	// return the horse
	fmt.Fprintf(w, "%v", horseName)
	return

}

// a func to update a horse
func updateHorse(w http.ResponseWriter, r *http.Request) {
	// get the params from the url
	vars := mux.Vars(r)
	horseName := vars["name"]
	// check if the horse exists
	horseExists := false
	for _, horse := range horses {
		if horse.Name == horseName {
			horseExists = true
			break

		}
	}
	if horseExists == false {
		fmt.Fprintf(w, "Horse does not exist")
		return
	}
	// get the new horse info
	// get the body of the request
	newHorse := types.Horse{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err := json.NewDecoder(r.Body).Decode(&newHorse)
		if err != nil {
			if err != nil {
				fmt.Fprintf(w, "Error reading body")
				fmt.Fprintf(w, "Error decoding body")
				log.Printf("err decoding json body. Err: %s", err)
				return
			}
		}
		// check if the new horse info is valid
		if newHorse.Name == "" {
			fmt.Fprintf(w, "Horse name cannot be empty")
			return

		}
		if newHorse.Color == "" {
			fmt.Fprintf(w, "Horse color cannot be empty")
			return

		}
		// update the horse info
		for i, horse := range horses {
			if horse.Name == horseName {
				horses[i].Name = newHorse.Name
				horses[i].Color = newHorse.Color
				break
			}
		}
		// return the updated horse
		fmt.Fprintf(w, "%v", horses)
		return
	}
}

// a func to monitor the app
func Monitor(w http.ResponseWriter, r *http.Request) {
	// get the metrics
	metrics := getMetrics()
	// return the metrics
	fmt.Fprintf(w, "%v", metrics)
	return

}

// a func to get the health of the app
func Health(w http.ResponseWriter, r *http.Request) {
	// return the health
	fmt.Fprintf(w, "Healthy")
	return
}

// a func to get the metrics
func getMetrics() string {
	// get the metrics
	metrics := "Horses: " + strconv.Itoa(len(horses)) + "\n"
	metrics += "Bets: " + strconv.Itoa(len(Bets)) + "\n"
	return metrics
}

// a func to get the static assets
func staticHandler() http.Handler {
	return http.FileServer(http.Dir("./static"))
}
