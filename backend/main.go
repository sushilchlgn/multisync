package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"multisync-backend/browser"
	"multisync-backend/commands"
	"multisync-backend/sessions"
)

func main() {

	r := gin.Default()
	dispatcher := commands.NewDispatcher()

	// health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// create session
	r.POST("/sessions", func(c *gin.Context) {

		var req struct {
			Name   string `json:"name"`
			Device string `json:"device"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		session := sessions.CreateSession(req.Name, req.Device)
		c.JSON(200, session)
	})

	// list sessions
	r.GET("/sessions", func(c *gin.Context) {
		c.JSON(200, sessions.GetSessions())
	})

	// delete session
	r.DELETE("/sessions/:id", func(c *gin.Context) {
		id := c.Param("id")
		sessions.DeleteSession(id)

		browser.Stop(id)

		c.JSON(200, gin.H{"deleted": id})
	})

	// START browser (Phase 4.3 core)
	r.POST("/sessions/:id/start", func(c *gin.Context) {

		id := c.Param("id")

		instance, err := browser.Start(id)
		if err != nil {
			c.JSON(500, gin.H{
				"error":   "failed to start browser",
				"details": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"sessionId": instance.SessionID,
			"status":    "running",
			"ready":     true,
		})
	})

	// STOP browser
	r.POST("/sessions/:id/stop", func(c *gin.Context) {

		id := c.Param("id")
		browser.Stop(id)

		c.JSON(200, gin.H{
			"stopped": id,
		})
	})

	// GET browser instance
	r.GET("/sessions/:id/browser", func(c *gin.Context) {

		id := c.Param("id")

		instance, exists := browser.Get(id)
		if !exists {
			c.JSON(404, gin.H{"error": "browser not running"})
			return
		}

		c.JSON(200, gin.H{
			"sessionId": instance.SessionID,
			"status":    "running",
		})
	})

	r.POST("/command", func(c *gin.Context) {

	var cmd commands.Command

	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := dispatcher.Dispatch(cmd)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "executed",
	})
})

	r.Run(":8080")
}
