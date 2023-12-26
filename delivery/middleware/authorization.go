package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quocbang/oauth2/errors"
	"github.com/quocbang/oauth2/utils/token"
)

type authKey string

const (
	AuthorizationKey authKey = "x-server-auth-key"
)

type auth struct {
	secretKey string
}

func NewAuthorization(secretKey string) IAuthorization {
	return &auth{
		secretKey: secretKey,
	}
}

func (a *auth) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		givenToken, ok := c.Get(string(AuthorizationKey)).(string)
		if !ok {
			return errors.Error{
				StatusCode: http.StatusUnauthorized,
				ErrorCode:  errors.Code_ERROR_MISSING_TOKEN,
				Details:    "missing token",
			}
		}

		token := token.JWT{
			SecretKey: a.secretKey,
		}
		claims, err := token.VerifyToken(givenToken)
		if err != nil {
			return errors.Error{
				StatusCode: http.StatusForbidden,
				ErrorCode:  errors.Code_ERROR_VERIFY_TOKEN_FAILED,
				Details:    err.Error(),
				Raw:        err,
			}
		}

		// set to main context
		r := c.Request()
		ctx := context.WithValue(r.Context(), AuthorizationKey, claims)
		r = r.WithContext(ctx) // append context to request

		return next(c.Echo().NewContext(r, c.Response())) // forward with context.Context value
	}
}

type IAuthorization interface {
	Authorization(echo.HandlerFunc) echo.HandlerFunc
}
