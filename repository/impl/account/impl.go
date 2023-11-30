package account

import (
	"context"

	"github.com/quocbang/oauth2/repository"
	"github.com/quocbang/oauth2/repository/orm/models"
	"github.com/quocbang/oauth2/utils/provider"
	"gorm.io/gorm"
)

type accountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) repository.IAccount {
	return &accountService{
		db: db,
	}
}

func (a *accountService) Create(ctx context.Context, account *models.Account) error {
	return a.db.Create(&account).Error
}

func (a *accountService) GetByProviderID(ctx context.Context, provider provider.Provider, userID string) (*models.Account, error) {
	account := models.Account{}
	if err := a.db.Where("provider = ? and provider_user_id = ?", provider, userID).Take(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
