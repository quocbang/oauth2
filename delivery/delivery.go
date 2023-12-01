package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

// func (response) Error(c echo.Context, err error) error {
// 	var errMessage string

// 	return c.JSON(err.HTTPCode, map[string]interface{}{
// 		"code":    err.ErrorCode,
// 		"message": err.Message,
// 		"info":    errMessage,
// 	})
// }
