package list

func Concat[T any](inputs ...[]T) []T {
	if len(inputs) == 0 {
		return []T{}
	}

	size := Reduce(0, func(acc int, each []T) int {
		return acc + len(each)
	}, inputs)

	output := Empty[T](size)

	Iter2D(func(t T) {
		output = append(output, t)
	}, inputs)

	return output
}
