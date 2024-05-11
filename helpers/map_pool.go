package helpers

import (
	"github.com/energeist/tournament-calculator/models"
)

func MapPool() []models.GameMap {
	alcyone := models.GameMap{
		Name:         "Alcyone",
		Height:       144,
		Width:        144,
		RushDistance: 36,
		TvZ:          51.5,
		ZvP:          56.1,
		PvT:          45.8,
	}

	amphion := models.GameMap{
		Name:         "Amphion",
		Height:       140,
		Width:        140,
		RushDistance: 36,
		TvZ:          54.5,
		ZvP:          53.8,
		PvT:          36.0,
	}

	crimsonCourt := models.GameMap{
		Name:         "Crimson Court",
		Height:       124,
		Width:        148,
		RushDistance: 34,
		TvZ:          69.2,
		ZvP:          32.3,
		PvT:          42.5,
	}

	dynasty := models.GameMap{
		Name:         "Dynasty",
		Height:       144,
		Width:        110,
		RushDistance: 33,
		TvZ:          80.0,
		ZvP:          39.3,
		PvT:          52.1,
	}

	ghostRiver := models.GameMap{
		Name:         "Ghost River",
		Height:       142,
		Width:        128,
		RushDistance: 31,
		TvZ:          53.8,
		ZvP:          51.2,
		PvT:          51.5,
	}

	goldenaura := models.GameMap{
		Name:         "Goldenaura",
		Height:       140,
		Width:        140,
		RushDistance: 35,
		TvZ:          49.5,
		ZvP:          42.3,
		PvT:          48.3,
	}

	oceanborn := models.GameMap{
		Name:         "Oceanborn",
		Height:       140,
		Width:        142,
		RushDistance: 34,
		TvZ:          54.1,
		ZvP:          52.5,
		PvT:          48.6,
	}

	postYouth := models.GameMap{
		Name:         "Post-Youth",
		Height:       116,
		Width:        144,
		RushDistance: 30,
		TvZ:          45.5,
		ZvP:          40.0,
		PvT:          30.6,
	}

	siteDelta := models.GameMap{
		Name:         "Site Delta",
		Height:       136,
		Width:        148,
		RushDistance: 35,
		TvZ:          45.3,
		ZvP:          50.8,
		PvT:          50.3,
	}

	gameMaps := []models.GameMap{alcyone, amphion, crimsonCourt, dynasty, ghostRiver, goldenaura, oceanborn, postYouth, siteDelta}

	return gameMaps
}
