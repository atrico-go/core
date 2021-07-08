package unit_tests

import (
	"context"
	"testing"

	"github.com/atrico-go/core/syncEx"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_Event_ManualInitiallyReset(t *testing.T) {
	// Arrange
	event := syncEx.Event{}

	// Act
	status1 := event.IsSet()
	result1 := event.Wait(createContext())

	// Assert
	Assert(t).That(status1, is.False, "1st IsSet not set")
	Assert(t).That(result1, is.EqualTo(context.DeadlineExceeded), "1st Event reset")
}

func Test_Event_ManualInitiallyResetThenSet(t *testing.T) {
	// Arrange
	event := syncEx.Event{}

	// Act
	set := event.Set()
	status1 := event.IsSet()
	result1 := event.Wait(createContext())
	status2 := event.IsSet()
	result2 := event.Wait(createContext())

	// Assert
	Assert(t).That(set, is.True, "Set change")
	Assert(t).That(status1, is.True, "1st IsSet set")
	Assert(t).That(result1, is.Nil, "1st Event set")
	Assert(t).That(status2, is.True, "2nd IsSet set")
	Assert(t).That(result2, is.Nil, "2nd Event set")
}

func Test_Event_ManualInitiallyResetThenReset(t *testing.T) {
	// Arrange
	event := syncEx.Event{}

	// Act
	status1 := event.IsSet()
	reset := event.Reset()
	status2 := event.IsSet()
	result1 := event.Wait(createContext())

	// Assert
	Assert(t).That(status1, is.False, "1st IsSet not set")
	Assert(t).That(reset, is.False, "Reset no change")
	Assert(t).That(status2, is.False, "2nd IsSet not set")
	Assert(t).That(result1, is.EqualTo(context.DeadlineExceeded), "1st Event reset")
}

func Test_Event_ManualInitiallySet(t *testing.T) {
	// Arrange
	event := syncEx.Event{}
	event.Set()

	// Act
	status1 := event.IsSet()
	result1 := event.Wait(createContext())
	status2 := event.IsSet()
	result2 := event.Wait(createContext())

	// Assert
	Assert(t).That(status1, is.True, "1st IsSet set")
	Assert(t).That(result1, is.Nil, "1st Event set")
	Assert(t).That(status2, is.True, "2nd IsSet set")
	Assert(t).That(result2, is.Nil, "2nd Event set")
}

func Test_Event_ManualInitiallySetThenSet(t *testing.T) {
	// Arrange
	event := syncEx.Event{}
	event.Set()

	// Act
	set := event.Set()
	result1 := event.Wait(createContext())
	result2 := event.Wait(createContext())

	// Assert
	Assert(t).That(set, is.False, "Set no change")
	Assert(t).That(result1, is.Nil, "1st Event set")
	Assert(t).That(result2, is.Nil, "2nd Event set")
}

func Test_Event_ManualInitiallySetThenReset(t *testing.T) {
	// Arrange
	event := syncEx.Event{}
	event.Set()

	// Act
	reset := event.Reset()
	result1 := event.Wait(createContext())

	// Assert
	Assert(t).That(reset, is.True, "Reset change")
	Assert(t).That(result1, is.EqualTo(context.DeadlineExceeded), "1st Event reset")
}
func Test_Event_AutoInitiallyReset(t *testing.T) {
	// Arrange
	event := syncEx.Event{AutoReset: true}

	// Act
	result1 := event.Wait(createContext())

	// Assert
	Assert(t).That(result1, is.EqualTo(context.DeadlineExceeded), "1st Event reset")
}

func Test_Event_AutoInitiallyResetThenSet(t *testing.T) {
	// Arrange
	event := syncEx.Event{AutoReset: true}

	// Act
	set := event.Set()
	result1 := event.Wait(createContext())
	result2 := event.Wait(createContext())

	// Assert
	Assert(t).That(set, is.True, "Set change")
	Assert(t).That(result1, is.Nil, "1st Event set")
	Assert(t).That(result2, is.EqualTo(context.DeadlineExceeded), "2nd Event reset")
}

func Test_Event_AutoInitiallyResetThenReset(t *testing.T) {
	// Arrange
	event := syncEx.Event{AutoReset: true}

	// Act
	reset := event.Reset()
	result1 := event.Wait(createContext())

	// Assert
	Assert(t).That(reset, is.False, "Reset no change")
	Assert(t).That(result1, is.EqualTo(context.DeadlineExceeded), "1st Event reset")
}

func Test_Event_AutoInitiallySet(t *testing.T) {
	// Arrange
	event := syncEx.Event{AutoReset: true}
	event.Set()

	// Act
	result1 := event.Wait(createContext())
	result2 := event.Wait(createContext())

	// Assert
	Assert(t).That(result1, is.Nil, "1st Event set")
	Assert(t).That(result2, is.EqualTo(context.DeadlineExceeded), "2nd Event reset")
}

func Test_Event_AutoInitiallySetThenSet(t *testing.T) {
	// Arrange
	event := syncEx.Event{AutoReset: true}
	event.Set()

	// Act
	set := event.Set()
	result1 := event.Wait(createContext())
	result2 := event.Wait(createContext())

	// Assert
	Assert(t).That(set, is.False, "Set no change")
	Assert(t).That(result1, is.Nil, "1st Event set")
	Assert(t).That(result2, is.EqualTo(context.DeadlineExceeded), "2nd Event reset")
}

func Test_Event_AutoInitiallySetThenReset(t *testing.T) {
	// Arrange
	event := syncEx.Event{AutoReset: true}
	event.Set()

	// Act
	reset := event.Reset()
	result1 := event.Wait(createContext())

	// Assert
	Assert(t).That(reset, is.True, "Reset change")
	Assert(t).That(result1, is.EqualTo(context.DeadlineExceeded), "1st Event reset")
}
