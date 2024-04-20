package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// initialize GORM database globally
var db *gorm.DB

// Define structs
type Test struct {
	ID     string `json:"id"`
	Stuff  string `json:"name"`
	Number int    `json:"number"`
}

func main() {
	fmt.Println("Hello, World!")

	// initialize GORM and connect to SQLite database withs test.db file
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema, create Test table with fields if it doesn't exist
	db.AutoMigrate(&Test{})

	r := gin.New()

	// Define routes

	r.GET("/ping", pingHandler)
	r.GET("/test", listTestHandler)
	r.POST("/test", createTestHandler)
	r.DELETE("/test/:id", deleteTestHandler)

	r.Run() // listen and serve on port 8080
}

// Define route handler functions

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func listTestHandler(c *gin.Context) {
	var tests []Test

	if result := db.Find(&tests); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, tests)
}

func createTestHandler(c *gin.Context) {
	var test Test

	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := db.Create(&test); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &test)
}

func deleteTestHandler(c *gin.Context) {
	id := c.Param("id")

	if result := db.Delete(&Test{}, id); result.Error != nil {
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
