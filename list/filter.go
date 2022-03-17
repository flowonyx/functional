package list

func Filter[T any](predicate func(T) bool, input []T) []T {
	output := Empty[T](len(input))

	Iter(func(t T) {
		if predicate(t) {
			output = append(output, t)
		}
	}, input)

	return output
}
