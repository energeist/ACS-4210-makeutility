package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/energeist/tournament-calculator/helpers"
	"github.com/energeist/tournament-calculator/models"
)

// Handler struct for GORM database
type Handler struct {
	db *gorm.DB
}

// Main function
func main() {
	fmt.Println("Hello, World!")

	ginPort := helpers.LoadFromDotEnv("GIN_PORT")
	fmt.Println("Server Port: ", ginPort)

	// initialize GORM and connect to SQLite database withs test.db file
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema, create Test table with fields if it doesn't exist
	db.AutoMigrate(&models.Test{}, &models.Player{}, &models.GameMap{})

	handler := newHandler(db)

	r := gin.New()

	// Define test routes

	r.GET("/ping", pingHandler)
	r.GET("/test", handler.listTestHandler)
	r.POST("/test", handler.createTestHandler)
	r.DELETE("/test/:id", handler.deleteTestHandler)

	// Define player routes
	r.GET("/player", handler.listPlayerHandler)
	r.GET("/player/:id", handler.listPlayerHandler)
	r.POST("/player", handler.createPlayerHandler)

	// Define map routes
	r.GET("/map", handler.listMapHandler)
	r.POST("/map", handler.createMapHandler)

	r.Run(":" + ginPort) // listen and serve on port specified in .env file
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

// Define a ping route for testing
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Define CRUD handlers for Test struct as methods of Handler struct
func (h *Handler) listTestHandler(c *gin.Context) {
	var tests []models.Test

	if result := h.db.Find(&tests); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, tests)
}

func (h *Handler) createTestHandler(c *gin.Context) {
	var test models.Test

	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Create(&test); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &test)
}

func (h *Handler) deleteTestHandler(c *gin.Context) {
	id := c.Param("id")

	if result := h.db.Delete(&models.Test{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "record deleted",
	// })

	c.Status(http.StatusNoContent)
}

// Define CRUD handlers for Player struct as methods of Handler struct
func (h *Handler) listPlayerHandler(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		var player models.Player

		if result := h.db.First(&player, id); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(200, player)
		return
	} else {
		var players []models.Player

		if result := h.db.Find(&players); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(200, players)
		return
	}
}

func (h *Handler) createPlayerHandler(c *gin.Context) {
	var player models.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Create(&player); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &player)
}

// Define CRUD handlers for Map struct as methods of Handler struct
func (h *Handler) listMapHandler(c *gin.Context) {
	var gameMaps []models.GameMap

	if result := h.db.Find(&gameMaps); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gameMaps)
}

func (h *Handler) createMapHandler(c *gin.Context) {
	var mapObj models.GameMap

	if err := c.ShouldBindJSON(&mapObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Create(&mapObj); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &mapObj)
}
