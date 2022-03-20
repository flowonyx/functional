package list

import "golang.org/x/exp/slices"

// Except returns values that are not in itemsToExclude.
func Except[T comparable](itemsToExclude []T, values ...T) []T {
	return Filter(func(t T) bool {
		return !slices.Contains(itemsToExclude, t)
	}, values...)
}
