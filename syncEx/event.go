package syncEx

import (
	"sync"
	"time"
)

// Synchronisation event
// Event can be set or reset
// Wait will block until event is set
type Event interface {
	// Set the event
	// returns true if event was set (i.e. wasn't already set)
	Set() bool
	// Reset the event
	// returns true if event was reset (i.e. wasn't already reset)
	Reset() bool
	// Wait until event is set
	Wait(timeout time.Duration) bool
}

func NewEvent(initialValue bool) Event {
	return newEvent(false, initialValue)
}

func NewAutoResetEvent(initialValue bool) Event {
	return newEvent(true, initialValue)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type event struct {
	autoReset   bool
	set         bool
	accessMutex sync.Mutex
	waitMutex   sync.Mutex
}

func (e *event) Set() bool {
	return e.changeState(true)
}

func (e *event) Reset() bool {
	return e.changeState(false)
}

func (e *event) Wait(timeout time.Duration) bool {
	result := LockWithTimeout(&e.waitMutex, timeout)
	if result && !e.autoReset {
		e.waitMutex.Unlock()
	}
	return result
}

func newEvent(autoReset, initialValue bool) Event {
	ev := event{autoReset: autoReset, set: initialValue}
	if !initialValue {
		ev.waitMutex.Lock()
	}
	return &ev
}

func (e *event) changeState(set bool) bool {
	e.accessMutex.Lock()
	defer e.accessMutex.Unlock()
	if e.set != set {
		e.set = set
		if set {
			e.waitMutex.Unlock()
		} else {
			e.waitMutex.Lock()
		}
		return true
	}
	return false
}
