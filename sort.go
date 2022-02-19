package functional

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Sort[T constraints.Ordered](input []T) []T {
	c := slices.Clone(input)
	slices.Sort(c)
	return c
}

func SortDescending[T constraints.Ordered](input []T) []T {
	c := slices.Clone(input)
	slices.SortFunc(c, func(a, b T) bool { return a > b })
	return c
}

func SortBy[T any, Key constraints.Ordered](projection Projection[T, Key], input []T) []T {
	return SortWith(func(t1, t2 T) bool {
		p1 := projection(t1)
		p2 := projection(t2)
		return p1 < p2
	}, input)
}

func SortByDescending[T any, Key constraints.Ordered](projection Projection[T, Key], input []T) []T {
	return SortWith(func(t1, t2 T) bool {
		p1 := projection(t1)
		p2 := projection(t2)
		return p1 > p2
	}, input)

}

func SortWith[T any](less func(T, T) bool, input []T) []T {
	output := slices.Clone(input)
	slices.SortFunc(output, less)
	return output
}

func Reverse[T any](input []T) []T {
	output := make([]T, len(input))
	j := LastIndexOf(output)
	for i := range input {
		output[j] = input[i]
		j--
	}
	return output
}
