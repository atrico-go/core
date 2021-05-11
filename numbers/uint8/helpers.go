package uint8Ex

const MinValue = uint8(0)
const MaxValue = ^uint8(0)

func Min(i,j uint8) uint8 {
	if i < j {
		return i
	}
	return j
}

func Max(i,j uint8) uint8 {
	if i > j {
		return i
	}
	return j
}