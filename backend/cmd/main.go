package main

import (
	"fmt"
	"log"
	"os"

	"todoapp-backend/internal/logger"
	"todoapp-backend/internal/model"
	"todoapp-backend/internal/repository"

	"github.com/joho/godotenv"
)

// Get database configuration from .env file.
func loadEnv() *repository.Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	dbConfig := &repository.Config{}

	dbConfig.URI = os.Getenv("POSTGRES_URI")
	if dbConfig.URI == "" {
		log.Panicln("\"POSTGRES_URI\" not set in .env file")
	}

	return dbConfig
}

func main() {
	logger.Setup()
	defer logger.Cleanup()

	dbConfig := loadEnv()
	log.Println("Loaded .env file")

	if err := repository.Connect(dbConfig); err != nil {
		log.Panicln("Error connecting to database:", err)
	}
	log.Println("Connected to database")

	defer func() {
		repository.Disconnect()
		log.Println("Disconnected from database")
	}()

	if err := repository.Ping(); err != nil {
		log.Panicln("Unable to ping database:", err)
	}

	todoRepository := repository.NewTodoRepository()

	todos, err := todoRepository.GetAllTodos()
	if err != nil {
		log.Panicln("Error getting todos:", err)
	}

	fmt.Println(todos)

	_, err = todoRepository.NewTodo(&model.Todo{Content: fmt.Sprintf("Task %v", len(todos)+1)})
	if err != nil {
		log.Panicln("Error creating todo:", err)
	}

	todos, err = todoRepository.GetAllTodos()
	if err != nil {
		log.Panicln("Error getting todos:", err)
	}

	fmt.Println(todos)

	_, err = todoRepository.UpdateTodo(&model.Todo{ID: len(todos), Content: "Updated"})
	if err != nil {
		log.Panicln("Error updating todo:", err)
	}

	todos, err = todoRepository.GetAllTodos()
	if err != nil {
		log.Panicln("Error getting todos:", err)
	}

	fmt.Println(todos)

	err = todoRepository.DeleteTodoByID(len(todos))
	if err != nil {
		log.Panicln("Error deleting todo:", err)
	}

	todos, err = todoRepository.GetAllTodos()
	if err != nil {
		log.Panicln("Error getting todos:", err)
	}

	fmt.Println(todos)

	fmt.Println("Hello, World!")
}
