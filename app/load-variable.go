package app

import (
	"github.com/draa-ayacucho/core/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func (a *App) loadVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	switch os.Getenv("GLOBAL_ENV") {
	case "develop":
		a.variables = *config.NewVariableConfig(config.Environment{})
	default:
		log.Fatal("variable GLOBAL_ENV not set")
	}
}
