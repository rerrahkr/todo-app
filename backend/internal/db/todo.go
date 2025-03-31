package db

import (
	"context"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create a new todo.
func NewTodo(content string) (*Todo, error) {
	cmd := `INSERT INTO todos (content) VALUES ($1) RETURNING *`
	row := pool.QueryRow(context.Background(), cmd, content)

	todo := &Todo{}
	if err := row.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return nil, err
	}

	return todo, nil
}

// Get all registered todos.
func GetAllTodos() ([]Todo, error) {
	cmd := `SELECT * FROM todos`
	rows, err := pool.Query(context.Background(), cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
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

// Get todo which has the given ID.
func GetTodo(id int) (*Todo, error) {
	cmd := `SELECT * FROM todos WHERE id = $1`
	row := pool.QueryRow(context.Background(), cmd, id)

	todo := &Todo{}
	if err := row.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return nil, err
	}

	return todo, nil
}

// Update a content of todo which has the given ID.
func UpdateTodo(id int, content string) (*Todo, error) {
	cmd := `UPDATE todos SET content = $1, updated_at = $2 WHERE id = $3 RETURNING *`
	row := pool.QueryRow(context.Background(), cmd, content, time.Now(), id)

	todo := &Todo{}
	if err := row.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return nil, err
	}

	return todo, nil
}

// Delete todo which has the given ID.
func DeleteTodo(id int) error {
	cmd := `DELETE FROM todos WHERE id = $1`
	_, err := pool.Exec(context.Background(), cmd, id)
	return err
}
