package service

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NewTodoRequest struct {
	Content string `json:"content" validate:"required"`
}

type NewTodoResponse Todo

type GetAllTodosResponse struct {
	Todos []Todo `json:"todos"`
}

type GetTodoByIDRequest struct {
	ID int `json:"id" validate:"required"`
}

type GetTodoByIDResponse Todo

type UpdateTodoRequest struct {
	ID      int    `json:"id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdateTodoResponse Todo

type DeleteTodoByIDRequest struct {
	ID int `json:"id" validate:"required"`
}
