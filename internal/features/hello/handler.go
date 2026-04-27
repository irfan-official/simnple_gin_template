package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SimpleHelloHandler(c *gin.Context) {

	data, err := SimpleHelloService("Hello")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": data,
	})
}
