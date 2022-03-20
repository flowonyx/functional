package list

// Concat accepts any number of slices and concatenates them into a single slice.
func Concat[T any](values ...[]T) []T {
	if len(values) == 0 {
		return []T{}
	}

	size := SumBy(func(t []T) int { return len(t) }, values)

	output := make([]T, size)

	i := 0

	Iter2D(func(t T) {
		output[i] = t
		i++
	}, values)

	return output
}
