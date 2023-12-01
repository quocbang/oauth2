package migrations

import (
	"errors"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func Up(db *gorm.DB, migratePath string, databaseName string) error {
	getDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed go get db, error: %v", err)
	}

	driver, err := postgres.WithInstance(getDB, &postgres.Config{MigrationsTable: "migration"})
	if err != nil {
		return fmt.Errorf("failed to instance, error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migratePath, databaseName, driver)
	if err != nil {
		return fmt.Errorf("failed to new database instance, error: %v", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	zap.L().Info("Migrate successfully")
	return nil
}
