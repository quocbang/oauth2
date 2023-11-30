package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID                   uuid.UUID `json:"id"`
	RefreshToken         string    `json:"refresh_token"`
	ProviderRefreshToken string    `json:"provider_refresh_token"`
	ClientIP             string    `json:"client_id"`
	ClientAgent          string    `json:"client_agent"`
	Expires              time.Time `json:"expires"`
	CreatedAt            string    `json:"created_at"`
	CreatedBy            uuid.UUID `json:"created_by"`
}

func (Session) TableName() string {
	return "session"
}

func (s *Session) BeforeCreate(*gorm.Tx) error {
	s.ID = uuid.New()
	return nil
}
