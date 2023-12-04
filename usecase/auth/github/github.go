package github

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"gorm.io/gorm"

	"github.com/quocbang/oauth2/config"
	"github.com/quocbang/oauth2/delivery/middleware"
	"github.com/quocbang/oauth2/presenter"
	"github.com/quocbang/oauth2/repository"
	"github.com/quocbang/oauth2/repository/orm/models"
	"github.com/quocbang/oauth2/usecase/service"
	"github.com/quocbang/oauth2/utils/auth"
	"github.com/quocbang/oauth2/utils/provider"
	"github.com/quocbang/oauth2/utils/token"
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

func (s *oauth2Service) GetAuthURL(ctx context.Context) (*presenter.GetAuthURLResponse, error) {
	return &presenter.GetAuthURLResponse{
		Url: s.getURL(),
	}, nil
}

func (s *oauth2Service) Oauth2Login(ctx context.Context, code string) (*presenter.Oauth2LoginResponse, error) {
	t, err := s.config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	githubUserInfo, err := auth.GetGithubUserInfo(ctx, t)
	if err != nil {
		return nil, err
	}

	// get our account
	var account *models.Account
	account, err = s.repo.Account().GetByProviderID(ctx, *provider.Provider_GOOGLE.Enum(), fmt.Sprintf("%d", githubUserInfo.ID))

	if err != nil {
		// create new user if not exist in database
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// create new account
			account = &models.Account{
				Name:           githubUserInfo.Name,
				Email:          "",
				Provider:       provider.Provider_GITHUB,
				ProviderUserID: fmt.Sprintf("%d", githubUserInfo.ID),
				Image:          githubUserInfo.AvatarURL,
			}
			err := s.repo.Account().Create(ctx, account)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// generate new our token
	accessTokenDuration, err := time.ParseDuration(s.internalAuth.AccessTokenDuration)
	if err != nil {
		return nil, err // TODO: should return with custom error
	}
	refreshTokenDuration, err := time.ParseDuration(s.internalAuth.RefreshTokenDuration)
	if err != nil {
		return nil, err
	}
	jwt := &token.JWT{
		SecretKey: s.internalAuth.SecretKey,
		User: token.UserInfo{
			ID:       account.ID,
			Provider: provider.Provider_GOOGLE,
			Name:     account.Name,
			Email:    account.Email,
			Image:    account.Image,
		},
		AccessTokenDuration:  accessTokenDuration,
		RefreshTokenDuration: refreshTokenDuration,
	}
	generateTokenReply, err := jwt.GenerateToken()
	if err != nil {
		return nil, err
	}

	// create a session
	session := &models.Session{
		ID:                   generateTokenReply.AccessTokenClaim.SessionID,
		RefreshToken:         generateTokenReply.RefreshToken,
		ProviderRefreshToken: t.RefreshToken,
		ClientIP:             middleware.GetClientIP(ctx),
		ClientAgent:          middleware.GetClientAgent(ctx),
		Expires:              generateTokenReply.RefreshTokenClaim.ExpiresAt.Time,
		CreatedBy:            account.ID,
	}
	if err := s.repo.Session().Create(ctx, session); err != nil {
		return nil, err
	}

	return &presenter.Oauth2LoginResponse{
		SessionID:    generateTokenReply.AccessTokenClaim.SessionID,
		AccessToken:  generateTokenReply.AccessToken,
		TokenExpires: generateTokenReply.AccessTokenClaim.ExpiresAt.Time,
	}, nil
}
