package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/energeist/tournament-calculator/helpers"
	"github.com/energeist/tournament-calculator/models"
	"github.com/energeist/tournament-calculator/utils"
)

// const rating_factor = 1354
const numIterations = 10000000

func main() {
	serverPort := helpers.LoadFromDotEnv("GIN_PORT")
	// APIKey := helpers.LoadFromDotEnv("ALIGULAC_API_KEY")

	fmt.Println("Gathering players from database")

	// Get all players from db
	playersData, err := helpers.GetRequest(helpers.ServerURL("player", serverPort))
	if err != nil {
		fmt.Println("Error in playersData request: ", err)
		return
	}

	var players []models.Player
	if err := json.Unmarshal(playersData, &players); err != nil {
		fmt.Println("Error in unmarshalling players data: ", err)
		return
	}

	fmt.Println("Gathering maps from database")

	// Get all maps from db
	mapsData, err := helpers.GetRequest(helpers.ServerURL("gameMap", serverPort))
	if err != nil {
		fmt.Println("Error in mapsData get request: ", err)
		return
	}

	var maps []models.GameMap
	if err := json.Unmarshal(mapsData, &maps); err != nil {
		fmt.Println("Error in unmarshalling maps data: ", err)
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

	player1 := players[player1Index]
	player2 := players[player2Index]

	fmt.Printf("Initializing match between %s (%s) and %s (%s)\n", player1.Tag, player1.Race, player2.Tag, player2.Race)

	// create a Match between the two players and store in db
	match := models.Match{
		Player1ID: player1.ID,
		Player2ID: player2.ID,
	}

	var lastMatch *http.Response
	lastMatch, err = helpers.PostRequest(helpers.ServerURL("match", serverPort), match)
	if err != nil {
		fmt.Println("Error storing match:", err)
	}

	bodyBytes, _ := io.ReadAll(lastMatch.Body)

	var matchResponse models.Match
	err = json.Unmarshal(bodyBytes, &matchResponse)
	if err != nil {
		fmt.Println("Error unmarshalling matchResponse:", err)
	}

	fmt.Printf("Match between %s and %s stored in db\n", players[player1Index].Tag, players[player2Index].Tag)

	fmt.Println("Calculating outcome of match...")
	winner, predictedChance := utils.SimulateWithMapBalance(players[player1Index], players[player2Index], maps, numIterations)
	// query db for the last match entered

	var loser models.Player
	if winner == players[player1Index] {
		loser = players[player2Index]
	} else {
		loser = players[player1Index]
	}

	// store Result in db
	result := models.Result{
		MatchID:          matchResponse.ID,
		WinnerID:         winner.ID,
		LoserID:          loser.ID,
		WinnerPercentage: predictedChance,
	}

	_, err = helpers.PostRequest(helpers.ServerURL("result", serverPort), result)
	if err != nil {
		fmt.Println("Error storing result:", err)
	}

	fmt.Printf("\nResult stored in db: %.5f%% predicted win for %s\n", predictedChance, winner.Tag)
}
