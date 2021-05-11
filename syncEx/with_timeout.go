package syncEx

import (
	"context"
	"sync"
)

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

// Wait for a WaitGroup with a context
func WaitWithContext(wg *sync.WaitGroup, ctx context.Context) error {
	return WithContext(func() { wg.Wait() }, ctx)
}

// Lock a Mutex with a context
func LockWithContext(m *sync.Mutex, ctx context.Context) error {
	return WithContext(func() { m.Lock() }, ctx)
}
