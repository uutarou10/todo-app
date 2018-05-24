package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/uutarou10/todo-app/models"
)

func TodoIndexHandler(c echo.Context) error {
	db := getDB(c)
	rows, err := db.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	var todos models.Todos
	for rows.Next() {
		var (
			id          int
			title       string
			description string
			isDone      bool
			projectID   int
			createdAt   time.Time
			updatedAt   time.Time
		)

		err := rows.Scan(&id, &title, &description, &isDone, &projectID, &createdAt, &updatedAt)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Internal server error")
		}

		todos = append(todos, models.Todo{
			ID:          id,
			Title:       title,
			Description: description,
			IsDone:      isDone,
			ProjectID:   projectID,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}
	return c.JSON(http.StatusOK, todos)
}

func TodoHandler(c echo.Context) error {
	return nil
}
