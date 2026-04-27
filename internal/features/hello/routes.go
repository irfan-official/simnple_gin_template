package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesV1(v1 *gin.RouterGroup) {
	Hello := v1.Group("/hello")

	Hello.GET("/", SimpleHelloHandler )

	Hello.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Developed by Irfan",
		})
	})
}