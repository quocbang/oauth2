package notification

import "net/http"

type handlers struct {
	secretKey string
}

func NewWebsocketHandler(secretKey string) INotification {
	return &handlers{
		secretKey: secretKey,
	}
}

type INotification interface {
	ReceiveNotification(http.ResponseWriter, *http.Request)
}
