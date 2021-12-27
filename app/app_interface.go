package app

type IAppLoader interface {
	// Route Loads the routes on the app
	Route()

	// Storage Loads the type of storage on the app
	Storage()
}
