package app

import (
	"fmt"

	"github.com/Leonardo-Antonio/api-matriz_qds/internal/app/custom_middleware"
	"github.com/Leonardo-Antonio/api-matriz_qds/internal/matrix"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type appModule struct {
	app   *echo.Echo
	group *echo.Group
	port  string
}

func NewAppModule(app *echo.Echo, port string) *appModule {
	return &appModule{
		app:   app,
		group: app.Group("/api/v1"),
		port:  port,
	}
}

func (server *appModule) Middlewares() {
	server.app.Use(middleware.CORS())
	server.app.Use(middleware.Logger())
	server.app.Use(custom_middleware.Headers())
}

func (server *appModule) Routers() {
	matrix.Module(server.group)
}

func (server *appModule) Listening() {
	server.app.Start(fmt.Sprintf(":%s", server.port))
}
