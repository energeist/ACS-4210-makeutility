package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/energeist/tournament-calculator/helpers"
	"github.com/energeist/tournament-calculator/models"
)

const rating_factor = 1354

func main() {
	fmt.Println("Hello, World!")

	serverPort := helpers.LoadFromDotEnv("GIN_PORT")
	APIKey := helpers.LoadFromDotEnv("ALIGULAC_API_KEY")

	// Get top X players from Aligulac API
	topXPlayers := 50

	existingPlayersData, err := helpers.GetRequest(helpers.ServerURL("player", serverPort))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var existingPlayers []models.Player
	if err := json.Unmarshal(existingPlayersData, &existingPlayers); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if len(existingPlayers) < topXPlayers {
		url := "http://aligulac.com/api/v1/player/?current_rating__isnull=false&order_by=-current_rating__rating&limit=" + strconv.Itoa(topXPlayers) + "&apikey=" + APIKey

		multiplePlayersData, err := helpers.GetRequest(url)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		var APIResponsePlayers APIResponsePlayers
		if err := json.Unmarshal(multiplePlayersData, &APIResponsePlayers); err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println("APIResponsePlayers.Objects[0]: ")
		fmt.Println(APIResponsePlayers.Objects[0])

		// Store each player in the database
		for _, player := range APIResponsePlayers.Objects {
			_, err := helpers.PostRequest(helpers.ServerURL("player", serverPort), player)
			if err != nil {
				fmt.Println("Error storing player:", err)
			}
		}
	} else {
		fmt.Println("Top " + strconv.Itoa(topXPlayers) + " players already seeded")
	}
}

type APIResponsePlayers struct {
	Objects []models.Player `json:"objects"`
}
