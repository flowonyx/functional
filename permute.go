package functional

func Permute[T any](indexMap Projection[int, int], input []T) []T {
	output := make([]T, len(input))
	Iteri(func(i int, t T) {
		newI := indexMap(i)
		output[newI] = t
	}, input)
	return output
}
