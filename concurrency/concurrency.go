package concurrency

import (
	"sync"
	"time"
)

func fetchUserData(userID int) string {
	time.Sleep(80 * time.Millisecond)

	return "user data"
}

func fetchUserRecommendations(userID int) string {
	time.Sleep(120 * time.Millisecond)
	return "user recommendations"
}

func fetcUserLikes(userID int) string {
	time.Sleep(50 * time.Millisecond)
	return "user likes"
}

func fetchUserDataCh(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)

	respch <- "user data"

	wg.Done()
}

func fetchUserRecommendationsCh(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)

	respch <- "user recommendations"

	wg.Done()
}

func fetcUserLikesCh(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)

	respch <- "user likes"

	wg.Done()
}
