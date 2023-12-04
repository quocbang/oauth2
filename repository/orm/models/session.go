package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID                   uuid.UUID `json:"id"`
	RefreshToken         string    `json:"refresh_token"`
	ProviderRefreshToken string    `json:"provider_refresh_token"`
	ClientIP             string    `json:"client_id"`
	ClientAgent          string    `json:"client_agent"`
	Expires              time.Time `json:"expires"`
	CreatedAt            time.Time `json:"created_at"`
	CreatedBy            uuid.UUID `json:"created_by"`
}

func (Session) TableName() string {
	return "session"
}
