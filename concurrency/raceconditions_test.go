package concurrency

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestDataRaceConditionsMu(t *testing.T) {
	var state int32
	var mu sync.RWMutex

	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			state += int32(i)
			mu.Unlock()
		}(i)
	}
}

func TestDataRaceConditionsAtomic(t *testing.T) {
	var state int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			//state += int32(i)
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}
