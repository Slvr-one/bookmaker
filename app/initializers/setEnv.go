package inits

import (
	"fmt"
	"os"

	h "github.com/Slvr-one/bookmaker/app/handlers"
)

func SetEnv(host string, mPort string, sPort string) (MH string, MP string, SP string) {
	mongoHost, set := os.LookupEnv("mongoHost")
	if !set {
		msg := fmt.Sprintf("mongoHost wasn't set, default is %s", host)
		h.LogToFile(msg)
		mongoHost = host
	}

	mongoPort, set := os.LookupEnv("mongoPort")
	if !set {
		msg := fmt.Sprintf("mongoPort wasn't set, default is %s", mPort)
		h.LogToFile(msg)
		mongoPort = mPort
	}

	serverPort, set := os.LookupEnv("serverPort")
	if !set {
		msg := fmt.Sprintf("serverPort env wasn't set, default is %s", mPort)
		h.LogToFile(msg)
		serverPort = sPort
	}

	return mongoHost, mongoPort, serverPort
}
