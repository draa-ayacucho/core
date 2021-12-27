package app

type IAppLoader interface {
	Logger()
	Route()
	Storage()
	Variable()
}
