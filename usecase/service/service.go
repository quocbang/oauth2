package service

import (
	"context"

	"github.com/quocbang/oauth2/presenter"
)

type IOAuth2 interface {
	Login(context.Context) (url string, err error)
	Oauth2Login(ctx context.Context, code string) (*presenter.Oauth2LoginResponse, error)
}
