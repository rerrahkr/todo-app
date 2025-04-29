package router

import (
	"net/http"
	"todoapp-backend/internal/controller"

	"github.com/rs/cors"
)

type TodoRouter interface {
	SetupRoutes(allowedOrigin string) http.Handler
}

type todoRouter struct {
	controller controller.TodoController
}

func NewTodoRouter(c controller.TodoController) TodoRouter {
	return &todoRouter{
		controller: c,
	}
}

func (ro *todoRouter) handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ro.controller.GetAllTodos(w, r)

	case http.MethodPost:
		ro.controller.NewTodo(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ro *todoRouter) handleTodoByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ro.controller.GetTodoByID(w, r)

	case http.MethodPatch:
		ro.controller.UpdateTodoByID(w, r)

	case http.MethodDelete:
		ro.controller.DeleteTodoByID(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ro *todoRouter) SetupRoutes(allowedOrigin string) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", ro.handleTodos)
	mux.HandleFunc("/todos/", ro.handleTodoByID)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	})

	return c.Handler(mux)
}
