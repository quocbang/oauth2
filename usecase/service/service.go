package service

import "context"

type Product interface {
	Create(context.Context) error
}
