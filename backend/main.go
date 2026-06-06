package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	go handleMessages()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		handleConnections(c.Writer, c.Request)
	})

	r.Run(":8080")
}