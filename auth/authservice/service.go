package authservice

import (
	"golang.org/x/oauth2"

	"github.com/quocbang/oauth2/auth"
	"github.com/quocbang/oauth2/auth/google"
)

type Auth struct {
	Google auth.IOAuth2
}

func NewAuth(googleEndPoint oauth2.Endpoint) Auth {
	return Auth{
		Google: google.NewGoogleOauth2(googleEndPoint),
	}
}
