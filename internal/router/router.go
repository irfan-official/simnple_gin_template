package router

import (
	"net/http"
	"github.com/irfan-official/simnple_gin_template/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	healthHandler := handler.NewHealthHandler()

	api_v1 := r.Group("/api/v1")

	apiV1.GET("/health", healthHandler.Health)

	return r
}