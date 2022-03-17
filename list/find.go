package list

import (
	"fmt"

	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/option"
)

func Find[T any](predicate func(T) bool, input []T) (T, error) {
	v, err := findFunc(predicate, input, 0, LastIndexOf(input))
	if err != nil {
		return v, fmt.Errorf("Find(%v): %w", input, err)
	}
	return v, nil
}

func FindBack[T any](predicate func(T) bool, input []T) (T, error) {
	v, err := findFunc(predicate, input, LastIndexOf(input), 0)
	if err != nil {
		return v, fmt.Errorf("FindBack(%v): %w", input, err)
	}
	return v, nil
}

func findFunc[T any](predicate func(T) bool, input []T, start, end int) (T, error) {
	index := -1

	DoRangeUntil(func(i int) bool {
		if predicate(input[i]) {
			index = i
			return true
		}
		return false
	}, start, end)

	if index >= 0 {
		return input[index], nil
	}

	return *(new(T)), errors.KeyNotFoundErr
}

func TryFind[T any](predicate func(T) bool, input []T) option.Option[T] {
	output, err := Find(predicate, input)
	if err != nil {
		return option.None[T]()
	}
	return option.Some(output)
}

func TryFindBack[T any](predicate func(T) bool, input []T) option.Option[T] {
	output, err := FindBack(predicate, input)
	if err != nil {
		return option.None[T]()
	}
	return option.Some(output)
}
