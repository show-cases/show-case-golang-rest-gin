package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPingRoute(rg *gin.RouterGroup) {
	// simple health check
	ping := rg.Group("/ping")
	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
