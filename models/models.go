package models

// Test struct
type Test struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}

// Player struct

type Player struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Rating    int    `json:"rating"`
	VsProtoss int    `json:"vs_protoss"`
	VsTerran  int    `json:"vs_terran"`
	VsZerg    int    `json:"vs_zerg"`
}

// Map struct
type GameMap struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	RushDistance int    `json:"rush_distance"`
	Tvz          string `json:"tvz"`
	ZvP          string `json:"zvp"`
	PvT          string `json:"pvt"`
}

// Match struct
type Match struct {
	ID        string `json:"id"`
	Player1ID string `json:"player1_id"`
	Player2ID string `json:"player2_id"`
	MapID     string `json:"map_id"`
	Timestamp string `json:"timestamp"`
}

// Result struct
type Result struct {
	ID      string `json:"id"`
	MatchID string `json:"match_id"`
	Winner  string `json:"winner"`
	Loser   string `json:"loser"`
}

type Bracket struct {
	ID        string `json:"id"`
	Players   []Player
	Timestamp string `json:"timestamp"`
}

// TODO: ModelWeights struct to be incorporated later
type ModelWeights struct {
}
