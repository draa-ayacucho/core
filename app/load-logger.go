package app

import "github.com/draa-ayacucho/core/logger"

func (a *App) loadLogger() {
	a.newLogger = &logger.NewLogger{}
	a.newLogger.Init(a.variables)
}
