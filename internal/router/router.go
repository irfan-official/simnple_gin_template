package router

import (
	"github.com/irfan-official/simnple_gin_template/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	healthHandler := handler.NewHealthHandler()

	r.GET("/health", healthHandler.Health)

	return r
}