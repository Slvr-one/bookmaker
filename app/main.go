package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// http://localhost:9000/invest/%7BDangerous%7D/%7B500%7D

var (
	horses    []Horse
	MainBoard Board

	defaultServerPort = "9099"  //default port to serve app
	DefaultMongoPort  = "27017" // default port for mongoDB connection
	// defaultHost       = "localhost"

	conn = &Conn{}
)

// func init() {
// 	// rand.Seed(time.Now().UnixNano())
// 	// SqlDB()
// 	MainBoard.Title = "welcom to the Garrison; what we have today: "
// 	MainBoard.Footer = "hope to see tou here again"

// 	envLoadErr := godotenv.Load()
// 	Check(envLoadErr, "No .env file found")

// 	mongoPort, set := os.LookupEnv("mongoPort")
// 	if !set {
// 		LogToFile("mongoPort wasn't set, default is 27017")
// 		mongoPort = DefaultMongoPort
// 	}

// 	mongoHost, set := os.LookupEnv("mongoHost")
// 	if !set {
// 		LogToFile("mongoHost wasn't set, default is localhost")
// 		mongoHost = defaultHost
// 	}

// 	mongodbUrl := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
// 	conn.Client = MongoDB(mongodbUrl)
// }

// main
func main() {
	initLog()
	start := time.Now()
	defer End(start)
	MainBoard.Date = &start

	serverPort, set := os.LookupEnv("serverPort")
	if !set {
		LogToFile("serverPort env wasn't set, default is 9090.")
		serverPort = defaultServerPort
	}

	listenAddr := flag.String("listenaddr", serverPort, "port to serve the app")

	// // checking for a local .env file containing vars - redundant as of now
	// envLoadErr := godotenv.Load()
	// helpers.HandleErr(envLoadErr, "err loading .env file.")

	// init a server client with custom spec - for listen & serve
	mainServer := &http.Server{
		Addr:           ":" + *listenAddr,
		Handler:        InitR(), //router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // "1 times 2, 20 times" or 1048576 - standard size of header :)
		// BaseContext: func(l net.Listener) context.Context {
		// 	ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
		// 	return ctx
		// },
	}

	// Hardcoded data - @todo: add database
	horses = append(horses, Horse{Name: "Monahen Boy", Color: "brown", Record: &Record{Wins: 8, Loses: 3}})
	horses = append(horses, Horse{Name: "Dangerous", Color: "brown:white", Record: &Record{Wins: 7, Loses: 1}})
	horses = append(horses, Horse{Name: "Black Beauty", Color: "black", Record: &Record{Wins: 4, Loses: 5}})
	horses = append(horses, Horse{Name: "horse 4", Color: "black", Record: &Record{Wins: 4, Loses: 5}})

	log.Printf("🌏 set to listen on port: %v", serverPort)
	portInfo := fmt.Sprintf("set to listen on port: %v", serverPort)
	LogToFile(portInfo)

	log.Fatal(mainServer.ListenAndServe())

	// c := cors.New(cors.Options{
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	// })

	// handler := c.Handler(router)
	// http.ListenAndServe(":"+port, middlewares.LogRequest(handler))
	// log.Fatal(http.ListenAndServe(":"+serverPort, router))
	// log.Fatal(http.ListenAndServe(":"+*listenAddr, router))
}
