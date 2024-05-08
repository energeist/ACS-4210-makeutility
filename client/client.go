package main

import (
	"fmt"

	"github.com/energeist/tournament-calculator/helpers"
)

const rating_factor = 1354

func main() {
	fmt.Println("Hello, World!")

	serverPort := helpers.LoadFromDotEnv("GIN_PORT")
	APIKey := helpers.LoadFromDotEnv("ALIGULAC_API_KEY")

	fmt.Println("Server Port: ", serverPort)

	// Get top 50 players from Aligulac API
	url := "http://aligulac.com/api/v1/player/?current_rating__isnull=false&order_by=-current_rating__rating&limit=50&apikey=" + APIKey
	body, err := helpers.GetRequest(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(body))
}
