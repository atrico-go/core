package unit_tests

import (
	"testing"

	"github.com/atrico-go/core/syncEx"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_Semaphore(t *testing.T) {
	// Arrange
	semaphore := syncEx.NewSemaphore(2)

	// Act
	result1 := semaphore.Wait(timeout)
	result2 := semaphore.Wait(timeout)
	result3 := semaphore.Wait(timeout)

	// Assert
	Assert(t).That(result1, is.True, "1st Wait ok")
	Assert(t).That(result2, is.True, "2nd Wait ok")
	Assert(t).That(result3, is.False, "3rd Wait timeout")
}

func Test_Semaphore_Release(t *testing.T) {
	// Arrange
	semaphore := syncEx.NewSemaphore(2)

	// Act
	result1 := semaphore.Wait(timeout)
	result2 := semaphore.Wait(timeout)
	semaphore.Release()
	result3 := semaphore.Wait(timeout)
	result4 := semaphore.Wait(timeout)

	// Assert
	Assert(t).That(result1, is.True, "1st Wait ok")
	Assert(t).That(result2, is.True, "2nd Wait ok")
	Assert(t).That(result3, is.True, "3rd Wait ok")
	Assert(t).That(result4, is.False, "4th Wait timeout")
}

func Test_Semaphore_Status(t *testing.T) {
	// Arrange
	semaphore := syncEx.NewSemaphore(2)
	var current [6]int
	var limit [6]int

	// Act
	current[0], limit[0] = semaphore.Status()
	semaphore.Wait(timeout)
	current[1], limit[1] = semaphore.Status()
	semaphore.Wait(timeout)
	current[2], limit[2] = semaphore.Status()
	semaphore.Wait(timeout)
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
	for i,lim := range limit {
		Assert(t).That(lim, is.EqualTo(2), "Limit %d", i)
	}
}
