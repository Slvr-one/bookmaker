package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// init a route handler for the server client
func InitR() *mux.Router {
	router := mux.NewRouter()
	// fs := http.FileServer(http.Dir("static"))

	// router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// router.Handle("/assets/", http.StripPrefix("/assets/", staticHandler()).Methods("GET")
	// router.HandleFunc("/GH/{name}", GetHorse).Methods("GET")       //Get a specific horse

	router.HandleFunc("/", Welcom).Methods("GET")
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/metrics", Monitor).Methods("GET")

	router.HandleFunc("/api/horses", GetHorses).Methods("GET") //List all available horses
	router.HandleFunc("/api/horses", CreateHorse).Methods("POST")
	router.HandleFunc("/api/horses/{name}", updateHorse).Methods("GET") //Get a specific horse
	router.HandleFunc("/api/horses/{name}", updateHorse).Methods("PUT") //Update a specific horse
	router.HandleFunc("/api/horses/{name}", DeleteHorse).Methods("DELETE")

	router.HandleFunc("/invest/{investor}/{horse}/{amount}", Invest).Methods("GET")

	// router.Handle("/", http.FileServer(http.Dir("templates/styles/")))
	// router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("templates/styles"))))

	styles := http.FileServer(http.Dir("./templates/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	// router.Handle("/styles/", styles)
	return router
}
