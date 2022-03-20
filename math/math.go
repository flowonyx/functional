// Package math provides some generic mathematical functions.
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

// Abs returns the absolute value of x.
func Abs[T constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// RoundInt returns the nearest integer as an int, rounding half away from zero.
func RoundInt[T constraints.Float](x T) int {
	return int(math.Round(float64(x)))
}

// Round returns the nearest integer as the float type of x, rounding half away from zero.
func Round[T constraints.Float](x T) T {
	return T(math.Round(float64(x)))
}

// RoundToEven returns the nearest integer as the float type of x, rounding ties to even.
func RoundToEven[T constraints.Float](x T) T {
	return T(math.RoundToEven(float64(x)))
}

// RoundToEvenInt returns the nearest integer as an int, rounding ties to even.
func RoundToEvenInt[T constraints.Float](x T) int {
	return int(math.RoundToEven(float64(x)))
}

// Cbrt returns the cube root of x.
func Cbrt[T numeric](x T) T {
	return T(math.Cbrt(float64(x)))
}

// CopySign returns a value with the magnitude of x and the sign of y.
func CopySign[T1, T2 constraints.Signed](x T1, y T2) T1 {
	return T1(math.Copysign(float64(x), float64(y)))
}

// Dim returns the maximum of x-y or 0.
func Dim[T numeric](x, y T) T {
	return T(math.Dim(float64(x), float64(y)))
}

// Pow returns x**y, the base-x exponential of y.
func Pow[T numeric](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

// Pow10 returns 10**n, the base-10 exponential of n.
func Pow10[T constraints.Integer](x T) T {
	return T(math.Pow10(int(x)))
}

// Remainder returns the IEEE 754 floating-point remainder of x/y.
func Remainder[T numeric](x, y T) T {
	return T(math.Remainder(float64(x), float64(y)))
}

// Sqrt returns the square root of x.
func Sqrt[T constraints.Integer](x T) T {
	return T(math.Sqrt(float64(x)))
}

// Max returns the maximum value of x or y.
func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Min returns the minimum value of x or y.
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// TryParseInt returns the integer parsed from s as an Option.
// If parsing fails, it returns None.
func TryParseInt[T constraints.Integer](s string) option.Option[T] {
	r, err := strconv.Atoi(s)
	if err != nil {
		return option.None[T]()
	}
	return option.Some(T(r))
}
