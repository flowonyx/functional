package list

// Empty makes a slice with the given cap but length of 0.
func Empty[T any](cap ...int) []T {
	if len(cap) > 0 {
		return make([]T, cap[0])[:0]
	}
	return make([]T, 0)
}

// IsEmpty tests if the slice of values is empty.
func IsEmpty[T any](values []T) bool {
	return len(values) == 0
}
