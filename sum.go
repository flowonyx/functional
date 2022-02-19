package functional

func Sum[T numeric | ~string](input []T) T {
	return Reduce(*(new(T)), func(accumulator, each T) T { return accumulator + each }, input)
}

func SumBy[T any, R numeric | ~string](projection Projection[T, R], input []T) R {
	return Reduce(*(new(R)), func(accumulator R, each T) R { return accumulator + projection(each) }, input)
}
