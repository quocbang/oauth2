package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quocbang/oauth2/delivery/middleware"
	"github.com/quocbang/oauth2/payload"
	"github.com/quocbang/oauth2/presenter"
)

func (h *Handlers) Callback(c echo.Context) error {
	var (
		ctx  = middleware.ToBuiltInContext(c)
		req  = payload.CallbackRequest{}
		resp *presenter.Oauth2LoginResponse
	)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp, err := h.Usecase.Auth().Google().Oauth2Login(ctx, req.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err) // TODO: should be echo error response
	}

	return c.JSON(http.StatusOK, resp)
}
