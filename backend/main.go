// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	go handleMessages()

// 	r := gin.Default()

// 	r.GET("/health", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"status": "ok"})
// 	})

// 	r.GET("/ws", func(c *gin.Context) {
// 		handleConnections(c.Writer, c.Request)
// 	})

// 	r.Run(":8080")
// }

package main

import (
	"net/http"

	// "sessions"

	"github.com/gin-gonic/gin"
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

	r.Run(":8080")
}