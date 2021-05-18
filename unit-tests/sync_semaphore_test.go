package unit_tests

import (
	"context"
	"testing"

	"github.com/atrico-go/core/syncEx"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_Semaphore(t *testing.T) {
	// Arrange
	semaphore := syncEx.Semaphore{Limit: 2}

	// Act
	result1 := semaphore.Wait(createContext())
	result2 := semaphore.Wait(createContext())
	result3 := semaphore.Wait(createContext())

	// Assert
	Assert(t).That(result1, is.Nil, "1st Wait ok")
	Assert(t).That(result2, is.Nil, "2nd Wait ok")
	Assert(t).That(result3, is.EqualTo(context.DeadlineExceeded), "3rd Wait timeout")
}

func Test_Semaphore_Release(t *testing.T) {
	// Arrange
	semaphore := syncEx.Semaphore{Limit: 2}

	// Act
	result1 := semaphore.Wait(createContext())
	result2 := semaphore.Wait(createContext())
	semaphore.Release()
	result3 := semaphore.Wait(createContext())
	result4 := semaphore.Wait(createContext())

	// Assert
	Assert(t).That(result1, is.Nil, "1st Wait ok")
	Assert(t).That(result2, is.Nil, "2nd Wait ok")
	Assert(t).That(result3, is.Nil, "3rd Wait ok")
	Assert(t).That(result4, is.EqualTo(context.DeadlineExceeded), "4th Wait timeout")
}

func Test_Semaphore_Status(t *testing.T) {
	// Arrange
	semaphore := syncEx.Semaphore{Limit: 2}
	var current [6]int
	var limit [6]int

	// Act
	current[0], limit[0] = semaphore.Status()
	semaphore.Wait(createContext())
	current[1], limit[1] = semaphore.Status()
	semaphore.Wait(createContext())
	current[2], limit[2] = semaphore.Status()
	semaphore.Wait(createContext())
	current[3], limit[3] = semaphore.Status()
	semaphore.Release()
	current[4], limit[4] = semaphore.Status()
	semaphore.Release()
	current[5], limit[5] = semaphore.Status()

	// Assert
	Assert(t).That(current[0], is.EqualTo(0), "Initial")
	Assert(t).That(current[1], is.EqualTo(1), "After 1 wait")
	Assert(t).That(current[2], is.EqualTo(2), "After 2 waits")
	Assert(t).That(current[3], is.EqualTo(2), "After 2 waits, 1 timeout")
	Assert(t).That(current[4], is.EqualTo(1), "After 2 waits, 1 release")
	Assert(t).That(current[5], is.EqualTo(0), "After 2 waits, 2 releases")
	for i, lim := range limit {
		Assert(t).That(lim, is.EqualTo(2), "Limit %d", i)
	}
}
