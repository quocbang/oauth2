package usecase

import (
	"golang.org/x/oauth2"

	"github.com/quocbang/oauth2/usecase/auth/google"
	"github.com/quocbang/oauth2/usecase/product"
	"github.com/quocbang/oauth2/usecase/service"
)

type UseCase struct {
	Oauth2  service.Auth
	Product service.Product
}

func NewUsecase(googleEndPoint oauth2.Endpoint) UseCase {
	return UseCase{
		Oauth2: service.Auth{
			Google: google.NewGoogleOauth2(googleEndPoint),
		},
		Product: product.NewProductService(),
	}
}
