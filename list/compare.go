package list

import (
	"fmt"

	"github.com/flowonyx/functional"
	"github.com/flowonyx/functional/errors"
	"golang.org/x/exp/constraints"
)

// Min finds the minimum value in values.
func Min[T constraints.Ordered](values ...T) (T, error) {
	if len(values) == 0 {
		return *(new(T)), fmt.Errorf("%w: Min cannot operate on an empty values", errors.BadArgumentErr)
	}
	return minMax(func(t T, min T) bool {
		return t < min
	}, values), nil
}

// MustMin is the same as Min but panics instead of returning an error.
func MustMin[T constraints.Ordered](values ...T) T {
	return functional.Must(Min(values...))
}

// Max finds the maximum value in values.
func Max[T constraints.Ordered](values ...T) (T, error) {
	if len(values) == 0 {
		return *(new(T)), fmt.Errorf("%w: Max cannot operate on an empty values", errors.BadArgumentErr)
	}
	return minMax(func(t T, max T) bool {
		return t > max
	}, values), nil
}

// MustMax is the same as Max but panics instead of returning an error.
func MustMax[T constraints.Ordered](values ...T) T {
	return functional.Must(Max(values...))
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
func MaxBy[T any, T2 constraints.Ordered](projection func(T) T2, values ...T) (T, error) {
	if len(values) == 0 {
		return *(new(T)), fmt.Errorf("%w: MaxBy cannot operate on an empty values", errors.BadArgumentErr)
	}
	return minMax(func(t1, t2 T) bool {
		return projection(t1) > projection(t2)
	}, values), nil
}

// MustMaxBy is the same as MaxBy but panics instead of returning an error.
func MustMaxBy[T any, T2 constraints.Ordered](projection func(T) T2, values ...T) T {
	return functional.Must(MaxBy(projection, values...))
}

// MinBy returns the minimum projection(value) in values.
func MinBy[T any, T2 constraints.Ordered](projection func(T) T2, values ...T) (T, error) {
	if len(values) == 0 {
		return *(new(T)), fmt.Errorf("%w: MinBy cannot operate on an empty values", errors.BadArgumentErr)
	}
	return minMax(func(t1, t2 T) bool {
		return projection(t1) < projection(t2)
	}, values), nil
}

// MustMinBy is the same as MinBy but panics instead of returning an error.
func MustMinBy[T any, T2 constraints.Ordered](projection func(T) T2, values ...T) T {
	return functional.Must(MinBy(projection, values...))
}
