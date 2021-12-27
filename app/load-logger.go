package app

func (a *App) loadLogger() {
	a.newLogger.Init(a.variables)
}
