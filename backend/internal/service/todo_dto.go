package service

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewTodoRequest struct {
	Content string `json:"content" validate:"required"`
}

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

type DeleteTodoByIDRequest struct {
	ID int `json:"id" validate:"required"`
}
