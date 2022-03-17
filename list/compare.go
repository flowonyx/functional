package list

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](input ...T) T {
	return minMax(func(t T, min T) bool {
		return t < min
	}, input)
}

func Max[T constraints.Ordered](input ...T) T {
	return minMax(func(t T, max T) bool {
		return t > max
	}, input)
}

func minMax[T any](test func(T, T) bool, input []T) T {
	if len(input) == 0 {
		panic("minMax cannot operate on an empty input")
	}
	if len(input) == 1 {
		return input[0]
	}
	current := input[0]
	Iter(func(t T) {
		if test(t, current) {
			current = t
		}
	}, input)

	return current
}

func MaxBy[T any, T2 constraints.Ordered](projection func(T) T2, input []T) T {
	return minMax(func(t1, t2 T) bool {
		return projection(t1) > projection(t2)
	}, input)
}

func MinBy[T any, T2 constraints.Ordered](projection func(T) T2, input []T) T {
	return minMax(func(t1, t2 T) bool {
		return projection(t1) < projection(t2)
	}, input)
}
