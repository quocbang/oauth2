package google

import (
	"golang.org/x/oauth2"

	"github.com/quocbang/oauth2/auth"
)

type oauth2Service struct {
	endPoint oauth2.Endpoint
}

func NewGoogleOauth2(endPoint oauth2.Endpoint) auth.IOAuth2 {
	return &oauth2Service{
		endPoint: endPoint,
	}
}

func (s *oauth2Service) Login() error {
	return nil
}
