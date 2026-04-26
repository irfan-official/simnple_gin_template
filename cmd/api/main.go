package main

import (
	"log"
	"net/http"
	"github.com/irfan-official/simnple_gin_template/internal/config"
	"github.com/irfan-official/simnple_gin_template/internal/router"
)

func main() {
	cfg := config.Load()

	r := router.NewRouter()

	log.Printf("Server running on %s\n", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
           "status": true,
		   "message": "Server is runing"
		})
	})
}