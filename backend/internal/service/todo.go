package service

import (
	"todoapp-backend/internal/model"
	"todoapp-backend/internal/repository"
)

type TodoService interface {
	NewTodo(req *NewTodoRequest) (*NewTodoResponse, error)
	GetAllTodos() (*GetAllTodosResponse, error)
	GetTodoByID(req *GetTodoByIDRequest) (*GetTodoByIDResponse, error)
	UpdateTodo(req *UpdateTodoRequest) (*UpdateTodoResponse, error)
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

func (s *todoService) NewTodo(req *NewTodoRequest) (*NewTodoResponse, error) {
	todo := &model.Todo{
		Content: req.Content,
	}

	newTodo, err := s.todoRepo.NewTodo(todo)
	if err != nil {
		return nil, err
	}

	res := &NewTodoResponse{
		ID:        newTodo.ID,
		Content:   newTodo.Content,
		CreatedAt: newTodo.CreatedAt,
		UpdatedAt: newTodo.UpdatedAt,
	}

	return res, nil
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

func (s *todoService) UpdateTodo(req *UpdateTodoRequest) (*UpdateTodoResponse, error) {
	todo := &model.Todo{
		ID:      req.ID,
		Content: req.Content,
	}

	updatedTodo, err := s.todoRepo.UpdateTodo(todo)
	if err != nil {
		return nil, err
	}

	res := &UpdateTodoResponse{
		ID:        updatedTodo.ID,
		Content:   updatedTodo.Content,
		CreatedAt: updatedTodo.CreatedAt,
		UpdatedAt: updatedTodo.UpdatedAt,
	}

	return res, nil
}

func (s *todoService) DeleteTodoByID(req *DeleteTodoByIDRequest) error {
	return s.todoRepo.DeleteTodoByID(req.ID)
}
