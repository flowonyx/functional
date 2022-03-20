package list

// ChunkBySize accepts a slice of values and returns a two dimensional slice
// where each inner slice has the length of chunkSize or smaller if at end of values.
func ChunkBySize[T any](chunkSize int, values []T) [][]T {
	output := Empty[[]T](len(values)/chunkSize + 1)
	temp := Empty[T](chunkSize)

	Iter(func(t T) {
		if len(temp) >= chunkSize {
			output = append(output, temp)
			temp = make([]T, chunkSize)[:0]
		}
		temp = append(temp, t)
	}, values)

	output = append(output, temp)
	return output
}
