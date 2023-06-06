package concurrency

import "testing"

func TestConcurrencyServer(t *testing.T) {
	s := &Server{
		msgch: make(chan Message),
	}
	go s.startAndListen()
	sendMessageToServer(s.msgch, "Hello world")

}
