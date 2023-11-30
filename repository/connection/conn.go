package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/quocbang/oauth2/repository"
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
		// TODO: should embed logger
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewRepository(p Postgres, opts ...Option) (repository.Repositories, error) {
	options := parseOptions(opts...)

	db, err := NewDatabase(p, options.schema)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database, error: %v", err)
	}

	return &dbConnection{
		db: db,
	}, nil
}
