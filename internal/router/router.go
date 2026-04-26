package router

import (
	"github.com/gin-gonic/gin"
	hello "github.com/irfan-official/simple_gin_template/internal/features/hello"
)

type RouteRegistrar func(rg *gin.RouterGroup)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")

	//  register your router here
	v1Routes := []RouteRegistrar{
		hello.RegisterRoutesV1,
	}

	for _, register := range v1Routes {
		register(v1)
	}

	return r
}