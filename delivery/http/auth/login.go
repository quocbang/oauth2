package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quocbang/oauth2/delivery"
	"github.com/quocbang/oauth2/delivery/middleware"
	"github.com/quocbang/oauth2/presenter"
)

func (h *Handlers) GoogleLogin(c echo.Context) error {
	var (
		ctx  = middleware.ToBuiltInContext(c)
		resp *presenter.LoginResponse
	)
	resp, err := h.Usecase.Auth().Google().Login(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return delivery.Response.Success(c, resp)
}

func (h *Handlers) GetGithubAuthURL(c echo.Context) error {
	var (
		ctx  = middleware.ToBuiltInContext(c)
		resp *presenter.LoginResponse
	)
	resp, err := h.Usecase.Auth().Github().Login(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return delivery.Response.Success(c, resp)
}
