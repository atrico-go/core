package unit_tests

import (
	"testing"
	"time"

	"github.com/atrico-go/core/syncEx"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_Event_InitiallySet(t *testing.T) {
	// Arrange
	event := syncEx.NewEvent(true)

	// Act
	result := event.Wait(time.Second)

	// Assert
	Assert(t).That(result, is.True, "Event set")
}

func Test_Event_InitiallySetThenSet(t *testing.T) {
	// Arrange
	event := syncEx.NewEvent(true)

	// Act
	set := event.Set()
	result := event.Wait(time.Second)

	// Assert
	Assert(t).That(set, is.False, "Set no change")
	Assert(t).That(result, is.True, "Event set")
}

func Test_Event_InitiallySetThenReset(t *testing.T) {
	// Arrange
	event := syncEx.NewEvent(true)

	// Act
	reset := event.Reset()
	result := event.Wait(time.Second)

	// Assert
	Assert(t).That(reset, is.True, "Reset change")
	Assert(t).That(result, is.False, "Event reset")
}

func Test_Event_InitiallyReset(t *testing.T) {
	// Arrange
	event := syncEx.NewEvent(false)

	// Act
	result := event.Wait(time.Second)

	// Assert
	Assert(t).That(result, is.False, "Event reset")
}

func Test_Event_InitiallyResetThenSet(t *testing.T) {
	// Arrange
	event := syncEx.NewEvent(false)

	// Act
	set := event.Set()
	result := event.Wait(time.Second)

	// Assert
	Assert(t).That(set, is.True, "Set change")
	Assert(t).That(result, is.True, "Event set")
}

func Test_Event_InitiallyResetThenReset(t *testing.T) {
	// Arrange
	event := syncEx.NewEvent(false)

	// Act
	reset := event.Reset()
	result := event.Wait(time.Second)

	// Assert
	Assert(t).That(reset, is.False, "Reset no change")
	Assert(t).That(result, is.False, "Event reset")
}
