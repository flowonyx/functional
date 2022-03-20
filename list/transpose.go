package list

// Transpose returns the transpose of the sequence of slices.
func Transpose[T any](values [][]T) [][]T {
	output := ZeroCreate2D[T](Len2(values), len(values))

	Iteri2D(func(i1, i2 int, t T) {
		output[i2][i1] = t
	}, values)

	return output
}
