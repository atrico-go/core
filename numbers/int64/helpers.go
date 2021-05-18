package int64Ex

import (
	"github.com/atrico-go/core/numbers/uint64"
)

const MinValue = -MaxValue - 1
const MaxValue = int(uint64Ex.MaxValue >> 1)

func Min(i, j int64) int64 {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}
