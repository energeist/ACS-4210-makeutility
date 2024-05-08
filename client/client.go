package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/energeist/tournament-calculator/helpers"
	"github.com/energeist/tournament-calculator/models"
)

const rating_factor = 1354

func main() {
	serverPort := helpers.LoadFromDotEnv("GIN_PORT")
	APIKey := helpers.LoadFromDotEnv("ALIGULAC_API_KEY")

	// Get top X players from Aligulac API
	topXPlayers := 50

	seedTopPlayers(serverPort, APIKey, topXPlayers)

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

	// create a Match between the two players and store in db
	match := models.Match{
		Player1ID: players[player1Index].ID,
		Player2ID: players[player2Index].ID,
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

	// perform calculation
	// lots of iterations, randomly assign map from the pool to each iteration
	// calculation will yield a win probability for each player

	// result = calculateOutcome(players[player1Index], players[player2Index], maps)

	// query db for the last match entered

	// coinflip to get a random winner for now
	winner := coinflip(players[player1Index], players[player2Index])
	var loser models.Player
	if winner == players[player1Index] {
		loser = players[player2Index]
	} else {
		loser = players[player1Index]
	}

	// store Result in db
	result := models.Result{
		MatchID:  matchResponse.ID,
		WinnerID: winner.ID,
		LoserID:  loser.ID,
	}

	_, err = helpers.PostRequest(helpers.ServerURL("result", serverPort), result)
	if err != nil {
		fmt.Println("Error storing result:", err)
	}

	fmt.Printf("\nResult stored in db: %s wins\n", winner.Tag)
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

// func calculateOutcome(player1, player2 models.Player, maps []models.GameMap) models.Result {
// initialize a map of probabilities that each player wins
// calculate win probability for player1
// win probability for player2 is 1 - win probability for player1
// store the result in the db
// return the result
// }

func coinflip(player1, player2 models.Player) models.Player {
	// randomly select a player to win

	if rand.Float64() < 0.5 {
		return player1
	}
	return player2
}
