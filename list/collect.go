package list

func Collect[T any, R any](projection func(T) []R, input []T) []R {
	mapped := Map(projection, input)
	return Concat(mapped...)
}
