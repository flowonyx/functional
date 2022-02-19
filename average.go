package functional

import "golang.org/x/exp/constraints"

type numeric interface {
	constraints.Float | constraints.Integer
}

func Average[T numeric](input []T) T {
	sum := Sum(input)
	return sum / T(len(input))
}

func AverageBy[T any, R numeric](input []T, projection Projection[T, R]) R {
	r := Map(projection, input)
	return Average(r)
}
