package uint64Ex

const MinValue = uint64(0)
const MaxValue = ^uint64(0)

func Min(i,j uint64) uint64 {
	if i < j {
		return i
	}
	return j
}

func Max(i,j uint64) uint64 {
	if i > j {
		return i
	}
	return j
}