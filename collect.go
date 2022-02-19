package functional

func Collect[T any, R any](projection Projection[T, []R], input []T) []R {
	mapped := Map(projection, input)
	return Concat(mapped...)
}
