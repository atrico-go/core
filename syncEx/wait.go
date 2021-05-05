package syncEx

import (
	"sync"
	"time"
)

// An object that can be waited on
type Waitable interface {
	// Wait on this object with a timeout
	// returns true if wait returned or false if timeout
	Wait(timeout time.Duration) bool
}

// Wait for a waitgroup with a timeout
func WaitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return true // completed normally
	case <-time.After(timeout):
		return false // timed out
	}
}
