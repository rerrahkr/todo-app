package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Database configurations.
type config struct {
	URI string
}

// Get database configuration from .env file.
func loadEnv() *config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	dbConfig := &config{}

	dbConfig.URI = os.Getenv("POSTGRES_URI")
	if dbConfig.URI == "" {
		log.Panicln("\"POSTGRES_URI\" not set in .env file")
	}

	return dbConfig
}

var dbConfig *config

func init() {
	dbConfig = loadEnv()
	log.Println("Loaded .env file")
}

// DB pool.
var pool *pgxpool.Pool

// Connect to the database.
func Connect() (err error) {
	pool, err = pgxpool.New(context.Background(), dbConfig.URI)
	return
}

// Ping the database.
func Ping() error {
	return pool.Ping(context.Background())
}

// Disconnect from the database.
func Disconnect() {
	pool.Close()
}
