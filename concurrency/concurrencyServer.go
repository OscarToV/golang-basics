package concurrency

import "fmt"

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func gracefullQuitServer(quitch chan struct{}) {
	close(quitch)
}

func (s *Server) startAndListen() {
	// you can name your for loop
running:
	for {
		select {
		// block here until receive a message
		case msg := <-s.msgch:
			fmt.Printf("received message from : %s payload: %s", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("The server is doing a gracefull shutdow")
			// logic for the gracefull shutdown
			break running
		default:
		}
	}
	fmt.Println("the server is shitdown")
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "testUser",
		Payload: payload,
	}

	msgch <- msg

	fmt.Println("sending message")
}
