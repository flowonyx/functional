package math

import (
	"math"
	"strconv"

	"github.com/flowonyx/functional/option"
	"golang.org/x/exp/constraints"
)

type numeric interface {
	constraints.Float | constraints.Integer
}

func Abs[T constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func RoundInt[T constraints.Float](x T) int {
	return int(math.Round(float64(x)))
}

func Round[T constraints.Float](x T) T {
	return T(math.Round(float64(x)))
}

func RoundToEven[T constraints.Float](x T) T {
	return T(math.RoundToEven(float64(x)))
}

func RoundToEvenInt[T constraints.Float](x T) int {
	return int(math.RoundToEven(float64(x)))
}

func Cbrt[T numeric](x T) T {
	return T(math.Cbrt(float64(x)))
}

func CopySign[T1, T2 constraints.Signed](x T1, y T2) T1 {
	return T1(math.Copysign(float64(x), float64(y)))
}

func Dim[T numeric](x, y T) T {
	return T(math.Dim(float64(x), float64(y)))
}

func Pow[T numeric](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Pow10[T constraints.Integer](x T) T {
	return T(math.Pow10(int(x)))
}

func Remainder[T numeric](x, y T) T {
	return T(math.Remainder(float64(x), float64(y)))
}

func Sqrt[T constraints.Integer](x T) T {
	return T(math.Sqrt(float64(x)))
}

func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func TryParseInt[T constraints.Integer](input string) option.Option[T] {
	r, err := strconv.Atoi(input)
	if err != nil {
		return option.None[T]()
	}
	return option.Some(T(r))
}
