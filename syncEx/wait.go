package syncEx

import (
	"context"
)

// An object that can be waited on
type Waitable interface {
	// Wait on this object with a context
	// returns nil if wait returned or error if cancelled/timeout
	Wait(ctx context.Context) error
}
