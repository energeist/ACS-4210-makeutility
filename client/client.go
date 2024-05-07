package main

import (
	"fmt"

	"github.com/energeist/tournament-calculator/helpers"
)

func main() {
	fmt.Println("Hello, World!")

	serverPort := helpers.LoadFromDotEnv("GIN_PORT")
	APIKey := helpers.LoadFromDotEnv("ALIGULAC_API_KEY")

	fmt.Println("Server Port: ", serverPort)
	fmt.Println("API Key: ", APIKey)

	// Create a new player
	// player := models.Player{
	// 	Name:      "Test Player",
	// 	Rating:    1000,
	// 	VsProtoss: 50,
	// 	VsTerran:  50,
	// 	VsZerg:    50,
	// }

	// playerJSON, err := json.Marshal(player)
	// if err != nil {
	// 	fmt.Println("Error marshaling player to JSON: ", err)
	// 	return
	// }

	// resp, err := http.Post("http://localhost:"+serverPort+"/player", "application/json", bytes.NewBuffer(playerJSON))
	// if err != nil {
	// 	fmt.Println("Error creating player: ", err)
	// }

	// fmt.Println("Response: ", resp)

	// Seed the db with maps
	// helpers.GenerateDBSeed()
}
