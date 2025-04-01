package main

import (
	"fmt"
	"log"

	"todoapp-backend/internal/logger"
	"todoapp-backend/internal/repository"
)

func main() {
	logger.Setup()
	defer logger.Cleanup()

	if err := repository.Connect(); err != nil {
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

	_, err = todoRepository.NewTodo(fmt.Sprintf("Task %v", len(todos)+1))
	if err != nil {
		log.Panicln("Error creating todo:", err)
	}

	todos, err = todoRepository.GetAllTodos()
	if err != nil {
		log.Panicln("Error getting todos:", err)
	}

	fmt.Println(todos)

	_, err = todoRepository.UpdateTodo(len(todos), "Updated")
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
