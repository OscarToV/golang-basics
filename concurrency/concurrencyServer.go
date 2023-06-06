package concurrency

import "fmt"

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch chan Message
}

func (s *Server) startAndListen() {
	for {
		// block here until receive a message
		msg := <-s.msgch
		fmt.Printf("received message from : %s payload: %s", msg.From, msg.Payload)
	}
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "testUser",
		Payload: payload,
	}

	msgch <- msg

	fmt.Println("sending message")
}
