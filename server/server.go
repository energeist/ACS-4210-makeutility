package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/energeist/tournament-calculator/helpers"
	"github.com/energeist/tournament-calculator/models"
	"github.com/energeist/tournament-calculator/server/db"
	"github.com/energeist/tournament-calculator/server/routes"
)

// Main function
func main() {
	serverPort := helpers.LoadFromDotEnv("GIN_PORT")
	APIKey := helpers.LoadFromDotEnv("ALIGULAC_API_KEY")

	ginPort := helpers.LoadFromDotEnv("GIN_PORT")
	fmt.Println("Server Port: ", ginPort)

	// initialize GORM and connect to SQLite database withs tournament.db file
	database, err := gorm.Open(sqlite.Open("tournament.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema, create tables if they do not exist
	database.AutoMigrate(&models.Player{}, &models.GameMap{}, &models.Match{}, &models.Result{})

	// Seed database with maps if required
	var count int64
	database.Model(&models.GameMap{}).Count(&count)
	if count < 9 {
		helpers.GenerateDBSeed(database)
	} else {
		fmt.Println("Maps already seeded")
	}

	// Seed database with players if required
	// Get top X players from Aligulac API
	topXPlayers := 50

	database.Model(&models.Player{}).Count(&count)
	if count < int64(topXPlayers) {
		helpers.SeedTopPlayers(database, serverPort, APIKey, topXPlayers)
	} else {
		fmt.Println("Top " + strconv.Itoa(topXPlayers) + " players already seeded")
	}
	// helpers.SeedTopPlayers(serverPort, APIKey, topXPlayers)

	handler := db.NewHandler(database)

	r := gin.New()

	routes.SetupRoutes(r, handler)
	r.Run(":" + ginPort) // listen and serve on port specified in .env file
}
