package models

import "time"

// Test struct
type Test struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Name   string `json:"name" gorm:"type:varchar(100)"`
	Number int    `json:"number" gorm:"type:int"`
}

// Player struct

type Player struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	Rating    uint16    `json:"rating" gorm:"type:uint"`
	VsProtoss uint16    `json:"vs_protoss" gorm:"type:uint"`
	VsTerran  uint16    `json:"vs_terran" gorm:"type:uint"`
	VsZerg    uint16    `json:"vs_zerg" gorm:"type:uint"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

// Map struct
type GameMap struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name" gorm:"type:varchar(100)"`
	Height       uint8     `json:"height" gorm:"type:uint"`
	Width        uint8     `json:"width" gorm:"type:uint"`
	RushDistance uint8     `json:"rush_distance" gorm:"type:uint"`
	TvZ          float32   `json:"tvz" gorm:"type:float"`
	ZvP          float32   `json:"zvp" gorm:"type:float"`
	PvT          float32   `json:"pvt" gorm:"type:float"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:datetime"`
}

// Match struct
type Match struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Player1   Player    `json:"player1_id" gorm:"type:uint"`
	Player2   Player    `json:"player2_id" gorm:"type:uint"`
	GameMap   GameMap   `json:"map_id" gorm:"type:uint"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

// Result struct
type Result struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Match     Match     `json:"match_id" gorm:"type:int"`
	Winner    Player    `json:"winner" gorm:"foreignKey:WinnerID"`
	Loser     Player    `json:"loser" gorm:"foreignKey:LoserID"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

type Bracket struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Players   []Player  `json:"players" gorm:"type:json"`
	Timestamp string    `json:"timestamp" gorm:"type:datetime"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

// TODO: ModelWeights struct to be incorporated later
type ModelWeights struct {
}
