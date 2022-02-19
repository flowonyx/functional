package functional

import "github.com/flowonyx/functional/math"

func SplitAt[T any](index int, input []T) ([]T, []T) {
	if len(input) <= index {
		panic("cannot split at given index: out of range")
	}
	return input[0:index], input[index:]
}

func SplitInto[T any](count int, input []T) [][]T {
	if len(input) < count {
		count = len(input)
	}
	output := Empty[[]T](count)
	c := math.RoundToEven(float64(len(input)) / float64(count))
	lastIndex := 0
	DoRange(func(i int) {
		output = append(output, input[lastIndex:i])
		lastIndex = i
	}, c, LastIndexOf(input), c)
	output = append(output, input[lastIndex:])
	return output
}
