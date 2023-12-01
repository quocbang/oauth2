package connection

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	migrations "github.com/quocbang/oauth2/migration"
	"github.com/quocbang/oauth2/repository"
	"github.com/quocbang/oauth2/repository/logging"
)

type dbConnection struct {
	db     *gorm.DB
	txFlag bool
}

type options struct {
	schema string
}

type Option func(*options)

func WithSchema(schema string) Option {
	return func(o *options) {
		o.schema = schema
	}
}

func parseOptions(opts ...Option) options {
	opt := &options{}
	for _, f := range opts {
		f(opt)
	}
	return *opt
}

type Postgres struct {
	Address  string
	Port     int
	Name     string
	Username string
	Password string
}

func NewDatabase(p Postgres, schema string) (*gorm.DB, error) {
	dsnStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		p.Address,
		p.Username,
		p.Password,
		p.Name,
		p.Port,
	)
	if schema != "" {
		dsnStr = fmt.Sprintf("%s search_path=%s", dsnStr, schema)
	}
	db, err := gorm.Open(postgres.Open(dsnStr), &gorm.Config{
		Logger: logging.NewGormLogger(),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewRepository(p Postgres, migrationPath string, opts ...Option) (repository.Repositories, error) {
	options := parseOptions(opts...)

	db, err := NewDatabase(p, options.schema)
	if err != nil {
		return nil, err
	}

	// migrate
	if err := migrations.Up(db, migrationPath, p.Name); err != nil {
		log.Fatalf("failed to up migrate, error: %v", err)
	}

	return &dbConnection{
		db: db,
	}, nil
}
