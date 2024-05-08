// Package helpers provides utility functions for the tournament-calculator project.
package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
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
