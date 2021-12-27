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
	router    *mux.Router
	variables config.VariableConfig
	newLogger *logger.NewLogger
}

// NewApp Create a new instance of the App
func NewApp() *App {
	return &App{
		router:    mux.NewRouter(),
		newLogger: &logger.NewLogger{},
	}
}

// Storage Loads the type of storage on the app
func (a *App) Storage() {}

// Route Loads the routes on the app
func (a *App) Route() {}

// Initialize Sets the initial configuration for the app
func (a *App) Initialize() {
	a.Variable()
	a.Logger()
	a.Storage()
	a.Route()
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

	if err := a.Server(ctxCancel); err != nil {
		a.newLogger.ErrorLogger.Printf("failed to serve: %v\n", err)
	}
}
