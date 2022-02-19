package functional

func ChunkBySize[T any](chunkSize int, input []T) [][]T {
	output := Empty[[]T](len(input)/chunkSize + 1)
	temp := Empty[T](chunkSize)

	Iter(func(t T) {
		if len(temp) >= chunkSize {
			output = append(output, temp)
			temp = make([]T, chunkSize)[:0]
		}
		temp = append(temp, t)
	}, input)

	output = append(output, temp)
	return output
}
