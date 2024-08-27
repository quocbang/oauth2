package notification

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/quocbang/oauth2/repository"
	"github.com/quocbang/oauth2/repository/orm/models"
)

type notificationService struct {
	db *gorm.DB
}

func NewNotificationService(db *gorm.DB) repository.INotification {
	return &notificationService{
		db: db,
	}
}

func (n *notificationService) Create(ctx context.Context, notifications models.Notifications) error {
	return n.db.Create(&notifications).Error
}

func (n *notificationService) GetList(ctx context.Context, userID uuid.UUID) ([]models.Notifications, error) {
	notifications := []models.Notifications{}
	if err := n.db.Where("receiver in (?) or receiver is null", userID).Order("created_at DESC").Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}
