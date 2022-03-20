package list

// Reduce applies f to each value in values, threading an accumulator argument through the computation.
// Apply the function to the first two elements of the list.
// Then feed this result into the function along with the third element and so on. Return the final result.
func Reduce[T any, R any](inital R, f func(accumulator R, each T) R, values []T) R {
	return reduce(inital, f, values, Iter[T])
}

func reduce[T any, R any](inital R, f func(accumulator R, each T) R, values []T, iterFunc func(func(T), []T)) R {
	output := inital
	iterFunc(func(t T) {
		output = f(output, t)
	}, values)
	return output
}

// ReduceBack applies f to each value in values in reverse, threading an accumulator argument through the computation.
// Apply the function to the first two elements of the list.
// Then feed this result into the function along with the third element and so on. Return the final result.
func ReduceBack[T any](initial T, f func(accumulator T, each T) T, values []T) T {
	return reduce(initial, f, values, IterRev[T])
}
