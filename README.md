# ACS-4210 - TOURNAMENT-CALCULATOR

[![Go Report Card](https://goreportcard.com/badge/github.com/energeist/tournament-calculator)](https://goreportcard.com/report/energeist/tournament-calculator)

Tournament Calculator is a simple utility that can be used to predict results of matches between players in a game.  Currently, it is configured to generate possible outcomes between professional Starcraft II players using publicly available data from the [Team Liquid](https://liquipedia.net/starcraft2/Main_Page) wiki and [Aligulac's](aligulac.com) player performance database.

This repository includes both a server and a client (in this case, a driver script).  The server uses `gin-gonic` for api routing and `Gorm` for data persistence in a local `SQLite` database file.

## USING THIS REPOSITORY

### BUILD

- Clone this repository and `cd` into the root directory
- `go build -o server ./server; go build -o client ./client;`
- Ensure you have a `.env` file in your project root with format mirroring that of `sample.env`
- - Server port can be anything that works for your system; `gin-gonic`'s standard port is `8080`.
- - Generate an API key for Aligulac queries from [Aligulac's API docs](http://aligulac.com/about/api/) if you don't already have one

### RUN

You'll need two terminal tabs running, one for the server and the other for the client.  The server will create a `tournament.db` file and seed it with `GameMap` and `Player` data the first time that it's run.

- After building the `server` and `client` executables, from the repository root run:
- - `./server/server` to start the database server running in the background
- - `./client/client` to perform a randomized match calculation using two players from Aligulac's top 50
- - `go test -v -bench=. ./...` to run all table and bench tests

## API ROUTES

### PING - `/api/v1/ping/`
- `GET` `/ping` - returns a `pong` message to ensure that the server responding

### PLAYERS - `/api/v1/player/`
  -`GET` `/api/v1/player` - retrieves a list of all `player`s
  -`GET` `/api/v1/player/:id` - retrieves a specific `player` by id
  -`POST` `/api/v1/player` - creates a `player` object according to the shape of the `Player` struct, which has a `Rating` struct nested inside:

```
type Player struct {
	Base
	Tag    string `json:"tag" gorm:"type:varchar(15)"`
	Race   string `json:"race" gorm:"type:varchar(1)"`
	Rating Rating `json:"current_rating" gorm:"embedded;embeddedPrefix:rating_"`
}

type Rating struct {
	CurrentRating float32 `json:"rating" gorm:"column:rating_rating"`
	VsProtoss     float32 `json:"tot_vp" gorm:"column:rating_vs_protoss"`
	VsTerran      float32 `json:"tot_vt" gorm:"column:rating_vs_terran"`
	VsZerg        float32 `json:"tot_vz" gorm:"column:rating_vs_zerg"`
}

type Base struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
```

### MAPS - `/api/v1/gameMap/`
  -`GET` `/api/v1/gameMap` - retrieves a list of all `gameMap`s
  -`GET` `/api/v1/gameMap/:id` - retrieves a specific `gameMap` by id
  -`POST` `/api/v1/gameMap` - creates a `gameMap` object according to the shape of the `GameMap` struct:

```
type GameMap struct {
	Base
	Name         string  `json:"name" gorm:"type:varchar(100)"`
	Height       uint8   `json:"height" gorm:"type:uint"`
	Width        uint8   `json:"width" gorm:"type:uint"`
	RushDistance uint8   `json:"rush_distance" gorm:"type:uint"`
	TvZ          float32 `json:"tvz" gorm:"type:float"`
	ZvP          float32 `json:"zvp" gorm:"type:float"`
	PvT          float32 `json:"pvt" gorm:"type:float"`
}
```

### MATCHES - `/api/v1/match/`
  -`GET` `/api/v1/match` - retrieves a list of all `match`s
  -`GET` `/api/v1/match/:id` - retrieves a specific `match` by id
  -`POST` `/api/v1/match` - creates a `match` object according to the shape of the `Match` struct:

```
type Match struct {
	Base
	Player1ID uint   `json:"player1_id"`
	Player1   Player `json:"player1" gorm:"foreignKey:Player1ID"`
	Player2ID uint   `json:"player2_id"`
	Player2   Player `json:"player2" gorm:"foreignKey:Player2ID"` 
}
```

### RESULTS - `/api/v1/result/`
  -`GET` `/api/v1/result` - retrieves a list of all `result`s
  -`GET` `/api/v1/result/:id` - retrieves a specific `result` by id
  -`POST` `/api/v1/result` - creates a `result` object according to the shape of the `Result` struct:

```
type Result struct {
	Base
	MatchID          uint    `json:"match_id"`
	WinnerID         uint    `json:"winner_id"`
	Winner           Player  `json:"winner" gorm:"foreignKey:WinnerID"`
	WinnerPercentage float32 `json:"winner_percentage"`
	LoserID          uint    `json:"loser_id"`
	Loser            Player  `json:"loser" gorm:"foreignKey:LoserID"`
}
```

## PROJECT DIRECTORY STRUCTURE

```
\tournament-calculator
|--\client
|----client.go
|----client
|--\helpers
|----helpers.go
|----map_pool.go
|----helpers_test.go
|--\models
|----models.go
|--\server
|----\db
|------handler.go
|------tournament.db
|----\handlers
|------map_handlers.go
|------match_handlers.go
|------ping_handlers.go
|------player_handlers.go
|------result_handlers.go
|----\routes
|------routes.go
|----server.go
|----server
|--\utils
|----matchmaker.go
|----simulation.go
|----simulation_test.go
|--.env
|--.gitignore
|--go.mod
|--go.sum
|--proposal.md
|--README.md
|--sample.env
