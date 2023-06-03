package basicChannel

import (
	"fmt"
)

type Server struct {
	users  map[string]string
	userch chan string
	quitch chan struct{}
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
free:
	for {
		select {
		case msg := <-s.userch:
			fmt.Println(msg)
		case <-s.quitch:
			fmt.Println("server need to quit")
			break free
		}
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func sendMessage(msgch chan<- string) {
	msgch <- "hello!"
}

func readMessage(msgch chan string) {
	msg := msgch
	fmt.Println(msg)
}
