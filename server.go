package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uutarou10/todo-app/context"
	"github.com/uutarou10/todo-app/handlers"
)

// App global values
type App struct {
	Echo *echo.Echo
	Host string
	Port int
	DB   *sql.DB
}

// NewApp creates App instance
func NewApp(host string, port int) *App {
	// FIXME: 環境変数とかから受け取るようにしないとなぁ
	db, err := sql.Open("mysql", "root:password@/todoapp")
	if err != nil {
		panic("Cannot establish db connection.")
	}

	return &App{
		Echo: echo.New(),
		Host: host,
		Port: port,
		DB:   db,
	}
}

// Run calls start server
func (a *App) Run() {
	a.RegisterRoutes()
	a.RegisterMiddlewares()
	a.Echo.Start(fmt.Sprintf("%s:%d", a.Host, a.Port))
}

// RegisterRoutes registration endpoints.
func (a *App) RegisterRoutes() {
	a.Echo.GET("/", handlers.HelloHandler)
	a.Echo.GET("/db", handlers.DBAccessTestHandler)
}

// RegisterMiddlewares registrationg middlewares.
func (a *App) RegisterMiddlewares() {
	// DB provider
	a.Echo.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		if a.DB == nil {
			panic("Not exist db connection.")
		}
		return func(c echo.Context) error {
			cc := &context.Context{
				c,
				a.DB,
			}

			return h(cc)
		}
	})
	a.Echo.Use(middleware.Logger())
}

func main() {
	app := NewApp("localhost", 3000)
	app.Run()
}
