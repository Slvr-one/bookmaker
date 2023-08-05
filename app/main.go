package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Slvr-one/bookmaker/api"
	"github.com/Slvr-one/bookmaker/db"
	h "github.com/Slvr-one/bookmaker/handlers"
	inits "github.com/Slvr-one/bookmaker/initializers"
	s "github.com/Slvr-one/bookmaker/structs"
	kratos "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DefaultServerPort = "9090"      //default port to serve app
	DefaultMongoPort  = "27017"     // default port for mongoDB connection
	DefaultHost       = "localhost" //default host for mongo database sever
)

var (
	Horses    []s.Horse
	MainBoard s.Board

	Conn    s.Conn
	connErr error

	// id        int
	// item      string
	// completed int
	// view     = template.Must(template.ParseFiles("./views/index.html"))
	// db = SqlDB()
)

func init() {
	// rand.Seed(time.Now().UnixNano())
	start := time.Now()
	inits.InitLog()

	inits.SetBoard(MainBoard, start)
	inits.LoadEnvVars()

	defer h.End(start)

}

// main
func main() {
	exampleData := s.Car{
		Id:        primitive.NewObjectID(),
		CreatedAt: time.Now().UTC(),
		Brand:     "Mercedes",
		Model:     "G-360",
		Year:      2002,
	}
	// db.Insert(exampleData)

	router := api.InitR()

	mongoHost, mongoPort, serverPort := inits.SetEnv(DefaultMongoPort, DefaultHost, DefaultServerPort)
	mongodbUrl := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	Conn.Client, connErr = db.MongoDB(mongodbUrl, MainBoard, Conn) // SqlDB()
	h.Check(connErr, "err on running mongoDB func to init connect")

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
