package main

import (
	"fmt"
	"log"

	"todoapp-backend/internal/db"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Panicln("Error connecting to database:", err)
	}
	log.Println("Connected to database")

	defer func() {
		db.Disconnect()
		log.Println("Disconnected from database")
	}()

	if err := db.Ping(); err != nil {
		log.Panicln("Unable to ping database:", err)
	}

	fmt.Println("Hello, World!")
}
