package service

import (
	"context"
)

type Product interface {
	Create(context.Context) error
}

type IOAuth2 interface {
	Login() error
}
