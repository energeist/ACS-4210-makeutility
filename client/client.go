package main

import (
	"encoding/json"
	"fmt"

	"github.com/energeist/tournament-calculator/helpers"
	"github.com/energeist/tournament-calculator/models"
)

const rating_factor = 1354

func main() {
	fmt.Println("Hello, World!")

	serverPort := helpers.LoadFromDotEnv("GIN_PORT")
	APIKey := helpers.LoadFromDotEnv("ALIGULAC_API_KEY")

	fmt.Println("Server Port: ", serverPort)

	// Get top 50 players from Aligulac API
	url := "http://aligulac.com/api/v1/player/?current_rating__isnull=false&order_by=-current_rating__rating&limit=50&apikey=" + APIKey
	multiplePlayersData, err := helpers.GetRequest(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Raw API response with mutliple players: ")
	fmt.Println(string(multiplePlayersData))

	var APIResponsePlayers APIResponsePlayers
	if err := json.Unmarshal(multiplePlayersData, &APIResponsePlayers); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Parsed API response with multiple players: ")
	fmt.Println(APIResponsePlayers)
}

type APIResponsePlayers struct {
	Objects []models.Player `json:"objects"`
}
