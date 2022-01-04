package app

import (
	"context"
	"github.com/draa-ayacucho/core/cmd/app/logger"
	"github.com/draa-ayacucho/core/configs"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	router    *gin.Engine
	variables configs.VariableConfig
	newLogger logger.NewLogger
	loader    *Loader
}

// NewApp Create a new App instance
func NewApp(l *Loader) *App {
	return &App{
		variables: configs.VariableConfig{},
		newLogger: logger.NewLogger{},
		loader:    l,
	}
}

// Variable Get app variables information
func (a *App) Variable() configs.VariableConfig {
	return a.variables
}

// Logger Get logger information of the app
func (a *App) Logger() logger.NewLogger {
	return a.newLogger
}

func (a *App) loadLogger() {
	a.newLogger.Init(a.variables)
}

func (a *App) loadStorage() {}

func (a *App) loadRoute() {
	a.router = gin.New()
	for _, l := range a.loader.GinRoute {
		switch l.Method {
		case http.MethodGet:
			a.router.GET(l.Path, l.Handler)
		case http.MethodPost:
			a.router.POST(l.Path, l.Handler)
		default:
			a.newLogger.ErrorLogger.Printf("method %s dont allowed\n", l.Method)
		}
	}
}

// Initialize Sets the initial configuration for the app
func (a *App) Initialize() {
	a.loadVariable()
	a.loadLogger()
	a.loadStorage()
	a.loadRoute()
}

// Run Starts the http server
func (a *App) Run() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	ctxCancel, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-signalChan
		a.newLogger.ErrorLogger.Printf("system call: %v\n", osCall)
		cancel()
	}()

	if err := a.loadServer(ctxCancel); err != nil {
		a.newLogger.ErrorLogger.Printf("failed to serve: %v\n", err)
	}
}
