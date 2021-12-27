package app

import (
	"github.com/draa-ayacucho/core/config"
	"github.com/draa-ayacucho/core/logger"
)

// Variable Get app variables information
func (a *App) Variable() config.VariableConfig {
	return a.variables
}

// Logger Get logger information of the app
func (a *App) Logger() logger.NewLogger {
	return a.newLogger
}
