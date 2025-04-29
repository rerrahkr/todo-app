package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
	controller_mock "todoapp-backend/internal/controller/mock"

	"go.uber.org/mock/gomock"
)

func TestNewTodo(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mockController := controller_mock.NewTodoControllerMock(mc)
	mockController.EXPECT().NewTodo(gomock.Any(), gomock.Any()).Times(1)

	router := NewTodoRouter(mockController)
	handler := router.SetupRoutes("*")

	req := httptest.NewRequest(http.MethodPost, "/todos", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
}

func TestGetTodos(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mockController := controller_mock.NewTodoControllerMock(mc)
	mockController.EXPECT().GetAllTodos(gomock.Any(), gomock.Any()).Times(1)

	router := NewTodoRouter(mockController)
	handler := router.SetupRoutes("*")

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
}

func TestGetTodoByID(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mockController := controller_mock.NewTodoControllerMock(mc)
	mockController.EXPECT().GetTodoByID(gomock.Any(), gomock.Any()).Times(1)

	router := NewTodoRouter(mockController)
	handler := router.SetupRoutes("*")

	req := httptest.NewRequest(http.MethodGet, "/todos/0", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
}

func TestUpdateTodoByID(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mockController := controller_mock.NewTodoControllerMock(mc)
	mockController.EXPECT().UpdateTodoByID(gomock.Any(), gomock.Any()).Times(1)

	router := NewTodoRouter(mockController)
	handler := router.SetupRoutes("*")

	req := httptest.NewRequest(http.MethodPatch, "/todos/0", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
}

func TestDeleteTodoByID(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mockController := controller_mock.NewTodoControllerMock(mc)
	mockController.EXPECT().DeleteTodoByID(gomock.Any(), gomock.Any()).Times(1)

	router := NewTodoRouter(mockController)
	handler := router.SetupRoutes("*")

	req := httptest.NewRequest(http.MethodDelete, "/todos/0", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
}
