package handlers

import (
	"net/http"

	"github.com/energeist/tournament-calculator/models"
	"github.com/gin-gonic/gin"

	"github.com/energeist/tournament-calculator/server/db"
)

type PlayerHandler struct {
	Handler *db.Handler
}

// Define CRUD handlers for Player struct as methods of Handler struct
func (h *PlayerHandler) ListPlayers(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		var player models.Player

		if result := h.Handler.DB.First(&player, id); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(200, player)
		return
	} else {
		var players []models.Player

		if result := h.Handler.DB.Find(&players); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(200, players)
		return
	}
}

func (h *PlayerHandler) CreatePlayers(c *gin.Context) {
	var player models.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.Handler.DB.Create(&player); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &player)
}
