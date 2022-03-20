package list

import (
	"github.com/flowonyx/functional/math"
	"github.com/flowonyx/functional/option"
	"golang.org/x/exp/constraints"
)

func checker[TInt constraints.Integer](start, end TInt, step int) (stepOut int, check func(i TInt) bool) {
	if start < end {
		return math.Abs(step), func(i TInt) bool { return i <= end }
	}
	return -math.Abs(step), func(i TInt) bool { return i >= end }
}

// Range creates a slice of Integers from start to end.
// If step is specified, the values will be spaced by that amount.
func Range[TInt constraints.Integer](start, end TInt, step ...int) []TInt {
	st := option.DefaultValue(1, TryHead(step))
	if st == 0 {
		st = 1
	}

	st, check := checker(start, end, st)

	ii := Empty[TInt](math.Abs(int(start) - int(end)))
	for i := start; check(i); i += TInt(st) {
		ii = append(ii, i)
	}

	return ii
}

// RangeTo creates a slice of Integers from 0 to end.
func RangeTo[TInt constraints.Integer](end TInt) []TInt {
	return Range(0, end)
}

// RangeChan creates a chan of Integers that it sends values form start to end.
// If step is specified, the values will be spaced by that amount.
func RangeChan[TInt constraints.Integer](start, end TInt, step ...int) <-chan TInt {
	output := make(chan TInt, 1)
	st := option.DefaultValue(1, TryHead(step))
	if st == 0 {
		st = 1
	}

	st, check := checker(start, end, st)

	go func() {
		for i := start; check(i); i += TInt(st) {
			output <- i
		}
		close(output)
	}()
	return output
}

// DoRange calls f repeatedly with values from start to end.
// If step is specified, the values will be spaced by that amount.
func DoRange[TInt constraints.Integer](f func(TInt), start, end TInt, step ...int) {
	st := option.DefaultValue(1, TryHead(step))
	st, check := checker(start, end, st)
	for i := start; check(i); i += TInt(st) {
		f(i)
	}
}

// DoRangeTo calls f repeatedly with values from 0 to end.
func DoRangeTo[TInt constraints.Integer](f func(TInt), end TInt) {
	st, check := checker(0, end, 1)
	for i := TInt(0); check(i); i += TInt(st) {
		f(i)
	}
}

// DoRangeToRev calls f repeatedly with values from end to 0.
func DoRangeToRev[TInt constraints.Integer](f func(TInt), end TInt) {
	st, check := checker(end, 0, 1)
	for i := end; check(i); i += TInt(st) {
		f(i)
	}
}

// DoRangeUntil calls f repeatedly with values from start to end until f returns true.
// If step is specified, the values will be spaced by that amount.
func DoRangeUntil[TInt constraints.Integer](f func(TInt) bool, start, end TInt, step ...int) {
	st := option.DefaultValue(1, TryHead(step))
	st, check := checker(start, end, st)
	for i := start; check(i); i += TInt(st) {
		if f(i) {
			return
		}
	}
}
