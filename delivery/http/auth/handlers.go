package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/quocbang/oauth2/usecase"
)

type Handlers struct {
	Usecase *usecase.UseCase
}

func NewAuthHandler(e *echo.Group, usecase *usecase.UseCase) {
	handlers := &Handlers{
		Usecase: usecase,
	}

	// google
	e.GET("/google/auth", handlers.GetGoogleAuthURL)
	e.GET("/google/callback", handlers.GoogleCallback)

	// github
	e.GET("/github/auth", handlers.GetGithubAuthURL)
	e.GET("/github/callback", handlers.GithubCallback)
}
