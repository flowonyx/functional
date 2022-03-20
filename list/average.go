package list

import (
	"golang.org/x/exp/constraints"
)

type numeric interface {
	constraints.Float | constraints.Integer
}

// Average calculates the average of all provided values.
func Average[T numeric](values ...T) T {
	if len(values) == 0 {
		panic("Average cannot operate on empty list of values")
	}
	sum := Sum(values)
	return sum / T(len(values))
}

// AverageBy applies projection to each value to get the numeric value to be used in the average calculation.
func AverageBy[T any, R numeric](projection func(T) R, values ...T) R {
	if len(values) == 0 {
		panic("AverageBy cannot operate on empty list of values")
	}
	r := Map(projection, values)
	return Average(r...)
}
