package utils

import (
	"math/rand"

	"github.com/energeist/tournament-calculator/models"
)

// split out the matchmaking logic from the client here
func SelectTwoRandomPlayers(players []models.Player, maps []models.GameMap) (player1, player2 models.Player) {

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

	player1 = players[player1Index]
	player2 = players[player2Index]

	return player1, player2
}
