package list

// Sum returns the result of adding all values together.
func Sum[T numeric | ~string](values []T) T {
	return Reduce(*(new(T)), func(accumulator, each T) T { return accumulator + each }, values)
}

// SumBy returns the result of adding all results of applying projection to each value.
func SumBy[T any, R numeric | ~string](projection func(T) R, values []T) R {
	return Reduce(*(new(R)), func(accumulator R, each T) R { return accumulator + projection(each) }, values)
}
