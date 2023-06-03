package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBasicCase(t *testing.T) {
	now := time.Now()

	userID := 10
	userData := fetchUserData(userID)
	userRecs := fetchUserRecommendations(userID)
	userLikes := fetcUserLikes(userID)

	fmt.Println(userData)
	fmt.Println(userRecs)
	fmt.Println(userLikes)

	fmt.Println(time.Since(now))
}

func TestRoutineCase(t *testing.T) {
	now := time.Now()
	userID := 10
	respch := make(chan string, 3)
	wg := &sync.WaitGroup{}

	go fetchUserDataCh(userID, respch, wg)
	wg.Add(1)
	go fetchUserRecommendationsCh(userID, respch, wg)
	wg.Add(1)
	go fetcUserLikesCh(userID, respch, wg)
	wg.Add(1)

	wg.Wait()

	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}
	fmt.Println(time.Since(now))

}
