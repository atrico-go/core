package syncEx

import (
	"sync"
	"time"
)

// Semaphore synchronisation object
// Limit on number of "active" holders
// Wait will block when Limit is reached
// Release frees up one count of the Limit
type Semaphore struct {
	Limit       int
	current     int
	accessMutex sync.Mutex
	waitMutex   sync.Mutex
}

func (s *Semaphore) Wait(timeout time.Duration) bool {
	result := LockWithTimeout(&s.waitMutex, timeout)
	if result {
		s.accessMutex.Lock()
		defer s.accessMutex.Unlock()
		s.current++
		if s.current < s.Limit {
			s.waitMutex.Unlock()
		}
	}
	return result
}

func (s *Semaphore) Release() {
	s.accessMutex.Lock()
	defer s.accessMutex.Unlock()
	s.current--
	if s.current < 0 {
		panic("too many releases")
	}
	if s.current == s.Limit-1 {
		s.waitMutex.Unlock()
	}
}

func (s *Semaphore) Status() (current, limit int) {
	return s.current, s.Limit
}
