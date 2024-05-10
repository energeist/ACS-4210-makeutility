package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/energeist/tournament-calculator/server/db"
)

type PingHandler struct {
	Handler *db.Handler
}

// Define a ping route for testing
func (h *PingHandler) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
