package list

// Permute returns a list with all elements permuted according to the specified permutation.
func Permute[T any](indexMap func(int) int, values []T) []T {
	output := make([]T, len(values))
	Iteri(func(i int, t T) {
		newI := indexMap(i)
		output[newI] = t
	}, values)
	return output
}
