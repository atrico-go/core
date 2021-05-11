package syncEx

import (
	"context"
	"sync"
	"time"
)

// Perform action or timeout
// return true if action completed, false if timed out
func WithTimeout(action func(), timeout time.Duration) bool {
	ctx,_ := context.WithTimeout(context.Background(), timeout)
	return WithContext(action, ctx) == nil
}

// Perform action with context (cancel and/or timeout)
// return nil if action completed, error if timed out/cancelled
func WithContext(action func(), ctx context.Context) error {
	c := make(chan struct{})
	go func() {
		defer close(c)
		action()
	}()
	select {
	case <-c:
		return nil // completed normally
	case <-ctx.Done():
		return ctx.Err() // timed out/cancelled
	}
}

// Wait for a WaitGroup with a timeout
func WaitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	return WithTimeout(func() { wg.Wait() }, timeout)
}
// Wait for a WaitGroup with a context
func WaitWithContext(wg *sync.WaitGroup, ctx context.Context) error {
	return WithContext(func() { wg.Wait() }, ctx)
}

// Lock a Mutex with a timeout
func LockWithTimeout(m *sync.Mutex, timeout time.Duration) bool {
	return WithTimeout(func() { m.Lock() }, timeout)
}
// Lock a Mutex with a context
func LockWithContext(m *sync.Mutex, ctx context.Context) error {
	return WithContext(func() { m.Lock() }, ctx)
}
