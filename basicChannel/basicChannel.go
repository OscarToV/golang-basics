package basicChannel

import "fmt"

type Server struct {
	users  map[string]string
	userch chan string
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
	}
}

func (s *Server) Start() {

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
