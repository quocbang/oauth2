package auth

import (
	"golang.org/x/oauth2"

	aGoogle "github.com/quocbang/oauth2/usecase/auth/google"
	"github.com/quocbang/oauth2/usecase/service"
)

type Auth struct {
	google service.IOAuth2
	github service.IOAuth2
}

func NewAuth(google oauth2.Config, github oauth2.Config) *Auth {
	return &Auth{
		google: aGoogle.NewGoogleOauth2(google.Endpoint),
		// add more here
	}
}

func (a *Auth) Google() service.IOAuth2 {
	return a.google
}
