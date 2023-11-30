package auth

import (
	"golang.org/x/oauth2"

	"github.com/quocbang/oauth2/config"
	"github.com/quocbang/oauth2/repository"
	aGoogle "github.com/quocbang/oauth2/usecase/auth/google"
	"github.com/quocbang/oauth2/usecase/service"
)

type Auth struct {
	google service.IOAuth2
}

func NewAuth(
	google oauth2.Config,
	github oauth2.Config,
	repo repository.Repositories,
	internalAuth config.InternalAuthInfo,
) *Auth {
	return &Auth{
		google: aGoogle.NewGoogleOauth2(google, repo, internalAuth),
		// add more here
	}
}

func (a *Auth) Google() service.IOAuth2 {
	return a.google
}
