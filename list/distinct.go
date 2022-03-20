package list

import (
	"golang.org/x/exp/slices"
)

// Distinct returns the values without repitition.
func Distinct[T comparable](values ...T) []T {
	output := make([]T, len(values))
	i := 0
	Iter(func(t T) {
		if !slices.Contains(output, t) {
			output[i] = t
			i++
		}
	}, values)
	return slices.Clip(output[0:i])
}

// DistinctBy applies projection to each value and returns the values whose projected value is not repeated.
func DistinctBy[T any, Key comparable](projection func(T) Key, values ...T) []T {
	output := make([]T, len(values))
	keys := map[Key]struct{}{}
	i := 0
	Iter(func(t T) {
		k := projection(t)
		if _, ok := keys[k]; !ok {
			output[i] = t
			i++
			keys[k] = struct{}{}
		}
	}, values)
	return slices.Clip(output[0:i])
}
