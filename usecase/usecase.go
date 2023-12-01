package usecase

import (
	"github.com/quocbang/oauth2/usecase/auth"
)

type UseCase struct {
	oauth2 *auth.Auth
}

func NewUsecase(oauth *auth.Auth) *UseCase {
	return &UseCase{
		oauth2: oauth,
	}
}

func (u *UseCase) Auth() *auth.Auth {
	return u.oauth2
}
