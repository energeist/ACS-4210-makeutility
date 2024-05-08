package handlers

import (
	"net/http"

	"github.com/energeist/tournament-calculator/models"
	"github.com/gin-gonic/gin"

	"github.com/energeist/tournament-calculator/server/db"
)

type MatchHandler struct {
	Handler *db.Handler
}

// Define CRUD handlers for Match struct as methods of Handler struct
func (h *MatchHandler) ListMatches(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		var match models.Match

		if result := h.Handler.DB.First(&match, id); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(200, match)
		return
	} else {
		var matches []models.Match

		if result := h.Handler.DB.Find(&matches); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(200, matches)
		return
	}
}

func (h *MatchHandler) CreateMatches(c *gin.Context) {
	var match models.Match

	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.Handler.DB.Where(models.Match{ID: match.ID}).FirstOrCreate(&match)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, &match)
}
