package inits

import (
	"errors"

	"github.com/ichtrojan/thoth"
	"github.com/rs/zerolog"
)

func InitLog() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// log.Print("hello world")

	backLogger, _ := thoth.Init("backlog")
	backLogger.Log(errors.New("This is the back Log: \n"))
}
