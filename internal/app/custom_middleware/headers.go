package custom_middleware

import (
	"net/http"

	"github.com/Leonardo-Antonio/api-matriz_qds/pkg/response"
	"github.com/labstack/echo/v4"
)

func Headers() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("Content-Type") != "application/json" {
				return c.JSON(
					http.StatusBadRequest,
					response.ResponseErr("la data se debe enviar en el formato 'application/json'", nil),
				)
			}
			return next(c)
		}
	}
}
