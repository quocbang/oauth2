package http

import (
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"

	"github.com/quocbang/oauth2/config"
	dAuth "github.com/quocbang/oauth2/delivery/http/auth"
	"github.com/quocbang/oauth2/repository/connection"
	"github.com/quocbang/oauth2/usecase"
	"github.com/quocbang/oauth2/usecase/auth"
)

type AuthConfig struct {
	Google oauth2.Config
	Github oauth2.Config
}

type HTTP struct {
	Database     config.DatabaseGroup
	Auth         AuthConfig
	InternalAuth config.InternalAuthInfo
	MigratePath  string
	// add more here
}

func (h HTTP) RegisterHTTPHandler(e *echo.Echo) {
	api := e.Group("/api")

	p := connection.Postgres{
		Address:  h.Database.Postgres.Address,
		Port:     h.Database.Postgres.Port,
		Name:     h.Database.Postgres.Name,
		Username: h.Database.Postgres.UserName,
		Password: h.Database.Postgres.Password,
	}
	options := []connection.Option{connection.WithSchema(h.Database.Postgres.Schema)}
	repo, err := connection.NewRepository(p, h.MigratePath, options...)
	if err != nil {
		log.Fatalf("failed to init repository layer , error: %v", err)
	}

	usecase := usecase.NewUsecase(
		auth.NewAuth(h.Auth.Google, h.Auth.Github, repo, h.InternalAuth),
	)

	// auth service
	dAuth.NewAuthHandler(api.Group("/auth"), usecase)
}
