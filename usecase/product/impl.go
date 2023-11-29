package product

import (
	"context"

	"github.com/quocbang/oauth2/usecase/service"
)

type productService struct {
}

func NewProductService() service.Product {
	return &productService{}
}

func (p *productService) Create(ctx context.Context) error {
	return nil
}
