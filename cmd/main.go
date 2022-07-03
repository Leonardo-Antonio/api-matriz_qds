package main

import (
	"github.com/Leonardo-Antonio/api-matriz_qds/internal/matrix"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	group := app.Group("/api/v1")
	matrix.Module(group)

	app.Start(":8000")
}
