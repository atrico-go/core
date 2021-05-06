package syncEx

import (
	"sync"
	"time"
)

// Semaphore synchronisation object
// Limit on number of "active" holders
// Wait will block when limit is reached
// Release frees up one count of the limit
type Semaphore interface {
	// Wait until below limit
	// On successful wait, available is reduced by one
	Wait(timeout time.Duration) bool
	// Release the semaphore
	// Available is increased by one
	Release()
	// Current status
	Status() (current, limit int)
}

func NewSemaphore(limit int) Semaphore {
	if limit < 1 {
		panic("limit must be positive")
	}
	s := semaphore{limit: limit}
	return &s
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type semaphore struct {
	limit       int
	current     int
	accessMutex sync.Mutex
	waitMutex   sync.Mutex
}

func (s *semaphore) Wait(timeout time.Duration) bool {
	result := LockWithTimeout(&s.waitMutex, timeout)
	if result {
		s.accessMutex.Lock()
		defer s.accessMutex.Unlock()
		s.current++
		if s.current < s.limit {
			s.waitMutex.Unlock()
		}
	}
	return result
}

func (s *semaphore) Release() {
	s.accessMutex.Lock()
	defer s.accessMutex.Unlock()
	s.current--
	if s.current < 0 {
		panic("too many releases")
	}
	if s.current == s.limit-1 {
		s.waitMutex.Unlock()
	}
}

func (s *semaphore) Status() (current, limit int) {
	return s.current, s.limit
}
