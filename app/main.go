package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Slvr-one/bookmaker/app/api"
	"github.com/Slvr-one/bookmaker/app/db"
	h "github.com/Slvr-one/bookmaker/app/handlers"
	inits "github.com/Slvr-one/bookmaker/app/initializers"
	s "github.com/Slvr-one/bookmaker/app/structs"
	kratos "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	DefaultServerPort = "9090"      //default port to serve app
	DefaultMongoPort  = "27011"     // default port for mongoDB connection
	DefaultHost       = "localhost" //default host for mongo database sever
)

var (
	Horses    []s.Horse
	MainBoard s.Board

	Conn    s.Conn
	connErr error
	start   = time.Now()

	// id        int
	// item      string
	// completed int
	// view     = template.Must(template.ParseFiles("./views/index.html"))
	// db = SqlDB()
)

func init() {
	// rand.Seed(time.Now().UnixNano())
	inits.InitLog()
	inits.LoadEnvVars()
	inits.SetBoard(MainBoard, start)
	h.Populate()

}

// main
func main() {
	defer h.EndProgram(start)

	router := api.InitR()
	c := Conn.Client

	// exampleData := s.Car{
	// 	Id:        primitive.NewObjectID(),
	// 	CreatedAt: time.Now().UTC(),
	// 	Brand:     "Mercedes",
	// 	Model:     "G-360",
	// 	Year:      2002,
	// }
	// db.Insert(exampleData)

	mongoHost, mongoPort, serverPort := inits.SetEnv(DefaultHost, DefaultMongoPort, DefaultServerPort)
	mongodbUrl := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	c, connErr = db.MongoDB(mongodbUrl, MainBoard, Conn) // SqlDB()
	h.Check(connErr, "err on running mongoDB func to init connect")

	listenAddr := flag.String("listenaddr", serverPort, "port to serve the app")

	portInfo := fmt.Sprintf("🌏 set to listen on port: %v", serverPort)
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
