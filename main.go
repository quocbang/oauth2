package main

import (
	"sync"

	"github.com/quocbang/oauth2/cmd"
	"github.com/quocbang/oauth2/delivery/middleware"
)

func main() {
	// init logger
	middleware.InitLogger(false)

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
