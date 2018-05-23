package main

import (
	"fmt"

	"github.com/uutarou10/todo-app/handlers"
	"github.com/labstack/echo"
)

// App global values
type App struct {
	Echo *echo.Echo
	Host string
	Port int
}

// NewApp creates App instance
func NewApp(host string, port int) *App {
	return &App{
		Echo: echo.New(),
		Host: host,
		Port: port,
	}
}

// Run calls start server
func (a *App) Run() {
	a.RegisterRoutes()
	a.Echo.Start(fmt.Sprintf("%s:%d", a.Host, a.Port))
}

// RegisterRoutes registration endpoints.
func (a *App) RegisterRoutes() {
	a.Echo.GET("/", handlers.HelloHandler)
}

func main() {
	app := NewApp("localhost", 8080)
	app.Run()
}
