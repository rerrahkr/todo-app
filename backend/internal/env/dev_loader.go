//go:build !prod

package env

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	log.Println("Loaded .env file")
}
