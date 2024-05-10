// Package helpers provides utility functions for the tournament-calculator project.
package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/energeist/tournament-calculator/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type APIResponsePlayers struct {
	Objects []models.Player `json:"objects"`
}

// Load environment variables from .env file
func LoadFromDotEnv(key string) string {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("No .env file found")
	}

	return os.Getenv(key)
}

func AligulacURL(endpoint, APIKey string, id int) string {
	return "https://api.aligulac.com/api/v1/" + endpoint + "/" + fmt.Sprint(id) + "/?apikey=" + APIKey
}

func ServerURL(endpoint, serverPort string) string {
	return "http://localhost:" + serverPort + "/" + endpoint
}

func GetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // Read the response body
	if err != nil {
		return nil, err
	}

	return body, nil
}

func PostRequest(url string, data interface{}) (*http.Response, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dataJSON))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GenerateDBSeed(db *gorm.DB) {
	fmt.Println("Seeding DB with maps")

	gameMaps := MapPool()

	for _, gameMap := range gameMaps {
		if result := db.Create(&gameMap); result.Error != nil {
			fmt.Println("Error creating map: ", result.Error)
		}
	}
}

func SeedTopPlayers(db *gorm.DB, serverPort, APIKey string, topXPlayers int) {
	// existingPlayersData, err := GetRequest(ServerURL("player", serverPort))
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	// var existingPlayers []models.Player
	// if err := json.Unmarshal(existingPlayersData, &existingPlayers); err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	// if len(existingPlayers) < topXPlayers {
	// 	fmt.Println("Seeding top " + strconv.Itoa(topXPlayers) + " players")
	fmt.Println("Calling Aligulac API")

	url := "http://aligulac.com/api/v1/player/?current_rating__isnull=false&order_by=-current_rating__rating&limit=" + strconv.Itoa(topXPlayers) + "&apikey=" + APIKey

	multiplePlayersData, err := GetRequest(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var APIResponsePlayers APIResponsePlayers
	if err := json.Unmarshal(multiplePlayersData, &APIResponsePlayers); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Store each player in the database
	// for _, player := range APIResponsePlayers.Objects {
	// 	_, err := PostRequest(ServerURL("player", serverPort), player)
	// 	if err != nil {
	// 		fmt.Println("Error storing player:", err)
	// 	}
	// }

	for _, player := range APIResponsePlayers.Objects {
		if result := db.Create(&player); result.Error != nil {
			fmt.Println("Error creating player: ", result.Error)
		}
	}
	// } else {
	// 	fmt.Println("Top " + strconv.Itoa(topXPlayers) + " players already seeded")
	// }
}
