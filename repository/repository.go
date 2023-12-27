package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

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
	Notification() INotification
}

type IAccount interface {
	Create(context.Context, *models.Account) error
	GetByProviderID(ctx context.Context, provider provider.Provider, id string) (*models.Account, error)
}

type ISession interface {
	Create(context.Context, *models.Session) error
}

type INotification interface {
	Create(context.Context, models.Notifications) error
	GetList(ctx context.Context, userID uuid.UUID) ([]models.Notifications, error)
}
