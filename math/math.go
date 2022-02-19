package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

type numeric interface {
	constraints.Float | constraints.Integer
}

func Abs[T constraints.Signed](input T) T {
	if input < 0 {
		return -input
	}
	return input
}

func Round[T constraints.Float](input T) int {
	return int(math.Round(float64(input)))
}

func RoundToEven[T constraints.Float](input T) int {
	return int(math.RoundToEven(float64(input)))
}
