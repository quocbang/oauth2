package notification

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/quocbang/oauth2/repository/orm/models"
)

const (
	MaxWebSocketReadSize = 5 * 1024 * 1024 // 5MB
	MaxHandShakeTimeOut  = time.Second * 10
)

var msg = make(chan models.Notifications)

type Participant struct {
	Conn *websocket.Conn
}

var notificationWS = &ws{
	participants: make(map[uuid.UUID]Participant),
} // TODO: should save in another place
type ws struct {
	mutex        sync.Mutex
	participants map[uuid.UUID]Participant
}

func (ws *ws) Join(userID uuid.UUID, conn *websocket.Conn) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	ws.participants[userID] = Participant{Conn: conn}
}

func (ws *ws) Leave(conn *websocket.Conn) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	for userID, p := range ws.participants {
		if p.Conn == conn {
			delete(ws.participants, userID) // delete from map
			return                          // exit function after found leave conn already
		}
	}
}
