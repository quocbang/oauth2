package session

import (
	"context"

	"github.com/quocbang/oauth2/repository"
	"github.com/quocbang/oauth2/repository/orm/models"
	"gorm.io/gorm"
)

type sessionService struct {
	db *gorm.DB
}

func NewSessionService(db *gorm.DB) repository.ISession {
	return &sessionService{
		db: db,
	}
}

func (s *sessionService) Create(ctx context.Context, ss *models.Session) error {
	return s.db.Create(&ss).Error
}
