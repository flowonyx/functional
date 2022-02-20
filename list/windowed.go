package list

func Windowed[T any](windowSize int, input []T) [][]T {
	output := Empty[[]T](len(input) / windowSize)
	temp := Empty[T](windowSize)
	DoRangeTo(func(i int) {
		DoRange(func(j int) {
			temp = append(temp, input[j])
		}, i, i+windowSize-1)
		output = append(output, temp)
		temp = Empty[T](windowSize)
	}, len(input)-windowSize)

	return output
}
