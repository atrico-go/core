package uint32Ex

const MinValue = uint32(0)
const MaxValue = ^uint32(0)

func Min(i, j uint32) uint32 {
	if i < j {
		return i
	}
	return j
}

func Max(i, j uint32) uint32 {
	if i > j {
		return i
	}
	return j
}
