package list

import (
	. "github.com/flowonyx/functional"
	"golang.org/x/exp/constraints"
)

func Equal[T comparable](a T) Predicate[T] {
	return func(b T) bool {
		return a == b
	}
}

func Equal2[T comparable](a, b T) bool {
	return a == b
}

func Not[T comparable](predicate Predicate[T]) Predicate[T] {
	return func(a T) bool {
		return !predicate(a)
	}
}

func Not2[T, T2 comparable](predicate Predicate2[T, T2]) Predicate2[T, T2] {
	return func(a T, b T2) bool {
		return !predicate(a, b)
	}
}

func GreaterThan[T constraints.Ordered](a T) Predicate[T] {
	return func(b T) bool {
		return b > a
	}
}

func GreaterThan2[T constraints.Ordered](a, b T) bool {
	return a > b
}

func LessThan[T constraints.Ordered](a T) Predicate[T] {
	return func(b T) bool {
		return b < a
	}
}

func LessThan2[T constraints.Ordered](a, b T) bool {
	return a < b
}

func ForAll[T any](predicate Predicate[T], input []T) bool {
	for i := range input {
		if !predicate(input[i]) {
			return false
		}
	}
	return true
}

func ForAll2[T any, T2 any](predicate Predicate2[T, T2], input1 []T, input2 []T2) bool {
	min := Min(len(input1), len(input2))
	for _, i := range RangeTo(min - 1) {
		if !predicate(input1[i], input2[i]) {
			return false
		}
	}
	return true
}
