package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path"
	"strconv"
	"todoapp-backend/internal/service"

	"github.com/go-playground/validator/v10"
)

type TodoController interface {
	NewTodo(w http.ResponseWriter, r *http.Request)
	GetAllTodos(w http.ResponseWriter, r *http.Request)
	GetTodoByID(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodoByID(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	todoService service.TodoService
}

func NewTodoController(service service.TodoService) TodoController {
	return &todoController{
		todoService: service,
	}
}

func (c *todoController) NewTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := &service.NewTodoRequest{}
	err = json.Unmarshal(body, req)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusBadRequest)
		return
	}

	if err = validator.New(validator.WithRequiredStructEnabled()).Struct(req); err != nil {
		http.Error(w, "Failed to create todo", http.StatusBadRequest)
		return
	}

	if err := c.todoService.NewTodo(req); err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	res, err := c.todoService.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
		return
	}
}

func (c *todoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	res, err := c.todoService.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
		return
	}
}

// Parse ID from path.
func parseID(pathStr string) (int, error) {
	idStr := path.Base(pathStr)
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid ID")
	}
	return id, nil
}

func (c *todoController) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	req := &service.GetTodoByIDRequest{ID: id}
	res, err := c.todoService.GetTodoByID(req)
	if err != nil {
		http.Error(w, "Failed to get todo", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to get todo", http.StatusInternalServerError)
		return
	}
}

func (c *todoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := &service.UpdateTodoRequest{}
	if err := json.Unmarshal(body, req); err != nil {
		http.Error(w, "Failed to update todo", http.StatusBadRequest)
		return
	}
	req.ID = id

	if err = validator.New(validator.WithRequiredStructEnabled()).Struct(req); err != nil {
		http.Error(w, "Failed to update todo", http.StatusBadRequest)
		return
	}

	if err := c.todoService.UpdateTodo(req); err != nil {
		// Try to create a new todo.
		if err := c.todoService.NewTodo(&service.NewTodoRequest{Content: req.Content}); err != nil {
			http.Error(w, "Failed to update todo", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	res, err := c.todoService.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
		return
	}
}

func (c *todoController) DeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	req := &service.DeleteTodoByIDRequest{ID: id}
	if err := c.todoService.DeleteTodoByID(req); err != nil {
		http.Error(w, "Failed to delete todo", http.StatusNotFound)
		return
	}

	res, err := c.todoService.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
		return
	}
}
