package models

import "time"

// Base struct for all models
type Base struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

// Player struct
type Player struct {
	Base
	Tag    string `json:"tag" gorm:"type:varchar(15)"`
	Race   string `json:"race" gorm:"type:varchar(1)"`
	Rating Rating `json:"current_rating" gorm:"embedded;embeddedPrefix:rating_"`
}

// Rating struct to capture nested JSON data
type Rating struct {
	CurrentRating float32 `json:"rating" gorm:"column:rating_rating"`
	VsProtoss     float32 `json:"tot_vp" gorm:"column:rating_vs_protoss"`
	VsTerran      float32 `json:"tot_vt" gorm:"column:rating_vs_terran"`
	VsZerg        float32 `json:"tot_vz" gorm:"column:rating_vs_zerg"`
}

// Map struct
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

// Match struct
type Match struct {
	Base
	Player1ID uint   `json:"player1_id"`
	Player1   Player `json:"player1" gorm:"foreignKey:Player1ID"`
	Player2ID uint   `json:"player2_id"`
	Player2   Player `json:"player2" gorm:"foreignKey:Player2ID"`
}

// Result struct
type Result struct {
	Base
	MatchID          uint    `json:"match_id"`
	WinnerID         uint    `json:"winner_id"`
	Winner           Player  `json:"winner" gorm:"foreignKey:WinnerID"`
	WinnerPercentage float32 `json:"winner_percentage"`
	LoserID          uint    `json:"loser_id"`
	Loser            Player  `json:"loser" gorm:"foreignKey:LoserID"`
}

// TODO: Expand functionality to calculate whole BO16 bracket, beyond single match
type Bracket struct {
	Base
	Players   []Player `json:"players" gorm:"type:json"`
	Timestamp string   `json:"timestamp" gorm:"type:datetime"`
}

// TODO: ModelWeights struct to be incorporated later
type ModelWeights struct {
}
