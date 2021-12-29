package app

import "github.com/gin-gonic/gin"

type Loader struct {
	GinRouteLoader []GinRouteLoader
	StorageLoader  StorageLoader
}

type GinRouteLoader struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	IsProtected bool
}

type StorageLoader struct{}
