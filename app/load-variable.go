package app

import (
	"github.com/draa-ayacucho/core/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func (a *App) Variable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	switch os.Getenv("GLOBAL_ENV") {
	case "develop":
		a.variables = *config.NewVariableConfig(config.Environment{})
	case "production":
		// The config.AwsSecretsManager{} (not implemented jet)
		// a.variables = *config.NewVariableConfig(config.AwsSecretsManager{})
		fallthrough
	default:
		log.Fatal("variable GLOBAL_ENV not set")
	}
}
