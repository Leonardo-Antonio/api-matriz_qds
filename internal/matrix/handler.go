package matrix

import (
	"net/http"

	"github.com/Leonardo-Antonio/api-matriz_qds/pkg/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service iService
}

func newHandler(service iService) *handler {
	return &handler{service}
}

func (h *handler) Rotate(ctx echo.Context) error {
	var body *matrixDTO
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			response.ResponseErr("la data enviada no tiene la estructura o el tipo correcto", nil),
		)
	}

	h.service.Rotate90(body)

	return ctx.JSON(
		http.StatusOK,
		response.ResponseSuccess("ok", body),
	)
}
