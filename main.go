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
	Stuff  string `json:"name"`
	Number int    `json:"number"`
}

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
