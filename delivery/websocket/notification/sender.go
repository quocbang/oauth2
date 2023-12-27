package notification

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/quocbang/oauth2/repository/orm/models"
)

func PushNotification(n models.Notifications) error {
	if reflect.DeepEqual(n, models.Notifications{}) {
		return fmt.Errorf("nil notify message")
	}
	msg <- n // send message to channel
	return nil
}

func (ws *ws) write() {
	for {
		receiveMessage := <-msg // channel always read message

		if receiveMessage.IsSendAll() {
			go ws.SendForAllClient(receiveMessage)
		} else {
			specifyClient := map[uuid.UUID]struct{}{}
			for _, receiver := range receiveMessage.Receiver {
				specifyClient[receiver] = struct{}{}
			}
			go ws.SendSpecifyClient(specifyClient, receiveMessage)
		}
	}
}

// SendForAllClient is send notification to all client that subscribe websocket channel
func (ws *ws) SendForAllClient(message models.Notifications) {
	for _, p := range ws.participants {
		go func(client *websocket.Conn) {
			if err := client.WriteJSON(message); err != nil {
				defer client.Close()
				ws.Leave(client) // kick this client if write failed
				return           // stop this process
			}
		}(p.Conn)
	}
}

// SendSpecifyClient is send notification to specify client
func (ws *ws) SendSpecifyClient(clients map[uuid.UUID]struct{}, message models.Notifications) {
	for userID := range clients {
		if participant, ok := ws.participants[userID]; ok {
			go func(client *websocket.Conn) {
				if err := participant.Conn.WriteJSON(message); err != nil {
					defer client.Close()
					ws.Leave(client) // kick this client if write failed
					return           // stop this process
				}
			}(participant.Conn)
		}
	}
}
