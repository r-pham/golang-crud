package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Important to begin function name with capital to be used in other files
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
