package google

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

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

func NewGoogleOauth2(config oauth2.Config, repo repository.Repositories, internalAuth config.InternalAuthInfo) service.IOAuth2 {
	return &oauth2Service{
		repo:         repo,
		config:       config,
		internalAuth: internalAuth,
	}
}

func (s *oauth2Service) formatUrl() string {
	values := url.Values{}
	values.Add("client_id", s.config.ClientID)
	values.Add("response_type", "code")
	values.Add("redirect_uri", s.config.RedirectURL)
	values.Add("scope", "openid email profile")
	return fmt.Sprintf("%s?%s", s.config.Endpoint.AuthURL, values.Encode())
}

func (s *oauth2Service) Login(ctx context.Context) (*presenter.LoginResponse, error) {
	return &presenter.LoginResponse{
		Url: s.formatUrl(),
	}, nil
}

func (s *oauth2Service) Oauth2Login(ctx context.Context, code string) (*presenter.Oauth2LoginResponse, error) {
	// get google token
	googleAuth, err := auth.GetGoogleOauthToken(code, s.config)
	if err != nil {
		return nil, err
	}

	// get user info with token
	googleUserInfo, err := googleAuth.GetGoogleUserInfo()
	if err != nil {
		return nil, err
	}

	// get our account
	var account *models.Account
	account, err = s.repo.Account().GetByProviderID(ctx, *provider.Provider_GOOGLE.Enum(), googleUserInfo.ID)
	if err != nil {
		// create new user if not exist in database
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// create new account
			account = &models.Account{
				Name:           googleUserInfo.Name,
				Email:          googleUserInfo.Email,
				Provider:       provider.Provider_GOOGLE,
				ProviderUserID: googleUserInfo.ID,
				Image:          googleUserInfo.Picture,
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
		ID:                   generateTokenReply.AccessTokenClaim.User.ID,
		RefreshToken:         generateTokenReply.RefreshToken,
		ProviderRefreshToken: googleAuth.IDToken, // TODO: should fix orr find refresh token
		ClientIP:             middleware.GetClientIP(ctx),
		ClientAgent:          middleware.GetClientAgent(ctx),
		Expires:              generateTokenReply.RefreshTokenClaim.ExpiresAt.Time,
		CreatedBy:            account.ID,
	}
	if err := s.repo.Session().Create(ctx, session); err != nil {
		return nil, err
	}

	// response our server token
	return &presenter.Oauth2LoginResponse{
		SessionID:    generateTokenReply.AccessTokenClaim.SessionID,
		AccessToken:  generateTokenReply.AccessToken,
		TokenExpires: generateTokenReply.AccessTokenClaim.ExpiresAt.Time,
	}, nil
}
