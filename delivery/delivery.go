package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quocbang/oauth2/errors"
)

type response struct{}

var Response response

func (response) Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    data,
	})
}

func (response) Error(c echo.Context, err error) error {
	serverErr, ok := err.(errors.Error)
	if !ok {
		return err // return origin error
	}

	return c.JSON(serverErr.StatusCode, map[string]interface{}{
		"code":    serverErr.ErrorCode,
		"message": serverErr.Details,
		"info":    serverErr.Raw.Error(),
	})
}
