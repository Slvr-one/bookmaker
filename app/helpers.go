package main

import (
	"errors"
	"fmt"
	"time"

	// clog "log"
	"math/rand"

	"github.com/fatih/color"
	"github.com/ichtrojan/thoth"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func initLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// log.Print("hello world")

	backLogger, _ := thoth.Init("backlog")
	backLogger.Log(errors.New("This is the back Log: \n"))
}
func Check(err error, msg string) {
	if err != nil {
		color.Green(msg)
		LogToFile(msg)
		log.Info().Msg(msg)
		log.Fatal()

		// clog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		// clog.Fatalf("Error: %v\n Error Message: %v", err, msg)
		// panic(err.Error())
		// os.Exit(-1)
	}
}

func LogToFile(msg string) {
	// init a general file for logging
	genLogger, _ := thoth.Init("log")
	genLogger.Log(errors.New(msg))
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func RandomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Great to see you, %v!",
		"Look who it is! its %v",
		"Hi, %v. Welcome!",
		"Hi %v, knew you'd come back",
		"%v, You again..",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	i := rand.Intn(len(formats))
	return formats[i]
}

func End(start time.Time) {
	// log to file <end session> msg
	t := time.Now()
	elapsed := t.Sub(start)
	msg := fmt.Sprintf("ended at -> %s, elapsed -> %s", t, elapsed)
	LogToFile(msg)
}
