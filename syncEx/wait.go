package syncEx

import (
	"time"
)

// An object that can be waited on
type Waitable interface {
	// Wait on this object with a timeout
	// returns true if wait returned or false if timeout
	Wait(timeout time.Duration) bool
}
