package list

import (
	"golang.org/x/exp/constraints"
)

// Min finds the minimum value in values.
func Min[T constraints.Ordered](values ...T) T {
	if len(values) == 0 {
		panic("Min cannot operate on an empty values")
	}
	return minMax(func(t T, min T) bool {
		return t < min
	}, values)
}

// Max finds the maximum value in values.
func Max[T constraints.Ordered](values ...T) T {
	if len(values) == 0 {
		panic("Max cannot operate on an empty values")
	}
	return minMax(func(t T, max T) bool {
		return t > max
	}, values)
}

func minMax[T any](test func(T, T) bool, values []T) T {

	if len(values) == 1 {
		return values[0]
	}
	current := values[0]
	Iter(func(t T) {
		if test(t, current) {
			current = t
		}
	}, values)

	return current
}

// MaxBy returns the maximum projection(value) in values.
func MaxBy[T any, T2 constraints.Ordered](projection func(T) T2, values ...T) T {
	if len(values) == 0 {
		panic("MaxBy cannot operate on an empty values")
	}
	return minMax(func(t1, t2 T) bool {
		return projection(t1) > projection(t2)
	}, values)
}

// MinBy returns the minimum projection(value) in values.
func MinBy[T any, T2 constraints.Ordered](projection func(T) T2, values ...T) T {
	if len(values) == 0 {
		panic("MinBy cannot operate on an empty values")
	}
	return minMax(func(t1, t2 T) bool {
		return projection(t1) < projection(t2)
	}, values)
}
