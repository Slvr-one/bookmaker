package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	horses    []Horse
	MainBoard Board

	defaultServerPort = "9090"  //default port to serve app
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
	router := InitR()
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

	// Hardcoded data - @todo: add database
	horses = append(horses, Horse{Name: "Monahen Boy", Color: "brown", Record: &Record{Wins: 8, Losses: 3}})
	horses = append(horses, Horse{Name: "Dangerous", Color: "brown:white", Record: &Record{Wins: 7, Losses: 1}})
	horses = append(horses, Horse{Name: "Black Beauty", Color: "black", Record: &Record{Wins: 4, Losses: 5}})
	horses = append(horses, Horse{Name: "horse 4", Color: "black", Record: &Record{Wins: 4, Losses: 5}})

	log.Printf("🌏 set to listen on port: %v", serverPort)
	portInfo := fmt.Sprintf("set to listen on port: %v", serverPort)
	LogToFile(portInfo)

	httpSrv := http.NewServer(http.Address(":" + *listenAddr))

	httpSrv.HandlePrefix("/", router)

	app := kratos.New(
		kratos.Name("gin"), kratos.Server(httpSrv),
	)
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	if err := app.Run(); err != nil {
		Check(err, "could'nt run app")
		// log.Fatal(err)
	}
}
