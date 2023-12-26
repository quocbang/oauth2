package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/quocbang/oauth2/proto/notification"
)

// Image contains images relative to notification
type Image struct {
}

// Notifications definitions
type Notifications struct {
	ID      uuid.UUID         `json:"id"`
	Kind    notification.Kind `json:"kind"`
	Type    string            `json:"type"`
	Title   string            `json:"title"`
	Content string            `json:"content"`
	Images  Image             `json:"images"`
	// receiver is nil that mean send to all user online
	Receiver  []uuid.UUID `json:"receiver"`
	CreatedAt time.Time   `json:"created_at"`
}

func (Notifications) TableName() string {
	return "notifications"
}
