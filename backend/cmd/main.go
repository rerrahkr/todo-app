package main

import (
	"fmt"
	"log"

	"todoapp-backend/internal/db"
	"todoapp-backend/internal/logger"
)

func main() {
	logger.Setup()
	defer logger.Cleanup()

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

	todos, err := db.GetAllTodos()
	if err != nil {
		log.Panicln("Error getting todos:", err)
	}

	fmt.Println(todos)

	_, err = db.NewTodo(fmt.Sprintf("Task %v", len(todos)+1))
	if err != nil {
		log.Panicln("Error creating todo:", err)
	}

	todos, err = db.GetAllTodos()
	if err != nil {
		log.Panicln("Error getting todos:", err)
	}

	fmt.Println(todos)

	_, err = db.UpdateTodo(len(todos), "Updated")
	if err != nil {
		log.Panicln("Error updating todo:", err)
	}

	todos, err = db.GetAllTodos()
	if err != nil {
		log.Panicln("Error getting todos:", err)
	}

	fmt.Println(todos)

	err = db.DeleteTodo(len(todos))
	if err != nil {
		log.Panicln("Error deleting todo:", err)
	}

	todos, err = db.GetAllTodos()
	if err != nil {
		log.Panicln("Error getting todos:", err)
	}

	fmt.Println(todos)

	fmt.Println("Hello, World!")
}
