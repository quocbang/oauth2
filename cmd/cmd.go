package cmd

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/quocbang/oauth2/config"
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

	// https://www.googleapis.com/auth/userinfo.email
	// https://www.googleapis.com/auth/userinfo.profile
	// https://api.github.com/user

	// 	- allow CORDS

	// get config

	// register router

	// serve
	log.Fatal(e.Start(":3000"))
}
