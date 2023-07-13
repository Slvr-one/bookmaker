package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Slvr-one/bookmaker/api"
	h "github.com/Slvr-one/bookmaker/handlers"
	s "github.com/Slvr-one/bookmaker/structs"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	DefaultServerPort = "9090"  //default port to serve app
	DefaultMongoPort  = "27017" // default port for mongoDB connection
	// defaultHost       = "localhost"
)

var (
	Horses    []s.Horse
	MainBoard s.Board

	conn = &s.Conn{}
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
	router := api.InitR()
	h.InitLog()
	start := time.Now()
	defer h.End(start)

	MainBoard.Date = &start

	serverPort, set := os.LookupEnv("serverPort")
	if !set {
		h.LogToFile("serverPort env wasn't set, default is 9090.")
		serverPort = DefaultServerPort
	}

	listenAddr := flag.String("listenaddr", serverPort, "port to serve the app")

	// Hardcoded data - @todo: add database
	Horses = append(Horses, s.Horse{Name: "Monahen Boy", Color: "brown", Record: &s.Record{Wins: 8, Losses: 3}})
	Horses = append(Horses, s.Horse{Name: "Dangerous", Color: "brown:white", Record: &s.Record{Wins: 7, Losses: 1}})
	Horses = append(Horses, s.Horse{Name: "Black Beauty", Color: "black", Record: &s.Record{Wins: 4, Losses: 5}})
	Horses = append(Horses, s.Horse{Name: "horse 4", Color: "black", Record: &s.Record{Wins: 4, Losses: 5}})

	log.Printf("🌏 set to listen on port: %v", serverPort)
	portInfo := fmt.Sprintf("set to listen on port: %v", serverPort)
	h.LogToFile(portInfo)

	httpSrv := http.NewServer(http.Address(":" + *listenAddr))

	httpSrv.HandlePrefix("/", router)

	app := kratos.New(
		kratos.Name("gin"), kratos.Server(httpSrv),
	)
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	if err := app.Run(); err != nil {
		h.Check(err, "app didnt start on $run")
		// log.Fatal(err)
	}
}
