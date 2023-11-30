package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/quocbang/oauth2/utils/provider"
)

type Account struct {
	ID             uuid.UUID         `json:"id"`
	Name           string            `json:"name"`
	Email          string            `json:"email"`
	Provider       provider.Provider `json:"provider"`
	ProviderUserID string            `json:"provider_user_id"`
	Image          string            `json:"image"`
	CreatedAt      time.Time         `json:"created_at"`
}

func (Account) TableName() string {
	return "account"
}

func (a *Account) BeforeCreate(*gorm.DB) error {
	a.ID = uuid.New()
	return nil
}
