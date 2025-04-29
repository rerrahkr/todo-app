package main

import (
	"fmt"
	"log"
	"net/http"

	"todoapp-backend/internal/controller"
	"todoapp-backend/internal/env"
	"todoapp-backend/internal/logger"
	"todoapp-backend/internal/repository"
	"todoapp-backend/internal/router"
	"todoapp-backend/internal/service"
)

func main() {
	logger.Setup()
	defer logger.Cleanup()

	dbConfig := &repository.Config{}

	var err error
	dbConfig.URI, err = env.GetDBURI()
	if err != nil {
		log.Panicln("Error getting database URI:", err)
	}

	apiPort, err := env.GetAPIPort()
	if err != nil {
		log.Panicln("Error getting API port:", err)
	}

	frontendURI, err := env.GetFrontendURI()
	if err != nil {
		log.Panicln("Error getting frontend URI:", err)
	}

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
	handler := router.SetupRoutes(frontendURI)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", apiPort), handler))
}
