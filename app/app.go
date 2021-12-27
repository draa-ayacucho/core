package app

import (
	"context"
	"github.com/draa-ayacucho/core/config"
	"github.com/draa-ayacucho/core/logger"
	"github.com/gorilla/mux"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Router    *mux.Router
	variables config.VariableConfig
	newLogger logger.NewLogger
	storage   func()
	route     func()
}

// NewApp Create a new App instance
func NewApp(ia IAppLoader) *App {
	return &App{
		Router:    mux.NewRouter(),
		variables: config.VariableConfig{},
		newLogger: logger.NewLogger{},
		storage:   ia.Storage,
		route:     ia.Route,
	}
}

// Initialize Sets the initial configuration for the app
func (a *App) Initialize() {
	a.loadVariable()
	a.loadLogger()
	a.storage()
	a.route()
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
