package app

import (
	"github.com/draa-ayacucho/core/config"
	"github.com/gin-gonic/gin"
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
		gin.SetMode(gin.DebugMode)
		a.variables = *config.NewVariableConfig(config.Environment{})
	case "production":
		gin.SetMode(gin.ReleaseMode)
		a.variables = *config.NewVariableConfig(config.Environment{})
	default:
		log.Fatal("variable GLOBAL_ENV not set")
	}
}
