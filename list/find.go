package list

import (
	"fmt"

	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/option"
)

// Find returns the first value that matches predicate.
// If no values match predicate, it returns a NotFoundErr.
func Find[T any](predicate func(T) bool, input ...T) (T, error) {
	v, err := findFunc(predicate, input, 0, LastIndexOf(input))
	if err != nil {
		return v, fmt.Errorf("Find(%v): %w", input, err)
	}
	return v, nil
}

// FindBack returns the last value that matches predicate.
// If no values match predicate, it returns a NotFoundErr.
func FindBack[T any](predicate func(T) bool, input ...T) (T, error) {
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

	return *(new(T)), errors.NotFoundErr
}

// TryFind returns an Option of the first value that matches predicate.
// If no values match predicate, it returns None.
func TryFind[T any](predicate func(T) bool, values ...T) option.Option[T] {
	output, err := Find(predicate, values...)
	if err != nil {
		return option.None[T]()
	}
	return option.Some(output)
}

// TryFindBack returns an Option of the last value that matches predicate.
// If no values match predicate, it returns None.
func TryFindBack[T any](predicate func(T) bool, values ...T) option.Option[T] {
	output, err := FindBack(predicate, values...)
	if err != nil {
		return option.None[T]()
	}
	return option.Some(output)
}
