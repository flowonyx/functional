package list

// Collect applies projection to a slice of values.
// projection returns a slice for each value, which then are all concatenated into one slice.
func Collect[T any, R any](projection func(T) []R, input []T) []R {
	mapped := Map(projection, input)
	return Concat(mapped...)
}
