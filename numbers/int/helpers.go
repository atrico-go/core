package intEx

import (
	"github.com/atrico-go/core/numbers/uint"
)

const MinValue = -MaxValue - 1
const MaxValue = int(uintEx.MaxValue >> 1)

func Min(i,j int) int {
	if i < j {
		return i
	}
	return j
}

func Max(i,j int) int {
	if i > j {
		return i
	}
	return j
}