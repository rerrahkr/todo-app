package repository

import (
	"context"
	"time"

	"todoapp-backend/internal/model"
)

// Interface of todo repository.
type TodoRepository interface {
	NewTodo(todo *model.Todo) (*model.Todo, error)
	GetAllTodos() ([]model.Todo, error)
	GetTodoByID(id int) (*model.Todo, error)
	UpdateTodo(todo *model.Todo) (*model.Todo, error)
	DeleteTodoByID(id int) error
}

// Instance of todo repository.
type todoRepository struct{}

// Create a new todo repository.
func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (*todoRepository) NewTodo(todo *model.Todo) (*model.Todo, error) {
	cmd := `INSERT INTO todos (content) VALUES ($1) RETURNING *`
	row := pool.QueryRow(context.Background(), cmd, todo.Content)

	newTodo := &model.Todo{}
	if err := row.Scan(&newTodo.ID, &newTodo.Content, &newTodo.CreatedAt, &newTodo.UpdatedAt); err != nil {
		return nil, err
	}

	return newTodo, nil
}

func (*todoRepository) GetAllTodos() ([]model.Todo, error) {
	cmd := `SELECT * FROM todos`
	rows, err := pool.Query(context.Background(), cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (*todoRepository) GetTodoByID(id int) (*model.Todo, error) {
	cmd := `SELECT * FROM todos WHERE id = $1`
	row := pool.QueryRow(context.Background(), cmd, id)

	todo := &model.Todo{}
	if err := row.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return nil, err
	}

	return todo, nil
}

func (*todoRepository) UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	cmd := `UPDATE todos SET content = $1, updated_at = $2 WHERE id = $3 RETURNING *`
	row := pool.QueryRow(context.Background(), cmd, todo.Content, time.Now(), todo.ID)

	newTodo := &model.Todo{}
	if err := row.Scan(&newTodo.ID, &newTodo.Content, &newTodo.CreatedAt, &newTodo.UpdatedAt); err != nil {
		return nil, err
	}

	return newTodo, nil
}

func (*todoRepository) DeleteTodoByID(id int) error {
	cmd := `DELETE FROM todos WHERE id = $1`
	_, err := pool.Exec(context.Background(), cmd, id)
	return err
}
