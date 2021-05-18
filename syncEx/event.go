package syncEx

import (
	"context"
	"sync"
)

// Synchronisation event
// Event can be set or reset
// Wait will block until event is set
type Event struct {
	AutoReset   bool
	set         bool
	accessMutex sync.Mutex
	waitMutex   sync.Mutex
	init        sync.Once
}

func (e *Event) Set() bool {
	return e.SetValue(true)
}

func (e *Event) Reset() bool {
	return e.SetValue(false)
}

func (e *Event) SetValue(value bool) bool {
	e.accessMutex.Lock()
	defer e.accessMutex.Unlock()
	e.initialise()
	if e.set != value {
		e.set = value
		if value {
			e.waitMutex.Unlock()
		} else {
			e.waitMutex.Lock()
		}
		return true
	}
	return false
}

func (e *Event) Wait(ctx context.Context) error {
	e.accessMutex.Lock()
	e.initialise()
	e.accessMutex.Unlock()
	err := LockWithContext(&e.waitMutex, ctx)
	if err == nil && !e.AutoReset {
		e.waitMutex.Unlock()
	}
	return err
}

func (e *Event) initialise() {
	e.init.Do(e.waitMutex.Lock) // Initially locked (event not set)
}
