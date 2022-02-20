package list

func Empty[T any](cap ...int) []T {
	if len(cap) > 0 {
		return make([]T, cap[0])[:0]
	}
	return make([]T, 0)
}

func IsEmpty[T any](input []T) bool {
	return len(input) == 0
}
