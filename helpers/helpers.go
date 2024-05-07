// Package helpers provides utility functions for the tournament-calculator project.
package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/energeist/tournament-calculator/models"
	"github.com/joho/godotenv"
)

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

func GetRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
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

func GenerateDBSeed() {
	alcyone := models.GameMap{
		Name:         "Alcyone",
		Height:       144,
		Width:        144,
		RushDistance: 36,
		TvZ:          51.5,
		ZvP:          56.1,
		PvT:          45.8,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	amphion := models.GameMap{
		Name:         "Amphion",
		Height:       140,
		Width:        140,
		RushDistance: 36,
		TvZ:          54.5,
		ZvP:          53.8,
		PvT:          36.0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	crimsonCourt := models.GameMap{
		Name:         "Crimson Court",
		Height:       124,
		Width:        148,
		RushDistance: 34,
		TvZ:          69.2,
		ZvP:          32.3,
		PvT:          42.5,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	dynasty := models.GameMap{
		Name:         "Dynasty",
		Height:       144,
		Width:        110,
		RushDistance: 33,
		TvZ:          80.0,
		ZvP:          39.3,
		PvT:          52.1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	ghostRiver := models.GameMap{
		Name:         "Ghost River",
		Height:       142,
		Width:        128,
		RushDistance: 31,
		TvZ:          53.8,
		ZvP:          51.2,
		PvT:          51.5,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	goldenaura := models.GameMap{
		Name:         "Goldenaura",
		Height:       140,
		Width:        140,
		RushDistance: 35,
		TvZ:          49.5,
		ZvP:          42.3,
		PvT:          48.3,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	oceanborn := models.GameMap{
		Name:         "Oceanborn",
		Height:       140,
		Width:        142,
		RushDistance: 34,
		TvZ:          54.1,
		ZvP:          52.5,
		PvT:          48.6,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	postYouth := models.GameMap{
		Name:         "Post-Youth",
		Height:       116,
		Width:        144,
		RushDistance: 30,
		TvZ:          45.5,
		ZvP:          40.0,
		PvT:          30.6,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	siteDelta := models.GameMap{
		Name:         "Site Delta",
		Height:       136,
		Width:        148,
		RushDistance: 35,
		TvZ:          45.3,
		ZvP:          50.8,
		PvT:          50.3,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	gameMaps := []models.GameMap{alcyone, amphion, crimsonCourt, dynasty, ghostRiver, goldenaura, oceanborn, postYouth, siteDelta}

	for _, gameMap := range gameMaps {
		resp, err := PostRequest(ServerURL("map", "8080"), gameMap)

		if err != nil {
			fmt.Println("Error creating map: ", err)
		}

		fmt.Println("Response: ", resp)
	}

}
