package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type key int

const (
	ClientIP key = iota + 1
	ClientAgent
	Logger
)

func contextWithClientIP(parent context.Context, c echo.Context) context.Context {
	return context.WithValue(parent, ClientIP, c.RealIP())
}

func contextWithClientAgent(parent context.Context, c echo.Context) context.Context {
	return context.WithValue(parent, ClientAgent, c.Request().UserAgent())
}

func GetClientIP(ctx context.Context) string {
	ip, ok := ctx.Value(ClientIP).(string)
	if !ok {
		return ""
	}
	return ip
}

func GetClientAgent(ctx context.Context) string {
	clientAgent, ok := ctx.Value(ClientAgent).(string)
	if !ok {
		return ""
	}
	return clientAgent
}

// ToBuiltInContext is convert context of echo to context built of Golang
// contains:
// - client ip
// - client agent
// - *zap.Logger
func ToBuiltInContext(c echo.Context) context.Context {
	ctx := c.Request().Context()
	ctx = contextWithClientIP(ctx, c)
	ctx = contextWithClientAgent(ctx, c)
	ctx = setLoggerToContext(ctx)
	return ctx
}

func setLoggerToContext(parent context.Context) context.Context {
	requestID := xid.New().String()
	logger := zap.L().With(zap.String("req_id", requestID))
	return context.WithValue(parent, Logger, logger)
}

func GetLoggerFormContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(Logger).(*zap.Logger)
	if ok {
		return logger
	}
	return zap.L()
}
