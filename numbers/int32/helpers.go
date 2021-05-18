package int32Ex

import (
	"github.com/atrico-go/core/numbers/uint32"
)

const MinValue = -MaxValue - 1
const MaxValue = int(uint32Ex.MaxValue >> 1)

func Min(i, j int32) int32 {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int32) int32 {
	if i > j {
		return i
	}
	return j
}
