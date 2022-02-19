package functional

import "golang.org/x/exp/slices"

func Skip[T any](count int, input []T) []T {
	return slices.Clone(input[count:])
}

func SkipWhile[T any](predicate Predicate[T], input []T) []T {
	for i := range input {
		if !predicate(input[i]) {
			return input[i:]
		}
	}
	return []T{}
}
