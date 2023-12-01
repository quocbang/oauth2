package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/quocbang/oauth2/utils/provider"
)

type TokenMaker interface {
	GenerateToken() (string, error)
	VerifyToken(string) (any, error)
}

type JWTClaimCustom struct {
	SessionID uuid.UUID
	User      UserInfo
	jwt.RegisteredClaims
}

type UserInfo struct {
	ID       uuid.UUID
	Provider provider.Provider
	Name     string
	Email    string
	Image    string
}

type JWT struct {
	SecretKey            string
	User                 UserInfo
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

type GenerateTokenReply struct {
	AccessToken       string
	RefreshToken      string
	AccessTokenClaim  JWTClaimCustom
	RefreshTokenClaim JWTClaimCustom
}

func (j JWT) GenerateToken() (*GenerateTokenReply, error) {
	if j.SecretKey == "" {
		return nil, fmt.Errorf("missing secret key")
	}

	generateTokenReply := &GenerateTokenReply{}
	// access token
	claims := JWTClaimCustom{
		SessionID: uuid.New(),
		User: UserInfo{
			ID:       j.User.ID,
			Name:     j.User.Name,
			Provider: j.User.Provider,
			Email:    j.User.Email,
			Image:    j.User.Image,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(j.AccessTokenDuration),
			},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return nil, err
	}
	generateTokenReply.AccessToken = accessToken
	generateTokenReply.AccessTokenClaim = claims

	// refresh token
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(j.RefreshTokenDuration),
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return nil, err
	}
	generateTokenReply.RefreshToken = refreshToken
	generateTokenReply.RefreshTokenClaim = claims

	return generateTokenReply, nil
}

func (j JWT) VerifyToken(token string) (*JWTClaimCustom, error) {
	if token == "" {
		return nil, fmt.Errorf("missing token")
	}

	if j.SecretKey == "" {
		return nil, fmt.Errorf("missing secret key")
	}

	claims := JWTClaimCustom{}
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors == jwt.ValidationErrorExpired {
				return nil, fmt.Errorf("token expired")
			}
			return nil, fmt.Errorf("invalid token")
		}
		return nil, err
	}

	return &claims, nil
}
