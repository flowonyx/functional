package functional

import "github.com/flowonyx/functional/option"

func Take[T any](count int, input []T) option.Option[[]T] {
	if count > LastIndexOf(input) {
		return option.None[[]T]()
	}
	return option.Some(input[:count])
}

func TakeWhile[T any](predicate Predicate[T], input []T) []T {
	for i := range input {
		if !predicate(input[i]) {
			return input[:i]
		}
	}
	return input
}

func Truncate[T any](count int, input []T) []T {
	if count > LastIndexOf(input) {
		return input
	}
	return input[:count]
}
