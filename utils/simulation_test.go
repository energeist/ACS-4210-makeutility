package utils

import (
	"math/rand"
	"testing"
	"time"

	"github.com/energeist/tournament-calculator/models"
)

func BenchmarkSimulateWithMapBalance(b *testing.B) {
	// Seed the random number generator for reproducibility in tests
	rand.Seed(time.Now().UnixNano())
	numIterations := 1

	// Create mock players
	player1 := models.Player{ID: 1, Tag: "Player1", Race: "P"}
	player2 := models.Player{ID: 2, Tag: "Player2", Race: "T"}

	// Create mock game maps
	maps := []models.GameMap{
		{Name: "Map1", PvT: 60, TvZ: 50, ZvP: 45},
		{Name: "Map2", PvT: 55, TvZ: 55, ZvP: 50},
		{Name: "Map3", PvT: 65, TvZ: 45, ZvP: 55},
	}

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		// Call the function to benchmark
		SimulateWithMapBalance(player1, player2, maps, numIterations)
	}
}
