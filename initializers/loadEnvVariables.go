package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

func LoadEnvVariables() {
	if FileExists(".env") {
		// Load .env file if it exists
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
