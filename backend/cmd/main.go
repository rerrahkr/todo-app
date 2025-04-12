package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"todoapp-backend/internal/controller"
	"todoapp-backend/internal/logger"
	"todoapp-backend/internal/repository"
	"todoapp-backend/internal/router"
	"todoapp-backend/internal/service"

	"github.com/joho/godotenv"
)

type APIConfig struct {
	Port int
}

// Get database configuration from .env file.
func loadEnv() (*repository.Config, *APIConfig) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	dbConfig := &repository.Config{}

	dbConfig.URI = os.Getenv("POSTGRES_URI")
	if dbConfig.URI == "" {
		log.Panicln(`"POSTGRES_URI" not set in .env file`)
	}

	apiConfig := &APIConfig{}

	apiConfig.Port, err = strconv.Atoi(os.Getenv("BACKEND_PORT"))
	if err != nil {
		log.Panicln(`"BACKEND_PORT" not set in .env file`)
	}

	return dbConfig, apiConfig
}

func main() {
	logger.Setup()
	defer logger.Cleanup()

	dbConfig, apiConfig := loadEnv()
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

	repository := repository.NewTodoRepository()
	service := service.NewTodoService(repository)
	controller := controller.NewTodoController(service)

	router := router.NewTodoRouter(controller)
	handler := router.SetupRoutes()

	http.ListenAndServe(fmt.Sprintf(":%d", apiConfig.Port), handler)
}
