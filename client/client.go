package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

	seedTopPlayers(serverPort, APIKey, topXPlayers)

	// Get all players from db
	playersData, err := helpers.GetRequest(helpers.ServerURL("player", serverPort))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var players []models.Player
	if err := json.Unmarshal(playersData, &players); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Get all maps from db
	mapsData, err := helpers.GetRequest(helpers.ServerURL("map", serverPort))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var maps []models.GameMap
	if err := json.Unmarshal(mapsData, &maps); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// select 2 random players from the list
	player1Index := rand.Intn(len(players) - 1)

	// ensure that two unique players are selected
	var player2Index int
	for {
		player2Index = rand.Intn(len(players) - 1)
		if player2Index != player1Index {
			break
		}
	}

	// create a Match between the two players and store in db
	match := models.Match{
		Player1: players[player1Index],
		Player2: players[player2Index],
	}

	_, err = helpers.PostRequest(helpers.ServerURL("match", serverPort), match)
	if err != nil {
		fmt.Println("Error storing match:", err)
	}
	// perform calculation
	// lots of iterations, randomly assign map from the pool to each iteration
	// calculation will yield a win probability for each player

	// store Result in db
}

type APIResponsePlayers struct {
	Objects []models.Player `json:"objects"`
}

func seedTopPlayers(serverPort, APIKey string, topXPlayers int) {
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
