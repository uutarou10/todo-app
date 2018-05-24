package handlers

import (
	"database/sql"
	"github.com/uutarou10/todo-app/context"
	"github.com/labstack/echo"
)

func getDB(c echo.Context) *sql.DB {
	cc, ok := c.(*context.Context)
	if !ok {
		panic("Cannnot cast to custom context!")
	}

	return cc.DB
}