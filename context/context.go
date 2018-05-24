package context

import (
	"database/sql"

	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	DB *sql.DB
}
