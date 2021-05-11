package unit_tests

import (
	"context"
	"errors"
	"time"
)

var timeoutError = errors.New("timmy")

func createContext() context.Context {
	timeout := 250 * time.Millisecond
	ctx,_ := context.WithTimeout(context.Background(), timeout)
	return ctx
}


