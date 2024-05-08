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
	resultHandler := &handlers.ResultHandler{Handler: handler}

	// Player routes
	r.GET("/player", playerHandler.ListPlayers)
	r.GET("/player/:id", playerHandler.ListPlayers)
	r.POST("/player", playerHandler.CreatePlayers)

	// Map routes
	r.GET("/gameMap", mapHandler.ListMaps)
	r.GET("/gameMap/:id", mapHandler.ListMaps)
	r.POST("/gameMap", mapHandler.CreateMaps)

	// Match routes
	r.GET("/match", matchHandler.ListMatches)
	r.GET("/match/:id", matchHandler.ListMatches)
	r.POST("/match", matchHandler.CreateMatches)

	// Result routes
	r.GET("/result", resultHandler.ListResults)
	r.GET("/result/:id", resultHandler.ListResults)
	r.POST("/result", resultHandler.CreateResults)
}
