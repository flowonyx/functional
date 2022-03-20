package list

import (
	"github.com/flowonyx/functional/option"
	"golang.org/x/exp/slices"
)

// Take returns an Option of the first count values in the slice.
// If count exceeds the number of values in the slice,
// it returns None.
// To return all values when count exceeds the number of values,
// use Truncate.
func Take[T any](count int, values []T) option.Option[[]T] {
	if count > LastIndexOf(values) {
		return option.None[[]T]()
	}
	return option.Some(slices.Clone(values[:count]))
}

// TakeWhile returns the the first values in the slice until
// predicate returns false.
func TakeWhile[T any](predicate func(T) bool, values []T) []T {
	for i := range values {
		if !predicate(values[i]) {
			return slices.Clone(values[:i])
		}
	}
	return slices.Clone(values)
}

// Truncate returns the first count values in the slie.
// If count exceeds the number of values in the slice,
// it will return all values.
func Truncate[T any](count int, values []T) []T {
	if count > LastIndexOf(values) {
		return slices.Clone(values)
	}
	return slices.Clone(values[:count])
}
