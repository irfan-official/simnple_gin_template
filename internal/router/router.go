package router

import (
	"github.com/irfan-official/simple_gin_template/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	healthHandler := handler.NewHealthHandler()

	apiV1 := r.Group("/api/v1")

	apiV1.GET("/health", healthHandler.Health)

	return r
}