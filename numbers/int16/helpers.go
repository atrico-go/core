package int16Ex

import (
	"github.com/atrico-go/core/numbers/uint16"
)

const MinValue = -MaxValue - 1
const MaxValue = int(uint16Ex.MaxValue >> 1)

func Min(i, j int16) int16 {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int16) int16 {
	if i > j {
		return i
	}
	return j
}
