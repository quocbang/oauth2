package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quocbang/oauth2/delivery"
	"github.com/quocbang/oauth2/delivery/middleware"
	"github.com/quocbang/oauth2/errors"
	"github.com/quocbang/oauth2/payload"
	"github.com/quocbang/oauth2/presenter"
)

func (h *Handlers) GoogleCallback(c echo.Context) error {
	var (
		ctx  = middleware.ToBuiltInContext(c)
		req  = payload.CallbackRequest{}
		resp *presenter.Oauth2LoginResponse
	)

	if err := c.Bind(&req); err != nil {
		return delivery.Response.Error(c, errors.Error{
			StatusCode: http.StatusBadRequest,
			ErrorCode:  errors.Code_ERROR_BAD_REQUEST,
			Details:    "bad request",
			Raw:        err,
		})
	}

	if err := req.Validate(); err != nil {
		return delivery.Response.Error(c, errors.Error{
			StatusCode: http.StatusBadRequest,
			ErrorCode:  errors.Code_ERROR_BAD_REQUEST,
			Details:    "bad request",
			Raw:        err,
		})
	}

	resp, err := h.Usecase.Auth().Google().Oauth2Login(ctx, req.Code)
	if err != nil {
		return delivery.Response.Error(c, err)
	}

	return delivery.Response.Success(c, resp)
}

func (h *Handlers) GithubCallback(c echo.Context) error {
	var (
		ctx  = middleware.ToBuiltInContext(c)
		req  = payload.CallbackRequest{}
		resp *presenter.Oauth2LoginResponse
	)

	if err := c.Bind(&req); err != nil {
		return delivery.Response.Error(c, errors.Error{
			StatusCode: http.StatusBadRequest,
			ErrorCode:  errors.Code_ERROR_BAD_REQUEST,
			Details:    "bad request",
			Raw:        err,
		})
	}

	if err := req.Validate(); err != nil {
		return delivery.Response.Error(c, errors.Error{
			StatusCode: http.StatusBadRequest,
			ErrorCode:  errors.Code_ERROR_BAD_REQUEST,
			Details:    "bad request",
			Raw:        err,
		})
	}

	resp, err := h.Usecase.Auth().Github().Oauth2Login(ctx, req.Code)
	if err != nil {
		return delivery.Response.Error(c, err)
	}

	return delivery.Response.Success(c, resp)
}
