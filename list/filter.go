package list

// Filter returns the values that match predicate.
func Filter[T any](predicate func(T) bool, values ...T) []T {
	output := make([]T, len(values))

	i := 0

	Iter(func(t T) {
		if predicate(t) {
			output[i] = t
			i++
		}
	}, values)

	return output
}
