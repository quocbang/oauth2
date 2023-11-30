package cmd

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
	"golang.org/x/oauth2/google"

	"github.com/quocbang/oauth2/config"
	"github.com/quocbang/oauth2/delivery/http"
	myMiddleware "github.com/quocbang/oauth2/delivery/middleware"
)

func Run() {
	// new API
	e := echo.New()

	// get config
	cfg := config.GetConfig()

	// init logger
	myMiddleware.InitLogger(cfg.DevMode)

	// logging + middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// register router
	// "https://www.googleapis.com/auth/userinfo.email"
	// "https://www.googleapis.com/auth/userinfo.profile"
	h := http.HTTP{
		Auth: http.AuthConfig{
			Google: oauth2.Config{
				ClientID:     cfg.Oauth2.Google.ClientID,
				ClientSecret: cfg.Oauth2.Google.ClientSecret,
				Endpoint:     google.Endpoint,
				RedirectURL:  cfg.Oauth2.Google.RedirectURL,
				Scopes:       []string{"openid", "email", "profile"},
			},
			Github: oauth2.Config{
				ClientID:     cfg.Oauth2.Github.ClientID,
				ClientSecret: cfg.Oauth2.Github.ClientSecret,
				Endpoint:     endpoints.GitHub,
				RedirectURL:  cfg.Oauth2.Github.RedirectURL,
				Scopes:       []string{"https://api.github.com/user"},
			},
		},
	}
	h.RegisterHTTPHandler(e)

	// serve
	log.Fatal(e.Start(":3000"))
}
