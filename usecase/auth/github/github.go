package github

import (
	"context"
	"fmt"
	"net/url"

	"golang.org/x/oauth2"

	"github.com/google/uuid"
	"github.com/quocbang/oauth2/config"
	"github.com/quocbang/oauth2/presenter"
	"github.com/quocbang/oauth2/repository"
	"github.com/quocbang/oauth2/usecase/service"
)

type oauth2Service struct {
	repo         repository.Repositories
	config       oauth2.Config
	internalAuth config.InternalAuthInfo
}

func NewGithubOauth2Service(config oauth2.Config, repo repository.Repositories, internalAuth config.InternalAuthInfo) service.IOAuth2 {
	return &oauth2Service{
		repo:         repo,
		config:       config,
		internalAuth: internalAuth,
	}
}

func (s *oauth2Service) getURL() string {
	values := url.Values{}
	values.Add("client_id", s.config.ClientID)
	values.Add("redirect_uri", s.config.RedirectURL)
	values.Add("scope", "user")
	values.Add("state", uuid.New().String())
	values.Add("allow_signup", "true")
	query := values.Encode()
	return fmt.Sprintf("%s?%s", s.config.Endpoint.AuthURL, query)
}

func (s *oauth2Service) Login(ctx context.Context) (*presenter.LoginResponse, error) {
	return &presenter.LoginResponse{
		Url: s.getURL(),
	}, nil
}

func (s *oauth2Service) Oauth2Login(ctx context.Context, code string) (*presenter.Oauth2LoginResponse, error) {
	return nil, nil
}
