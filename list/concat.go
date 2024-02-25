package list

import "slices"

// Concat accepts any number of slices and concatenates them into a single slice.
func Concat[T any](values ...[]T) []T {
	return slices.Concat(values...)
}
