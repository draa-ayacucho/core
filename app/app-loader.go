package app

import "github.com/gin-gonic/gin"

type Loader struct {
	GinRoute []GinRoute
	Storage  []Storage
}

type GinRoute struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
	Role    []string
}

type Storage struct{}
