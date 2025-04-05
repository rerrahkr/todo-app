package service

import (
	"errors"
	"testing"

	"todoapp-backend/internal/model"
)

type TodoRepositoryStub struct{}

func (s *TodoRepositoryStub) NewTodo(todo *model.Todo) (*model.Todo, error) {
	return todo, nil
}

func (s *TodoRepositoryStub) GetAllTodos() ([]model.Todo, error) {
	return []model.Todo{
		{ID: 0, Content: "A"},
		{ID: 1, Content: "B"},
	}, nil
}

func (s *TodoRepositoryStub) GetTodoByID(id int) (*model.Todo, error) {
	todo := model.Todo{ID: 0, Content: "0"}

	if todo.ID != id {
		return nil, errors.New("")
	}

	return &todo, nil
}

func (s *TodoRepositoryStub) UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	td := model.Todo{ID: 0, Content: "1"}

	if td.ID != todo.ID {
		return nil, errors.New("")
	}

	return &model.Todo{ID: td.ID, Content: td.Content}, nil
}

func (s *TodoRepositoryStub) DeleteTodoByID(id int) error {
	if id != 0 {
		return errors.New("")
	}

	return nil
}

var service TodoService

func TestMain(m *testing.M) {
	// Preprocess.
	service = NewTodoService(&TodoRepositoryStub{})

	m.Run()
}

func TestNewTodo(t *testing.T) {
	req := NewTodoRequest{"New!"}

	err := service.NewTodo(&req)
	if err != nil {
		t.Fatalf("Failed in TodoService.NewTodo: %v", err)
	}
}

func TestGetAllTodos(t *testing.T) {
	expected := GetAllTodosResponse{
		Todos: []Todo{
			{ID: 0, Content: "A"},
			{ID: 1, Content: "B"},
		},
	}

	actual, err := service.GetAllTodos()
	if err != nil {
		t.Fatalf("Failed in TodoService.GetAllTodos: %v", err)
	}

	if len(actual.Todos) != len(expected.Todos) {
		t.Fatal("Failed in TodoService.GetAllTodos: Invalid length of Todos")
	}

	for i, todo := range actual.Todos {
		if todo.ID != expected.Todos[i].ID ||
			todo.Content != expected.Todos[i].Content {
			t.Fatal("Failed in TodoService.GetAllTodos: Todo is not same content")
		}
	}
}

func TestGetTodoByID(t *testing.T) {
	t.Run("Give a valid ID", func(t *testing.T) {
		expected := model.Todo{ID: 0, Content: "0"}
		actual, err := service.GetTodoByID(&GetTodoByIDRequest{ID: 0})
		if err != nil {
			t.Fatalf("Failed in TodoService.GetTodoByID: %v", err)
		} else if actual.ID != expected.ID || actual.Content != expected.Content {
			t.Fatal("Failed in TodoService.GetTodoByID: Response is unexpected value")
		}
	})

	t.Run("Give an invalid ID", func(t *testing.T) {
		_, err := service.GetTodoByID(&GetTodoByIDRequest{ID: -1})
		if err == nil {
			t.Errorf("Failed in TodoService.GetTodoByID: No error occurs when an invalid ID is given")
		}

		_, err = service.GetTodoByID(&GetTodoByIDRequest{ID: 1})
		if err == nil {
			t.Errorf("Failed in TodoService.GetTodoByID: No error occurs when an invalid ID is given")
		}
	})
}

func TestUpdateTodo(t *testing.T) {
	t.Run("Give a valid ID", func(t *testing.T) {
		req := UpdateTodoRequest{ID: 0, Content: "ABC"}
		if err := service.UpdateTodo(&req); err != nil {
			t.Fatalf("Failed in TodoService.UpdateTodo: %v", err)
		}
	})

	t.Run("Give an invalid ID", func(t *testing.T) {
		req := UpdateTodoRequest{ID: 1}
		if err := service.UpdateTodo(&req); err == nil {
			t.Fatal("Failed in TodoService.UpdateTodo: No error occurs when an invalid ID is given")
		}
	})
}

func TestDeleteTodoByID(t *testing.T) {
	t.Run("Give a valid ID", func(t *testing.T) {
		req := DeleteTodoByIDRequest{ID: 0}
		if err := service.DeleteTodoByID(&req); err != nil {
			t.Fatalf("Failed in TodoService.DeleteTodo: %v", err)
		}
	})

	t.Run("Give an invalid ID", func(t *testing.T) {
		req := DeleteTodoByIDRequest{ID: 1}
		if err := service.DeleteTodoByID(&req); err == nil {
			t.Fatal("Failed in TodoService.DeleteTodo: No error occurs when an invalid ID is given")
		}
	})
}
