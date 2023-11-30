package repository

import (
	"context"
	"database/sql"

	"github.com/quocbang/oauth2/repository/orm/models"
	"github.com/quocbang/oauth2/utils/provider"
)

type Repositories interface {
	Begin(...*sql.TxOptions) Repositories
	Rollback() error
	Commit() error
	Transaction() error
	Account() IAccount
	Session() ISession
}

type IAccount interface {
	Create(context.Context, *models.Account) error
	GetByProviderID(ctx context.Context, provider provider.Provider, id string) (*models.Account, error)
}

type ISession interface {
	Create(context.Context, *models.Session) error
}
