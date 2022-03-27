package list

// Sum returns the result of adding all values together.
func Sum[T numeric | ~string](values []T) T {
	var r T
	for _, v := range values {
		r += v
	}
	return r
}

// SumBy returns the result of adding all results of applying projection to each value.
func SumBy[T any, R numeric | ~string](projection func(T) R, values []T) R {
	var r R
	for _, v := range values {
		r += projection(v)
	}
	return r
}
