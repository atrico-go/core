package uintEx

const MinValue = uint(0)
const MaxValue = ^uint(0)

func Min(i,j uint) uint {
	if i < j {
		return i
	}
	return j
}

func Max(i,j uint) uint {
	if i > j {
		return i
	}
	return j
}