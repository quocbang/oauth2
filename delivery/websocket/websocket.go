package websocket

import (
	"encoding/json"
	"net/http"

	"github.com/quocbang/oauth2/config"
	"github.com/quocbang/oauth2/delivery/websocket/notification"
	"github.com/quocbang/oauth2/repository/orm/models"
)

func NewWebsocketHandler(cfg config.Config) http.Handler {
	mux := http.NewServeMux()

	notificationHandler := notification.NewWebsocketHandler(cfg.InternalAuth.SecretKey)
	mux.HandleFunc("/v1/notification", notificationHandler.ReceiveNotification)
	mux.HandleFunc("/v1/posts", func(w http.ResponseWriter, r *http.Request) {
		message := models.Notifications{}
		if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
			return
		}
		err := notification.PushNotification(message)
		if err != nil {
			return
		}
	})
	return mux
}
