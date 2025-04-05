package service

import (
	"todoapp-backend/internal/model"
	"todoapp-backend/internal/repository"
)

type TodoService interface {
	NewTodo(req *NewTodoRequest) error
	GetAllTodos() (*GetAllTodosResponse, error)
	GetTodoByID(req *GetTodoByIDRequest) (*GetTodoByIDResponse, error)
	UpdateTodo(req *UpdateTodoRequest) error
	DeleteTodoByID(req *DeleteTodoByIDRequest) error
}

type todoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepo: repo,
	}
}

func (s *todoService) NewTodo(req *NewTodoRequest) error {
	todo := &model.Todo{
		Content: req.Content,
	}

	_, err := s.todoRepo.NewTodo(todo)

	return err
}

func (s *todoService) GetAllTodos() (*GetAllTodosResponse, error) {
	todos, err := s.todoRepo.GetAllTodos()
	if err != nil {
		return nil, err
	}

	resTodos := []Todo{}
	for _, todo := range todos {
		resTodos = append(resTodos, Todo{
			ID:        todo.ID,
			Content:   todo.Content,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		})
	}

	return &GetAllTodosResponse{Todos: resTodos}, nil
}

func (s *todoService) GetTodoByID(req *GetTodoByIDRequest) (*GetTodoByIDResponse, error) {
	todo, err := s.todoRepo.GetTodoByID(req.ID)
	if err != nil {
		return nil, err
	}

	return &GetTodoByIDResponse{
		ID:        todo.ID,
		Content:   todo.Content,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}, nil
}

func (s *todoService) UpdateTodo(req *UpdateTodoRequest) error {
	todo := &model.Todo{
		ID:      req.ID,
		Content: req.Content,
	}

	_, err := s.todoRepo.UpdateTodo(todo)

	return err
}

func (s *todoService) DeleteTodoByID(req *DeleteTodoByIDRequest) error {
	return s.todoRepo.DeleteTodoByID(req.ID)
}
