package syncEx

import (
	"sync"
	"time"
)


// Perform action or timeout
// return true if action completed, false if timed out
func WithTimeout(action func(), timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		action()
	}()
	select {
	case <-c:
		return true // completed normally
	case <-time.After(timeout):
		return false // timed out
	}
}

// Wait for a WaitGroup with a timeout
func WaitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	return WithTimeout(func() { wg.Wait() }, timeout)
}

// Lock a Mutex with a timeout
func LockWithTimeout(m *sync.Mutex, timeout time.Duration) bool {
	return WithTimeout(func() { m.Lock() }, timeout)
}
