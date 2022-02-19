package functional

import "github.com/flowonyx/functional/option"

func Find[T any](predicate Predicate[T], input []T) (T, error) {
	return findFunc(predicate, input, 0, LastIndexOf(input))
}

func FindBack[T any](predicate Predicate[T], input []T) (T, error) {
	return findFunc(predicate, input, LastIndexOf(input), 0)
}

func findFunc[T any](predicate Predicate[T], input []T, start, end int) (T, error) {
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

	return *(new(T)), KeyNotFoundErr
}

func TryFind[T any](predicate Predicate[T], input []T) option.Option[T] {
	output, err := Find(predicate, input)
	if err != nil {
		return option.None[T]()
	}
	return option.Some(output)
}

func TryFindBack[T any](predicate Predicate[T], input []T) option.Option[T] {
	output, err := FindBack(predicate, input)
	if err != nil {
		return option.None[T]()
	}
	return option.Some(output)
}
