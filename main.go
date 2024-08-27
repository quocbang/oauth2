package main

import (
	"log"
	"sync"

	"github.com/quocbang/oauth2/cmd"
	"github.com/quocbang/oauth2/delivery/middleware"
)

func main() {
	// init logger
	err := middleware.InitLogger(false)
	if err != nil {
		log.Fatalf("failed to init logger, error: %v", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		cmd.Run()
	}()

	go func() {
		defer wg.Done()
		cmd.RunWebsocket()
	}()

	wg.Wait()
}
