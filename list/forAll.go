package list

import (
	"golang.org/x/exp/constraints"
)

// ValueEqual returns a function that tests equality of a with the value passed the returned function.
// The primary purpose of this function is for use in ForAll or other predicates.
func ValueEqual[T comparable](a T) func(T) bool {
	return func(b T) bool {
		return a == b
	}
}

// ValueEqual2 tests equality of a with b.
// The primary purpose of this function is for use in ForAll2 or other predicates that accept two parameters.
func ValueEqual2[T comparable](a, b T) bool {
	return a == b
}

// Not returns a function that tests for the negative of the predicate.
// The primary purpose of this function is for use in ForAll or other predicates.
func Not[T comparable](predicate func(T) bool) func(T) bool {
	return func(a T) bool {
		return !predicate(a)
	}
}

// Not2 returns a function that tests for the negative of the two-parameter predicate.
// The primary purpose of this function is for use in ForAll2 or other predicates that accept two parameters.
func Not2[T, T2 comparable](predicate func(T, T2) bool) func(T, T2) bool {
	return func(a T, b T2) bool {
		return !predicate(a, b)
	}
}

// GreaterThan returns a function that tests whether the value passed the returned function is greater than a.
// The primary purpose of this function is for use in ForAll or other predicates.
func GreaterThan[T constraints.Ordered](a T) func(T) bool {
	return func(b T) bool {
		return b > a
	}
}

// GreaterThan2 tests whether b is greater than a.
// The primary purpose of this function is for use in ForAll2 or other predicates that accept two parameters.
func GreaterThan2[T constraints.Ordered](a, b T) bool {
	return b > a
}

// LessThan returns a function that tests whether the value passed the returned function is less than a.
// The primary purpose of this function is for use in ForAll or other predicates.
func LessThan[T constraints.Ordered](a T) func(T) bool {
	return func(b T) bool {
		return b < a
	}
}

// LessThan2 tests whether b is less than a.
// The primary purpose of this function is for use in ForAll2 or other predicates that accept two parameters.
func LessThan2[T constraints.Ordered](a, b T) bool {
	return b < a
}

// ForAll tests whether all values match predicate.
func ForAll[T any](predicate func(T) bool, values []T) bool {
	for i := range values {
		if !predicate(values[i]) {
			return false
		}
	}
	return true
}

// ForAll2 tests whether all pairs of values from values1 and values2 match predicate.
func ForAll2[T any, T2 any](predicate func(T, T2) bool, values1 []T, values2 []T2) bool {
	min := Min(len(values1), len(values2))
	if min == 0 {
		return true
	}
	for _, i := range RangeTo(min - 1) {
		if !predicate(values1[i], values2[i]) {
			return false
		}
	}
	return true
}
