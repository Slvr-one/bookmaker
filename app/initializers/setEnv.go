package inits

import (
	"fmt"
	"os"

	h "github.com/Slvr-one/bookmaker/handlers"
)

func SetEnv(host string, port string) (MH string, MP string) {
	mongoHost, set := os.LookupEnv("mongoHost")
	if !set {
		msg := fmt.Sprintf("mongoHost wasn't set, default is %s", host)
		h.LogToFile(msg)
		mongoHost = host
	}

	mongoPort, set := os.LookupEnv("mongoPort")
	if !set {
		msg := fmt.Sprintf("mongoPort wasn't set, default is %s", port)

		h.LogToFile(msg)
		mongoPort = port
	}

	return mongoHost, mongoPort
}
