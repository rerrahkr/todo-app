package model

import "time"

type Todo struct {
	ID        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
