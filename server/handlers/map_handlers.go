package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/energeist/tournament-calculator/models"
	"github.com/energeist/tournament-calculator/server/db"
)

type MapHandler struct {
	Handler *db.Handler
}

// Define CRUD handlers for Map struct as methods of Handler struct
func (h *MapHandler) ListMaps(c *gin.Context) {
	var gameMaps []models.GameMap

	if result := h.Handler.DB.Find(&gameMaps); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gameMaps)
}

func (h *MapHandler) CreateMaps(c *gin.Context) {
	var mapObj models.GameMap

	if err := c.ShouldBindJSON(&mapObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.Handler.DB.Create(&mapObj); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &mapObj)
}
