package router

import (
	"net/http"
	"todoapp-backend/internal/controller"
)

type TodoRouter interface {
	SetupRoutes() http.Handler
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

	case http.MethodPut:
		ro.controller.UpdateTodo(w, r)

	case http.MethodDelete:
		ro.controller.DeleteTodoByID(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ro *todoRouter) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", ro.handleTodos)
	mux.HandleFunc("/todos/", ro.handleTodoByID)

	return mux
}
