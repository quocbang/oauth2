package cmd

import (
	"log"

	"github.com/quocbang/oauth2/auth/authservice"

	"golang.org/x/oauth2"
)

func Run() {
	gEndPoint := oauth2.Endpoint{
		AuthURL:       "https://accounts.google.com/o/oauth2/auth",
		TokenURL:      "https://oauth2.googleapis.com/token",
		DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
	}

	auth2 := authservice.NewAuth(gEndPoint)
	// login with auth2
	err := auth2.Google.Login()
	if err != nil {
		log.Fatalln(err)
	}
}
