package basicChannel

import (
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {
	server := NewServer()
	server.Start()

	/* 	for i := 0; i < 10; i++ {
		go func(i int) {
			server.userch <- fmt.Sprintf("user_%d", i)
		}(i)
	} */
	go func() {
		time.Sleep(2 * time.Second)
		close(server.quitch)
	}()
	// This block
	select {}
}
