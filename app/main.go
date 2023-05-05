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

// Init horses var as a slice of Horse struct (slice: an array with no fixed Size, Type)
var (
	horses    []Horse
	MainBoard Board

	srvPort = "9099" //to serve app

	host      = "localhost"
	mongoPort = "27017" // to connect to mongoDB

	mongodbUrl = fmt.Sprintf("mongodb://%s:%s", host, mongoPort)
)

func init() {
	// rand.Seed(time.Now().UnixNano())
	// SqlDB()
	// MongoDB()
}

// main
func main() {
	initLog()
	start := time.Now()
	defer End(start)

	//get port env var (if provided by user, default / 9090)
	serverPort, set := os.LookupEnv("serverPort")
	// serverPort := os.Getenv("serverPort")
	if !set {
		LogToFile("serverPort env was'nt set, default is 9090.")
		serverPort = srvPort
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

		// TLSConfig: &tls.Config{
		// 	ClientAuth: tls.RequireAndVerifyClientCert,
		// 	ClientCAs:  certPool,
		// 	MinVersion: tls.VersionTLS12,
		// 	CipherSuites: []uint16{
		// 		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		// 	}
		// 	PreferServerCipherSuites: true,
		// 	CurvePreferences: []tls.CurveID{
		// 		tls.CurveP256,
		// 		tls.X25519,
		// 		tls.CurveP384,

		// 	}
		// // InsecureSkipVerify: true,
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
