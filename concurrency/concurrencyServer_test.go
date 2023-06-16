package concurrency

import (
	"testing"
	"time"
)

func TestConcurrencyServer(t *testing.T) {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.startAndListen()
	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "Hello world")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		gracefullQuitServer(s.quitch)
	}()
}
