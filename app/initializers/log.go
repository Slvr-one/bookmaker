package inits

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ichtrojan/thoth"
	"github.com/rs/zerolog"
)

type logger struct {
	GinCtx *gin.Context
}

func (_this *logger) Set(key string, value interface{}) {
	if _this.GinCtx != nil {
		_this.GinCtx.Set(key, value)
	}
}

func (_this *logger) Get(key string) interface{} {
	if _this.GinCtx != nil {
		value, _ := _this.GinCtx.Get(key)
		return value
	}
	return nil
}

func InitLog() {
	// var l logger
	// l.Set("one",)
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// log.Print("hello world")

	backLogger, _ := thoth.Init("backlog")
	backLogger.Log(errors.New("This is the back Log: \n"))
}
