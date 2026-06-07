package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"multisync-backend/browser"
	"multisync-backend/sessions"
)

func main() {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.POST("/sessions", func(c *gin.Context) {

		var req struct {
			Name   string `json:"name"`
			Device string `json:"device"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		session := sessions.CreateSession(
			req.Name,
			req.Device,
		)

		c.JSON(200, session)
	})

	r.GET("/sessions", func(c *gin.Context) {
		c.JSON(200, sessions.GetSessions())
	})

	r.DELETE("/sessions/:id", func(c *gin.Context) {

		id := c.Param("id")

		sessions.DeleteSession(id)

		c.JSON(200, gin.H{
			"deleted": id,
		})
	})

	r.POST("/sessions/:id/start", func(c *gin.Context) {

		id := c.Param("id")

		instance := browser.Start(id)

		c.JSON(200, instance)
	})

	r.POST("/sessions/:id/stop", func(c *gin.Context) {

		id := c.Param("id")

		browser.Stop(id)

		c.JSON(200, gin.H{
			"stopped": id,
		})
	})

	r.GET("/sessions/:id/browser", func(c *gin.Context) {

		id := c.Param("id")

		instance, exists := browser.Get(id)

		if !exists {
			c.JSON(404, gin.H{
				"error": "browser not running",
			})
			return
		}

		c.JSON(200, instance)
	})

	r.Run(":8080")
}
