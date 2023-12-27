package notification

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/quocbang/oauth2/delivery/middleware"
	"github.com/quocbang/oauth2/utils/token"
)

const (
	HeaderCookie = "Cookie"
)

func (h *handlers) ReceiveNotification(w http.ResponseWriter, r *http.Request) {
	var (
		claim *token.JWTClaimCustom
		err   error
	)

	upgrade := websocket.Upgrader{
		ReadBufferSize: MaxWebSocketReadSize,
		// HandshakeTimeout: MaxHandShakeTimeOut,
		CheckOrigin: func(r *http.Request) bool {
			t := token.JWT{
				SecretKey: h.secretKey,
			}
			var token string
			cookie := r.Header.Get(HeaderCookie)
			if strings.Contains(cookie, string(middleware.AuthorizationKey)) {
				tokens := strings.Split(cookie, "=")
				if len(tokens) > 1 {
					token = tokens[1]
				}
			}
			claim, err = t.VerifyToken(token)
			return err == nil
		},
	}

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte(`failed to upgrade`)) // TODO: should handle error
		log.Println(err)
		return
	}

	// join notification
	notificationWS.Join(claim.User.ID, conn)

	go notificationWS.write()
	go notificationWS.read(conn)
}

func (ws *ws) read(conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if e, ok := err.(*websocket.CloseError); ok {
				// TODO: should rehandle
				// close connection
				if e.Code == websocket.CloseNormalClosure {
					conn.Close()
					ws.Leave(conn)
					return
				}
			}
			return
		}
	}
}
