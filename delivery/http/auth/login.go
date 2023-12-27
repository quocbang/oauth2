package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/quocbang/oauth2/delivery"
	"github.com/quocbang/oauth2/presenter"
)

func (h *Handlers) GetGoogleAuthURL(c echo.Context) error {
	var (
		ctx  = c.Request().Context()
		resp *presenter.GetAuthURLResponse
	)
	resp, err := h.Usecase.Auth().Google().GetAuthURL(ctx)
	if err != nil {
		return delivery.Response.Error(c, err)
	}
	return delivery.Response.Success(c, resp)
}

func (h *Handlers) GetGithubAuthURL(c echo.Context) error {
	var (
		ctx  = c.Request().Context()
		resp *presenter.GetAuthURLResponse
	)
	resp, err := h.Usecase.Auth().Github().GetAuthURL(ctx)
	if err != nil {
		return delivery.Response.Error(c, err)
	}
	return delivery.Response.Success(c, resp)
}
