package connection

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	"github.com/quocbang/oauth2/repository"
	"github.com/quocbang/oauth2/repository/impl/account"
	"github.com/quocbang/oauth2/repository/impl/notification"
	"github.com/quocbang/oauth2/repository/impl/session"
)

func (d *dbConnection) Account() repository.IAccount {
	return account.NewAccountService(d.db)
}

func (d *dbConnection) Session() repository.ISession {
	return session.NewSessionService(d.db)
}

func (d *dbConnection) Notification() repository.INotification {
	return notification.NewNotificationService(d.db)
}

func (d *dbConnection) Begin(opts ...*sql.TxOptions) repository.Repositories {
	if d.txFlag {
		zap.L().Info("SQL TX TRACKING...", zap.String("transaction", "Already in transaction"))
		return d
	}
	return &dbConnection{
		db:     d.db.Begin(opts...),
		txFlag: true,
	}
}

func (d *dbConnection) Rollback() error {
	if !d.txFlag {
		return fmt.Errorf("not in transaction")
	}
	return d.Rollback()
}

func (d *dbConnection) Commit() error {
	if !d.txFlag {
		return fmt.Errorf("not in transaction")
	}
	return d.Commit()
}

func (d *dbConnection) Transaction() error {
	if !d.txFlag {
		return fmt.Errorf("not in transaction")
	}
	return d.Transaction()
}
