package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/energeist/tournament-calculator/server/db"
	"github.com/energeist/tournament-calculator/server/handlers"
)

func SetupRoutes(r *gin.Engine, handler *db.Handler) {
	playerHandler := &handlers.PlayerHandler{Handler: handler}
	mapHandler := &handlers.MapHandler{Handler: handler}
	matchHandler := &handlers.MatchHandler{Handler: handler}

	// Player routes
	r.GET("/player", playerHandler.ListPlayers)
	r.POST("/player", playerHandler.CreatePlayers)

	// Map routes
	r.GET("/gameMap", mapHandler.ListMaps)
	r.POST("/gameMap", mapHandler.CreateMaps)

	// Match routes
	r.GET("/match", matchHandler.ListMatches)
	r.POST("/match", matchHandler.CreateMatches)
}
