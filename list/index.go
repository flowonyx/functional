package list

import (
	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/option"
)

// Indexed converts values into Pairs of each value with its index.
func Indexed[T any](values []T) []Pair[int, T] {
	output := make([]Pair[int, T], len(values))
	for i := range values {
		output[i] = PairOf(i, values[i])
	}
	return output
}

func indexFunc[T any](search func(T) bool, input []T, reverse bool) int {
	if len(input) == 0 {
		return -1
	}
	index := -1
	start, end := 0, len(input)-1
	if reverse {
		start, end = end, start
	}
	DoRangeUntil(func(i int) bool {
		if search(input[i]) {
			index = i
			return true
		}
		return false
	}, start, end)

	return index
}

// IndexOf returns the first index within values of search.
// If search is not in values, it returns -1.
func IndexOf[T comparable](search T, values []T) int {
	return indexFunc(func(t T) bool { return t == search }, values, false)
}

// IndexOfBack returns the last index within values of search.
// If search is not in values, it returns -1.
func IndexOfBack[T comparable](search T, input []T) int {
	return indexFunc(func(t T) bool { return t == search }, input, true)
}

// IndexBy returns the first index within values that matches the search predicate.
// If no values match search, it returns -1.
func IndexBy[T any](search func(T) bool, values []T) int {
	return indexFunc(search, values, false)
}

// IndexByBack returns the last index within values that matches the search predicate.
// If no values match search, it returns -1.
func IndexByBack[T any](search func(T) bool, values []T) int {
	return indexFunc(search, values, true)
}

// TryIndexOf returns the first index within values of search as an Option.
// If search is not in values, it returns None.
// This function is probably only useful if you working heavily with Options.
func TryIndexOf[T comparable](search T, input []T) option.Option[int] {
	if i := IndexOf(search, input); i >= 0 {
		return option.Some(i)
	}
	return option.None[int]()
}

// TryIndexOfBack returns the last index within values of search as an Option.
// If search is not in values, it returns None.
// This function is probably only useful if you working heavily with Options.
func TryIndexOfBack[T comparable](search T, input []T) option.Option[int] {
	if i := IndexOfBack(search, input); i >= 0 {
		return option.Some(i)
	}
	return option.None[int]()
}

// TryIndexBy returns the first index within values that matches the search predicate as an Option.
// If search is not in values, it returns None.
// This function is probably only useful if you working heavily with Options.
func TryIndexBy[T any](search func(T) bool, input []T) option.Option[int] {
	if i := IndexBy(search, input); i >= 0 {
		return option.Some(i)
	}
	return option.None[int]()
}

// TryIndexByBack returns the last index within values that matches the search predicate as an Option.
// If search is not in values, it returns None.
// This function is probably only useful if you working heavily with Options.
func TryIndexByBack[T any](search func(T) bool, input []T) option.Option[int] {
	if i := IndexByBack(search, input); i >= 0 {
		return option.Some(i)
	}
	return option.None[int]()
}
