package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Define structs

// Handler struct
type Handler struct {
	db *gorm.DB
}

// Test struct
type Test struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}

// Plauer struct

type Player struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Rating    int    `json:"rating"`
	VsProtoss int    `json:"vs_protoss"`
	VsTerran  int    `json:"vs_terran"`
	VsZerg    int    `json:"vs_zerg"`
}

// Map struct
type Map struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	RushDistance int    `json:"rush_distance"`
	Tvz          string `json:"tvz"`
	ZvP          string `json:"zvp"`
	PvT          string `json:"pvt"`
}

// Match struct
type Match struct {
	ID        string `json:"id"`
	Player1ID string `json:"player1_id"`
	Player2ID string `json:"player2_id"`
	MapID     string `json:"map_id"`
	Timestamp string `json:"timestamp"`
}

// Result struct
type Result struct {
	ID      string `json:"id"`
	MatchID string `json:"match_id"`
	Winner  string `json:"winner"`
	Loser   string `json:"loser"`
}

type Bracket struct {
	ID        string `json:"id"`
	Players   []Player
	Timestamp string `json:"timestamp"`
}

// TODO: ModelWeights struct to be incorporated later
type ModelWeights struct {
}

// Main function
func main() {
	fmt.Println("Hello, World!")

	// initialize GORM and connect to SQLite database withs test.db file
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema, create Test table with fields if it doesn't exist
	db.AutoMigrate(&Test{})

	handler := newHandler(db)

	r := gin.New()

	// Define routes

	r.GET("/ping", pingHandler)
	r.GET("/test", handler.listTestHandler)
	r.POST("/test", handler.createTestHandler)
	r.DELETE("/test/:id", handler.deleteTestHandler)

	r.Run() // listen and serve on port 8080
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
	var tests []Test

	if result := h.db.Find(&tests); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, tests)
}

func (h *Handler) createTestHandler(c *gin.Context) {
	var test Test

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

	if result := h.db.Delete(&Test{}, id); result.Error != nil {
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
