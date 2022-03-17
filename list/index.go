package list

import (
	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/option"
)

func Indexed[T any](input []T) []Pair[int, T] {
	output := make([]Pair[int, T], len(input))
	for i := range input {
		output[i] = PairOf(i, input[i])
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

func IndexOf[T comparable](search T, input []T) int {
	return indexFunc(func(t T) bool { return t == search }, input, false)
}

func IndexOfBack[T comparable](search T, input []T) int {
	return indexFunc(func(t T) bool { return t == search }, input, true)
}

func IndexBy[T any](search func(T) bool, input []T) int {
	return indexFunc(search, input, false)
}

func IndexByBack[T any](search func(T) bool, input []T) int {
	return indexFunc(search, input, true)
}

func TryIndexOf[T comparable](search T, input []T) option.Option[int] {
	if i := IndexOf(search, input); i >= 0 {
		return option.Some(i)
	}
	return option.None[int]()
}

func TryIndexOfBack[T comparable](search T, input []T) option.Option[int] {
	if i := IndexOfBack(search, input); i >= 0 {
		return option.Some(i)
	}
	return option.None[int]()
}

func TryIndexBy[T any](search func(T) bool, input []T) option.Option[int] {
	if i := IndexBy(search, input); i >= 0 {
		return option.Some(i)
	}
	return option.None[int]()
}

func TryIndexByBack[T any](search func(T) bool, input []T) option.Option[int] {
	if i := IndexByBack(search, input); i >= 0 {
		return option.Some(i)
	}
	return option.None[int]()
}
