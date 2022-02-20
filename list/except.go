package list

import "golang.org/x/exp/slices"

func Except[T comparable](itemsToExclude []T, input []T) []T {
	return Filter(func(t T) bool {
		return !slices.Contains(itemsToExclude, t)
	}, input)
}
