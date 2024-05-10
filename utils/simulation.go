package utils

import (
	"math/rand"

	"github.com/energeist/tournament-calculator/models"
)

func SimulateWithMapBalance(player1, player2 models.Player, maps []models.GameMap, numIterations int) (winner models.Player, resultProbability float32) {
	player1Race := player1.Race
	player2Race := player2.Race

	// initialize a map of outcomes
	// key: map name, value: number of wins for player 1
	outcomesMap := make(map[string]int)

	// var matchUp string
	var mapBalance float32

	for i := 0; i < numIterations; i++ {
		// pick a random map from the pool
		randomMap := maps[rand.Intn(len(maps))]

		if player1Race == "P" && player2Race == "T" {
			// matchUp = "PvT"
			mapBalance = randomMap.PvT / 100
		} else if player1Race == "T" && player2Race == "P" {
			// matchUp = "TvP"
			mapBalance = 1 - (randomMap.PvT / 100)
		} else if player1Race == "T" && player2Race == "Z" {
			// matchUp = "TvZ"
			mapBalance = randomMap.TvZ / 100
		} else if player1Race == "Z" && player2Race == "T" {
			// matchUp = "ZvT"
			mapBalance = 1 - (randomMap.TvZ / 100)
		} else if player1Race == "Z" && player2Race == "P" {
			// matchUp = "ZvP"
			mapBalance = randomMap.ZvP / 100
		} else if player1Race == "P" && player2Race == "Z" {
			// matchUp = "PvZ"
			mapBalance = 1 - (randomMap.ZvP / 100)
		} else {
			mapBalance = 0.5
		}

		// calculate win probability for player1
		if rand.Float32() <= mapBalance {
			outcomesMap[randomMap.Name]++
		}
		outcomesMap["total"]++
	}

	// calculate win probability for player1

	var player1Wins int
	for mapName, wins := range outcomesMap {
		if mapName != "total" {
			player1Wins += wins
		}
	}
	aggregateWinProbability := float32(player1Wins) / float32(outcomesMap["total"])

	if aggregateWinProbability > 0.5 {
		resultProbability = aggregateWinProbability * 100
		return player1, resultProbability
	} else {
		resultProbability = (1 - aggregateWinProbability) * 100
		return player2, resultProbability
	}
}
