package syncEx

import (
	"sync"
	"time"
)

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
	ev := event{set: initialValue}
	if !initialValue {
		ev.wg.Add(1)
	}
	return &ev
}

type event struct {
	set   bool
	mutex sync.Mutex
	wg    sync.WaitGroup
}

func (e *event) Set() bool {
	return e.changeState(true)
}

func (e *event) Reset() bool {
	return e.changeState(false)
}

func (e *event) changeState(set bool) bool {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if e.set != set {
		e.set = set
		if set {
			e.wg.Done()
		} else {
			e.wg.Add(1)
		}
		return true
	}
	return false
}

func (e *event) Wait(timeout time.Duration) bool {
	return WaitWithTimeout(&e.wg, timeout)
}
