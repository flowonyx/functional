package functional

func Transpose[T any](input [][]T) [][]T {
	output := ZeroCreate2D[T](Len2(input), len(input))

	Iteri2D(func(i1, i2 int, t T) {
		output[i2][i1] = t
	}, input)

	return output
}
