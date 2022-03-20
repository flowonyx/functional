package list

// Windowed returns the values in sliding windows of the size specified by windowSize.
// Each window is returned as a fresh slice.
func Windowed[T any](windowSize int, values []T) [][]T {
	output := Empty[[]T](len(values) / windowSize)
	temp := Empty[T](windowSize)
	DoRangeTo(func(i int) {
		DoRange(func(j int) {
			temp = append(temp, values[j])
		}, i, i+windowSize-1)
		output = append(output, temp)
		temp = Empty[T](windowSize)
	}, len(values)-windowSize)

	return output
}
