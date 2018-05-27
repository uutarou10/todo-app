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
	db := getDB(c)
	requestID := c.Param("id")

	row := db.QueryRow("SELECT * FROM todos WHERE id=?", requestID)

	var (
		id          int
		title       string
		description string
		isDone      bool
		projectID   int
		createdAt   time.Time
		updatedAt   time.Time
	)
	row.Scan(&id, &title, &description, &isDone, &projectID, &createdAt, &updatedAt)
	if id == 0 {
		return c.String(http.StatusNotFound, "not found")
	}

	todo := models.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		IsDone:      isDone,
		ProjectID:   projectID,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	return c.JSON(http.StatusOK, todo)
}

func CreateTodoHandler(c echo.Context) error {
	db := getDB(c)
	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return c.String(http.StatusBadRequest, "Invalid params.")
	}

	result, err := db.Exec("INSERT INTO todos (title, description, isDone, projectId) VALUES (?, ?, ?, ?)", todo.Title, todo.Description, todo.IsDone, todo.ProjectID)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// LastInsertIdをみていい感じにする
	insertedID, _ := result.LastInsertId()
	insertedRow := db.QueryRow("SELECT * FROM todos where id=?", insertedID)

	var (
		id          int
		title       string
		description string
		isDone      bool
		projectID   int
		createdAt   time.Time
		updatedAt   time.Time
	)
	insertedRow.Scan(&id, &title, &description, &isDone, &projectID, &createdAt, &updatedAt)

	todo = &models.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		IsDone:      isDone,
		ProjectID:   projectID,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodoHandler(c echo.Context) error {
	db := getDB(c)
	requestId := c.Param("id")

	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	result, _ := db.Exec("UPDATE todos SET title=?, description=?, isDone=?, projectId=?, updatedAt=CURRENT_TIMESTAMP() where id=?", todo.Title, todo.Description, todo.IsDone, todo.ProjectID, requestId)
	if affectedRow, _ := result.RowsAffected(); affectedRow <= 0 {
		return c.NoContent(http.StatusNotFound)
	}

	updatedRow := db.QueryRow("SELECT * FROM todos where id=?", requestId)

	var (
		id          int
		title       string
		description string
		isDone      bool
		projectID   int
		createdAt   time.Time
		updatedAt   time.Time
	)
	updatedRow.Scan(&id, &title, &description, &isDone, &projectID, &createdAt, &updatedAt)

	todo = &models.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		IsDone:      isDone,
		ProjectID:   projectID,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	return c.JSON(http.StatusCreated, todo)
}

func DeleteTodoHandler(c echo.Context) error {
	db := getDB(c)
	id := c.Param("id")

	result, err := db.Exec("DELETE FROM todos where id=?", id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// 変更された行が存在しなかったら
	if affectedRows, _ := result.RowsAffected(); affectedRows <= 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}
