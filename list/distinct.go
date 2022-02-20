package list

import (
	. "github.com/flowonyx/functional"
	"golang.org/x/exp/slices"
)

func Distinct[T comparable](input []T) []T {
	output := Empty[T](len(input))
	Iter(func(t T) {
		if !slices.Contains(output, t) {
			output = append(output, t)
		}
	}, input)
	slices.Clip(output)
	return output
}

func DistinctBy[T any, Key comparable](projection Projection[T, Key], input []T) []T {
	output := Empty[T](len(input))
	keys := map[Key]struct{}{}
	Iter(func(t T) {
		k := projection(t)
		if _, ok := keys[k]; !ok {
			output = append(output, t)
			keys[k] = struct{}{}
		}
	}, input)
	slices.Clip(output)
	return output
}
