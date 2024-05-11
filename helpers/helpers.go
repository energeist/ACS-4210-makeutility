// Package helpers provides utility functions for the tournament-calculator project.
package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
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
	_, b, _, _ := runtime.Caller(0)    // Get the path to the current file
	basePath := filepath.Dir(b)        // Get the directory of the current file
	rootPath := filepath.Dir(basePath) // Assume the root is one level up from the file

	dotenvPath := filepath.Join(rootPath, ".env")
	if err := godotenv.Load(dotenvPath); err != nil {
		fmt.Println("No .env file found at", dotenvPath)
		return ""
	}

	return os.Getenv(key)
}

func AligulacURL(endpoint, APIKey string, id int) string {
	return "https://api.aligulac.com/api/v1/" + endpoint + "/" + fmt.Sprint(id) + "/?apikey=" + APIKey
}

func ServerURL(endpoint, serverPort string) string {
	return "http://localhost:" + serverPort + "/api/v1/" + endpoint
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

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("server returned error status: %d", resp.StatusCode)
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

	if resp.StatusCode >= 300 { // or use any other appropriate criteria
		return resp, fmt.Errorf("server returned error status: %d", resp.StatusCode)
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

	for _, player := range APIResponsePlayers.Objects {
		if result := db.Create(&player); result.Error != nil {
			fmt.Println("Error creating player: ", result.Error)
		}
	}
}
