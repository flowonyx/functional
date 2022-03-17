package list

import (
	"golang.org/x/exp/constraints"
)

func ValueEqual[T comparable](a T) func(T) bool {
	return func(b T) bool {
		return a == b
	}
}

func ValueEqual2[T comparable](a, b T) bool {
	return a == b
}

func Not[T comparable](predicate func(T) bool) func(T) bool {
	return func(a T) bool {
		return !predicate(a)
	}
}

func Not2[T, T2 comparable](predicate func(T, T2) bool) func(T, T2) bool {
	return func(a T, b T2) bool {
		return !predicate(a, b)
	}
}

func GreaterThan[T constraints.Ordered](a T) func(T) bool {
	return func(b T) bool {
		return b > a
	}
}

func GreaterThan2[T constraints.Ordered](a, b T) bool {
	return a > b
}

func LessThan[T constraints.Ordered](a T) func(T) bool {
	return func(b T) bool {
		return b < a
	}
}

func LessThan2[T constraints.Ordered](a, b T) bool {
	return a < b
}

func ForAll[T any](predicate func(T) bool, input []T) bool {
	for i := range input {
		if !predicate(input[i]) {
			return false
		}
	}
	return true
}

func ForAll2[T any, T2 any](predicate func(T, T2) bool, input1 []T, input2 []T2) bool {
	min := Min(len(input1), len(input2))
	if min == 0 {
		return true
	}
	for _, i := range RangeTo(min - 1) {
		if !predicate(input1[i], input2[i]) {
			return false
		}
	}
	return true
}
