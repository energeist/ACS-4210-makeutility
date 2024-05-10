package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/energeist/tournament-calculator/server/db"
	"github.com/energeist/tournament-calculator/server/handlers"
)

func SetupRoutes(r *gin.Engine, handler *db.Handler) {
	pingHandlers := &handlers.PingHandler{Handler: handler}
	playerHandler := &handlers.PlayerHandler{Handler: handler}
	mapHandler := &handlers.MapHandler{Handler: handler}
	matchHandler := &handlers.MatchHandler{Handler: handler}
	resultHandler := &handlers.ResultHandler{Handler: handler}

	// Ping route
	r.GET("api/v1/ping", pingHandlers.PingHandler)
	// Player routes
	r.GET("api/v1/player", playerHandler.ListPlayers)
	r.GET("api/v1/player/:id", playerHandler.ListPlayers)
	r.POST("api/v1/player", playerHandler.CreatePlayers)

	// Map routes
	r.GET("api/v1/gameMap", mapHandler.ListMaps)
	r.GET("api/v1/gameMap/:id", mapHandler.ListMaps)
	r.POST("api/v1/gameMap", mapHandler.CreateMaps)

	// Match routes
	r.GET("api/v1/match", matchHandler.ListMatches)
	r.GET("api/v1/match/:id", matchHandler.ListMatches)
	r.POST("api/v1/match", matchHandler.CreateMatches)

	// Result routes
	r.GET("api/v1/result", resultHandler.ListResults)
	r.GET("api/v1/result/:id", resultHandler.ListResults)
	r.POST("api/v1/result", resultHandler.CreateResults)
}
