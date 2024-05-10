package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/energeist/tournament-calculator/helpers"
	"github.com/energeist/tournament-calculator/models"
)

// const rating_factor = 1354

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
	winner, predictedChance := calculateOutcome(players[player1Index], players[player2Index], maps)
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

func calculateOutcome(player1, player2 models.Player, maps []models.GameMap) (winner models.Player, resultProbability float32) {
	numIterations := 10000000
	player1Race := player1.Race
	player2Race := player2.Race

	// initialize a map of outcomes
	// key: map name, value: number of wins for player 1
	outcomesMap := make(map[string]uint)

	// var matchUp string
	var mapBalance float32
	// var matchBalance string

	// loop 1000 times
	for i := 0; i < numIterations; i++ {
		// pick a random map from the pool
		randomMap := maps[rand.Intn(len(maps))]

		if player1Race == "P" && player2Race == "T" {
			// matchUp = "PvT"
			mapBalance = randomMap.PvT
		} else if player1Race == "T" && player2Race == "Z" {
			// matchUp = "TvZ"
			mapBalance = randomMap.TvZ
		} else if player1Race == "Z" && player2Race == "P" {
			// matchUp = "ZvP"
			mapBalance = randomMap.ZvP
		} else {
			mapBalance = 0.5
		}

		// fmt.Println("Matchup: ", matchUp, "Map: ", randomMap.Name, "Balance: ", mapBalance)

		// calculate win probability for player1
		if rand.Float32() <= mapBalance {
			outcomesMap[randomMap.Name]++
		}
		outcomesMap["total"]++
	}

	fmt.Println(outcomesMap)

	// calculate win probability for player1

	var player1Wins uint
	for mapName, wins := range outcomesMap {
		if mapName != "total" {
			player1Wins += wins
		}
	}
	aggregateWinProbability := float32(player1Wins) / float32(outcomesMap["total"])

	if aggregateWinProbability > 0.5 {
		fmt.Printf("Aggregate win probability: %.5f\n", aggregateWinProbability)
		resultProbability = aggregateWinProbability * 100
		return player1, resultProbability
	} else {
		fmt.Printf("Aggregate win probability: %.5f\n", aggregateWinProbability)
		resultProbability = (1 - aggregateWinProbability) * 100
		return player2, resultProbability
	}
}

// func coinflip(player1, player2 models.Player) models.Player {
// 	// randomly select a player to win

// 	if rand.Float64() < 0.5 {
// 		return player1
// 	}
// 	return player2
// }
