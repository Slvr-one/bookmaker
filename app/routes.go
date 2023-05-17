package main

import (
	"context"
	"fmt"

	// "log"

	// "net/http"

	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"

	// "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	// "github.com/go-kratos/kratos/v2/transport/http"
)

// init a route handler for the server client
func InitR() *gin.Engine {
	router := gin.Default()
	// fs := http.FileServer(http.Dir("static"))

	// Use Kratos middleware
	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))

	router.GET("/hello/:name", Welc)

	// router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// router.Handle("/assets/", http.StripPrefix("/assets/", staticHandler()).Methods("GET")
	router.GET("/GH/:name", GetHorse) //Get a specific horse

	router.GET("/", Welcom)
	router.GET("/health", Health)
	router.GET("/metrics", Monitor)

	router.GET("/api/horses", GetHorses) //List all available horses
	router.POST("/api/horses", CreateHorse)
	router.GET("/api/:name", GetHorses)          //Get a specific horse
	router.PUT("/api/horses/:name", UpdateHorse) //Update a specific horse
	router.DELETE("/api/horses/:name", DeleteHorse)

	// router.GET("/invest/:investor/:horse/:amount", Invest)

	return router
}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

func Welc(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "error" {
		// return kratos error
		kgin.Error(ctx, errors.Unauthorized("auth_error", "no authentication"))
	} else {
		ctx.JSON(200, map[string]string{"welcome": name})
	}
}

// using mux router -> (old)
// func InitR() *mux.Router {
// 	router := mux.NewRouter()
// 	// fs := http.FileServer(http.Dir("static"))

// 	// router.Handle("/assets/", http.StripPrefix("/assets/", fs))
// 	// router.Handle("/assets/", http.StripPrefix("/assets/", staticHandler()).Methods("GET")
// 	// router.HandleFunc("/GH/{name}", GetHorse).Methods("GET")       //Get a specific horse

// 	router.HandleFunc("/", Welcom).Methods("GET")
// 	router.HandleFunc("/health", Health).Methods("GET")
// 	router.HandleFunc("/metrics", Monitor).Methods("GET")

// 	router.HandleFunc("/api/horses", GetHorses).Methods("GET") //List all available horses
// 	router.HandleFunc("/api/horses", CreateHorse).Methods("POST")
// 	router.HandleFunc("/api/horses/{name}", updateHorse).Methods("GET") //Get a specific horse
// 	router.HandleFunc("/api/horses/{name}", updateHorse).Methods("PUT") //Update a specific horse
// 	router.HandleFunc("/api/horses/{name}", DeleteHorse).Methods("DELETE")

// 	router.HandleFunc("/invest/{investor}/{horse}/{amount}", Invest).Methods("GET")

// 	// router.Handle("/", http.FileServer(http.Dir("templates/styles/")))
// 	// router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("templates/styles"))))

// 	styles := http.FileServer(http.Dir("./templates/styles"))
// 	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
// 	// router.Handle("/styles/", styles)
// 	return router
// }
