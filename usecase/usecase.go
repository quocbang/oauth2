package usecase

import (
	"github.com/quocbang/oauth2/usecase/auth"
	"github.com/quocbang/oauth2/usecase/service"
)

type UseCase struct {
	oauth2  *auth.Auth
	product service.Product
}

func NewUsecase(oauth *auth.Auth, product service.Product) *UseCase {
	return &UseCase{
		oauth2:  oauth,
		product: product,
	}
}

func (u *UseCase) Auth() *auth.Auth {
	return u.oauth2
}

func (u *UseCase) Product() service.Product {
	return u.product
}
