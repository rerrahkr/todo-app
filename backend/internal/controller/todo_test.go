package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapp-backend/internal/service"
)

type TodoServiceMock struct{}

var expectedTodos = []service.Todo{
	{ID: 1, Content: "A"},
	{ID: 2, Content: "B"},
}

func (s *TodoServiceMock) NewTodo(req *service.NewTodoRequest) (*service.NewTodoResponse, error) {
	if req.Content != "New" {
		return nil, errors.New("Unexpected content")
	}

	return &service.NewTodoResponse{Content: req.Content}, nil
}

func (*TodoServiceMock) GetAllTodos() (*service.GetAllTodosResponse, error) {
	return &service.GetAllTodosResponse{
		Todos: expectedTodos,
	}, nil
}

func (*TodoServiceMock) GetTodoByID(req *service.GetTodoByIDRequest) (*service.GetTodoByIDResponse, error) {
	if req.ID != 1 {
		return nil, errors.New("")
	}

	return &service.GetTodoByIDResponse{ID: 1, Content: "A"}, nil
}

func (*TodoServiceMock) UpdateTodo(req *service.UpdateTodoRequest) (*service.UpdateTodoResponse, error) {
	if req.ID != 1 {
		return nil, errors.New("")
	}

	return &service.UpdateTodoResponse{ID: req.ID, Content: req.Content}, nil
}

func (*TodoServiceMock) DeleteTodoByID(req *service.DeleteTodoByIDRequest) error {
	if req.ID != 1 {
		return errors.New("")
	}

	return nil
}

var controller TodoController

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

	if w.Header().Get("Content-Type") != "application/json" {
		t.Fatal("Failed in TodoController.NewTodo: Content-Type is not application/json")
	}

	res := &service.NewTodoResponse{}
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatal("Failed in TodoController.NewTodo: Invalid response body")
	}

	if res.Content != body["content"] {
		t.Errorf("Failed in TodoController.NewTodo: Content of todo is %v but actual is %v", body["content"], res.Content)
	}
}

func TestGetAllTodos(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()

	controller.GetAllTodos(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Failed in TodoController.GetAllTodos: Status code is %v", w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Failed in TodoController.GetAllTodos: Content-Type is not application/json")
	}

	res := &service.GetAllTodosResponse{}
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("Failed in TodoController.GetAllTodos: Invalid response body")
	}

	if len(res.Todos) != len(expectedTodos) {
		t.Fatalf("Failed in TodoController.GetAllTodos: Count of todos is %v but actual is %v", len(expectedTodos), len(res.Todos))
	}

	for i, todo := range res.Todos {
		if todo.ID != expectedTodos[i].ID {
			t.Errorf("Failed in TodoController.GetAllTodos: ID of todo[%v] is %v but actual is %v", i, expectedTodos[i].ID, todo.ID)
		} else if todo.Content != expectedTodos[i].Content {
			t.Errorf("Failed in TodoController.GetAllTodos: Content of todo[%v] is %v but actual is %v", i, expectedTodos[i].Content, todo.Content)
		}
	}
}

func TestGetTodoByID(t *testing.T) {
	t.Run("Valid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
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

		if res.ID != 1 {
			t.Errorf("Failed in TodoController.GetTodoByID: ID of todo is %v but actual is %v", 1, res.ID)
		}
		if res.Content != "A" {
			t.Errorf("Failed in TodoController.GetTodoByID: Content of todo is %v but actual is %v", "A", res.Content)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/2", nil)
		w := httptest.NewRecorder()

		controller.GetTodoByID(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("Failed in TodoController.GetTodoByID: Status code is %v", w.Code)
		}
	})
}

func TestUpdateTodo(t *testing.T) {
	t.Run("Existed ID", func(t *testing.T) {
		body := map[string]string{
			"content": "New",
		}
		jsonBody, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPatch, "/todos/1", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		controller.UpdateTodoByID(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Failed in TodoController.UpdateTodo: Status code is %v", w.Code)
		}

		if w.Header().Get("Content-Type") != "application/json" {
			t.Fatal("Failed in TodoController.UpdateTodo: Content-Type is not application/json")
		}

		res := &service.UpdateTodoResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
			t.Fatal("Failed in TodoController.UpdateTodo: Invalid response body")
		}

		if res.ID != 1 {
			t.Errorf("Failed in TodoController.UpdateTodo: ID of todo is %v but actual is %v", 1, res.ID)
		}
		if res.Content != body["content"] {
			t.Errorf("Failed in TodoController.UpdateTodo: Content of todo is %v but actual is %v", body["content"], res.Content)
		}
	})

	t.Run("Non-existed ID", func(t *testing.T) {
		body := map[string]string{
			"content": "New",
		}
		jsonBody, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPatch, "/todos/2", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		controller.UpdateTodoByID(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("Failed in TodoController.UpdateTodo: Status code is %v", w.Code)
		}
	})
}

func TestDeleteTodoByID(t *testing.T) {
	t.Run("Valid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		w := httptest.NewRecorder()

		controller.DeleteTodoByID(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Failed in TodoController.DeleteTodo: Status code is %v", w.Code)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/2", nil)
		w := httptest.NewRecorder()

		controller.DeleteTodoByID(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("Failed in TodoController.DeleteTodo: Status code is %v", w.Code)
		}
	})
}
