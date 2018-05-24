package models

import (
	"time"
)

// Todo model
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Description string `json:"description"`
	IsDone    bool      `json:"is_done"`
	ProjectID int       `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTodo(title string, isDone bool, projectId int) *Todo {
	return &Todo{
		Title:     title,
		IsDone:    isDone,
		ProjectID: projectId,
	}
}

// Todos model
type Todos []Todo
