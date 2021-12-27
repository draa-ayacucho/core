package app

func (a *App) Logger() {
	a.newLogger.Init(a.variables)
}
