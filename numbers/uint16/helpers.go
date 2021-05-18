package uint16Ex

const MinValue = uint16(0)
const MaxValue = ^uint16(0)

func Min(i, j uint16) uint16 {
	if i < j {
		return i
	}
	return j
}

func Max(i, j uint16) uint16 {
	if i > j {
		return i
	}
	return j
}
