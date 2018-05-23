package models

import (
	"time"
)

// Todo model
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	IsDone    bool      `json:"isDone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Todos model
type Todos []Todo
