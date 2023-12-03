package presenter

import (
	"time"

	"github.com/google/uuid"
)

type Oauth2LoginResponse struct {
	SessionID    uuid.UUID `json:"session_id"`
	AccessToken  string    `json:"access_token"`
	TokenExpires time.Time `json:"token_expires"`
}

type LoginResponse struct {
	Url string
}
