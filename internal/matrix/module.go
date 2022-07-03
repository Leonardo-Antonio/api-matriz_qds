package matrix

import "github.com/labstack/echo/v4"

func Module(group *echo.Group) {
	handler := newHandler(newService())
	group.POST("/matrix", handler.Rotate)
}
