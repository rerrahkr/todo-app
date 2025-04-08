package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapp-backend/internal/service"
)

type TodoServiceMock struct{}

var expectedTodos = []service.Todo{
	{ID: 0, Content: "A"},
	{ID: 1, Content: "B"},
}

func (s *TodoServiceMock) NewTodo(req *service.NewTodoRequest) error {
	if req.Content != "New" {
		return errors.New("Unexpected content")
	}

	return nil
}

func (*TodoServiceMock) GetAllTodos() (*service.GetAllTodosResponse, error) {
	return &service.GetAllTodosResponse{
		Todos: expectedTodos,
	}, nil
}

func (*TodoServiceMock) GetTodoByID(req *service.GetTodoByIDRequest) (*service.GetTodoByIDResponse, error) {
	if req.ID != 0 {
		return nil, errors.New("")
	}

	return &service.GetTodoByIDResponse{ID: 0, Content: "A"}, nil
}

func (*TodoServiceMock) UpdateTodo(req *service.UpdateTodoRequest) error {
	if req.ID != 0 {
		return errors.New("")
	}

	return nil
}

func (*TodoServiceMock) DeleteTodoByID(req *service.DeleteTodoByIDRequest) error {
	if req.ID != 0 {
		return errors.New("")
	}

	return nil
}

var controller TodoController

func verifyTodosResponse(w *httptest.ResponseRecorder) error {
	if w.Header().Get("Content-Type") != "application/json" {
		return errors.New("Content-Type is not application/json")
	}

	res := &service.GetAllTodosResponse{}
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		return errors.New("Invalid response body")
	}

	if len(res.Todos) != len(expectedTodos) {
		return fmt.Errorf("Count of todos is %v but actual is %v", len(expectedTodos), len(res.Todos))
	}

	for i, todo := range res.Todos {
		if todo.ID != expectedTodos[i].ID {
			return fmt.Errorf("ID of todo[%v] is %v but actual is %v", i, expectedTodos[i].ID, todo.ID)
		} else if todo.Content != expectedTodos[i].Content {
			return fmt.Errorf("Content of todo[%v] is %v but actual is %v", i, expectedTodos[i].Content, todo.Content)
		}
	}

	return nil
}

func TestMain(m *testing.M) {
	// Preprocess.
	controller = NewTodoController(&TodoServiceMock{})

	m.Run()
}

func TestNewTodo(t *testing.T) {
	body := map[string]string{
		"content": "New",
	}
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	controller.NewTodo(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Failed in TodoController.NewTodo: Status code is %v", w.Code)
	}

	if err := verifyTodosResponse(w); err != nil {
		t.Fatalf("Failed in TodoController.NewTodo: %v", err)
	}
}

func TestGetAllTodos(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()

	controller.GetAllTodos(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Failed in TodoController.GetAllTodos: Status code is %v", w.Code)
	}

	if err := verifyTodosResponse(w); err != nil {
		t.Fatalf("Failed in TodoController.GetAllTodos: %v", err)
	}
}

func TestGetTodoByID(t *testing.T) {
	t.Run("Valid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/0", nil)
		w := httptest.NewRecorder()

		controller.GetTodoByID(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Failed in TodoController.GetTodoByID: Status code is %v", w.Code)
		}

		if w.Header().Get("Content-Type") != "application/json" {
			t.Fatal("Failed in TodoController.GetTodoByID: Content-Type is not application/json")
		}

		res := &service.GetTodoByIDResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
			t.Fatal("Failed in TodoController.GetTodoByID: Invalid response body")
		}

		if res.ID != 0 {
			t.Fatalf("Failed in TodoController.GetTodoByID: ID of todo is %v but actual is %v", 0, res.ID)
		} else if res.Content != "A" {
			t.Fatalf("Failed in TodoController.GetTodoByID: Content of todo is %v but actual is %v", "A", res.Content)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		w := httptest.NewRecorder()

		controller.GetTodoByID(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("Failed in TodoController.GetTodoByID: Status code is %v", w.Code)
		}
	})
}

func TestUpdateTodo(t *testing.T) {
	t.Run("Valid ID", func(t *testing.T) {
		body := map[string]string{
			"content": "New",
		}
		jsonBody, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPut, "/todos/0", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		controller.UpdateTodo(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Failed in TodoController.UpdateTodo: Status code is %v", w.Code)
		}

		if err := verifyTodosResponse(w); err != nil {
			t.Fatalf("Failed in TodoController.UpdateTodo: %v", err)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		body := map[string]string{
			"content": "New",
		}
		jsonBody, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPut, "/todos/1", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		controller.UpdateTodo(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("Failed in TodoController.UpdateTodo: Status code is %v", w.Code)
		}

		if err := verifyTodosResponse(w); err != nil {
			t.Fatalf("Failed in TodoController.UpdateTodo: %v", err)
		}
	})
}

func TestDeleteTodoByID(t *testing.T) {
	t.Run("Valid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/0", nil)
		w := httptest.NewRecorder()

		controller.DeleteTodoByID(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Failed in TodoController.DeleteTodo: Status code is %v", w.Code)
		}

		if err := verifyTodosResponse(w); err != nil {
			t.Fatalf("Failed in TodoController.DeleteTodo: %v", err)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		w := httptest.NewRecorder()

		controller.DeleteTodoByID(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("Failed in TodoController.DeleteTodo: Status code is %v", w.Code)
		}
	})
}
