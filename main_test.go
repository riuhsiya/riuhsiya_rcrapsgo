package mypackage

import (
	"sync"
	"testing"
)

func TestCounterWithMutex(t *testing.T) {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	if counter != 1000 {
		t.Errorf("Expected counter to be 1000, got %d", counter)
	}
}
