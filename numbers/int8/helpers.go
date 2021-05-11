package int8Ex

import (
	"github.com/atrico-go/core/numbers/uint8"
)

const MinValue = -MaxValue - 1
const MaxValue = int(uint8Ex.MaxValue >> 1)

func Min(i,j int8) int8 {
	if i < j {
		return i
	}
	return j
}

func Max(i,j int8) int8 {
	if i > j {
		return i
	}
	return j
}