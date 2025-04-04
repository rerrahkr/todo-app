package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Database configurations.
type Config struct {
	URI string
}

// DB pool.
var pool *pgxpool.Pool

// Connect to the database.
func Connect(config *Config) (err error) {
	pool, err = pgxpool.New(context.Background(), config.URI)
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
