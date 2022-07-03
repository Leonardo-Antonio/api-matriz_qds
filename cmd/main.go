package main

import (
	"github.com/Leonardo-Antonio/api-matriz_qds/internal/app"
	"github.com/labstack/echo/v4"
)

func main() {
	server := app.NewAppModule(echo.New(), "8000")
	server.Middlewares()
	server.Routers()
	server.Listening()
}
