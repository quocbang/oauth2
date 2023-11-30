package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

type key int

const (
	ClientIP key = iota + 1
	ClientAgent
)

func setNetWorkInfo(ctx echo.Context) context.Context {
	return nil
}
