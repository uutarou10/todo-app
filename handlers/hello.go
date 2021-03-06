package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/uutarou10/todo-app/context"
)

func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!!")
}

func DBAccessTestHandler(c echo.Context) error {
	cc, ok := c.(*context.Context)
	if !ok {
		return c.String(http.StatusInternalServerError, "Cannot cast echo.Context to custom context.")
	}

	rows, _ := cc.DB.Query("SELECT * FROM todos")
	defer rows.Close()

	columuns, _ := rows.Columns()

	return c.JSON(http.StatusOK, columuns)
}
