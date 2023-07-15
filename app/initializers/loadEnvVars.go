package inits

import (
	"github.com/Slvr-one/bookmaker/handlers"
	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	envLoadErr := godotenv.Load()
	handlers.Check(envLoadErr, "No .env file found")
}
