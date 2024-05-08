package handlers

import (
	"net/http"

	"github.com/energeist/tournament-calculator/models"
	"github.com/gin-gonic/gin"

	"github.com/energeist/tournament-calculator/server/db"
)

type ResultHandler struct {
	Handler *db.Handler
}

// Define CRUD handlers for Result struct as methods of Handler struct
func (h *ResultHandler) ListResults(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		var result models.Result
		if result := h.Handler.DB.Preload("Winner").Preload("Loser").First(&result, id); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}
		c.JSON(200, result)
	} else {
		var results []models.Result
		if result := h.Handler.DB.Preload("Winner").Preload("Loser").Find(&results); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}
		c.JSON(200, results)
	}
}

func (h *ResultHandler) CreateResults(c *gin.Context) {
	var resultObj models.Result

	if err := c.ShouldBindJSON(&resultObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.Handler.DB.Create(&resultObj); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &resultObj)
}
