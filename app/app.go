package app

import (
	"context"
	"github.com/draa-ayacucho/core/config"
	"github.com/draa-ayacucho/core/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	router    *gin.Engine
	variables config.VariableConfig
	newLogger logger.NewLogger
	loader    *Loader
	//storage   func()
	//route     func()
}

// NewApp Create a new App instance
func NewApp(l *Loader) *App {
	return &App{
		router:    gin.Default(),
		variables: config.VariableConfig{},
		newLogger: logger.NewLogger{},
		loader:    l,
		//storage:   ia.Storage,
		//route:     ia.Route,
	}
}

// Variable Get app variables information
func (a *App) Variable() config.VariableConfig {
	return a.variables
}

// Logger Get logger information of the app
func (a *App) Logger() logger.NewLogger {
	return a.newLogger
}

func (a *App) loadLogger() {
	a.newLogger.Init(a.variables)
}

func (a *App) loadRoute() {
	for _, l := range a.loader.GinRouteLoader {
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

func (a *App) loadStorage() {}

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
