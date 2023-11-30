package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) GoogleLogin(ctx echo.Context) error {
	url, err := h.Usecase.Auth().Google().Login(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.HTML(http.StatusOK, url)
}
