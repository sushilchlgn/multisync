package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"service": "multisync-backend",
		})
	})

	// test URL endpoint
	r.GET("/render", func(c *gin.Context) {
		url := c.Query("url")

		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "url is required",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ready to render",
			"url":     url,
		})
	})

	r.Run(":8080")
}