package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
	"golang.org/x/oauth2/google"

	"github.com/quocbang/oauth2/config"
	h "github.com/quocbang/oauth2/delivery/http"
	mdw "github.com/quocbang/oauth2/delivery/middleware"
	"github.com/quocbang/oauth2/delivery/websocket"
)

func Run() {
	// new API
	e := echo.New()

	// get config
	cfg := config.GetConfig()

	// logging + middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(mdw.WithBaseContextValues())

	// register router
	// "https://www.googleapis.com/auth/userinfo.email"
	// "https://www.googleapis.com/auth/userinfo.profile"
	httpConfig := h.HTTP{
		Auth: h.AuthConfig{
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
		Database:     cfg.Database,
		InternalAuth: cfg.InternalAuth,
		MigratePath:  cfg.MigratePath,
	}
	httpConfig.RegisterHTTPHandler(e)

	// serve
	log.Fatal(e.Start(":3000"))
}

func RunWebsocket() {
	// get config
	cfg := config.GetConfig()

	// register websocket handler
	handlers := websocket.NewWebsocketHandler(cfg)

	// serve
	log.Printf("Starting serve websocket on host: %s port: %d \n", cfg.Websocket.Host, cfg.Websocket.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Websocket.Host, cfg.Websocket.Port), handlers))
}
