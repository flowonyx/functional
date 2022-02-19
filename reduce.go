package functional

func Reduce[T any, R any](inital R, f func(accumulator R, each T) R, input []T) R {
	return reduce(inital, f, input, Iter[T])
}

func reduce[T any, R any](inital R, f func(accumulator R, each T) R, input []T, iterFunc func(func(T), []T)) R {
	output := inital
	iterFunc(func(t T) {
		output = f(output, t)
	}, input)
	return output
}

func ReduceBack[T any](initial T, f func(accumulator T, each T) T, input []T) T {
	return reduce(initial, f, input, IterRev[T])
}
