package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func LoadFromDotEnv(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	return os.Getenv(key)
}
