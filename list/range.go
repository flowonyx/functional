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

func RangeTo[TInt constraints.Integer](end TInt) []TInt {
	return Range(0, end)
}

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

func DoRange[TInt constraints.Integer](f func(TInt), start, end TInt, step ...int) {
	st := option.DefaultValue(1, TryHead(step))
	st, check := checker(start, end, st)
	for i := start; check(i); i += TInt(st) {
		f(i)
	}
}

func DoRangeTo[TInt constraints.Integer](f func(TInt), end TInt) {
	st, check := checker(0, end, 1)
	for i := TInt(0); check(i); i += TInt(st) {
		f(i)
	}
}

func DoRangeToRev[TInt constraints.Integer](f func(TInt), end TInt) {
	st, check := checker(end, 0, 1)
	for i := end; check(i); i += TInt(st) {
		f(i)
	}
}

func DoRangeUntil[TInt constraints.Integer](f func(TInt) bool, start, end TInt, step ...int) {
	st := option.DefaultValue(1, TryHead(step))
	st, check := checker(start, end, st)
	for i := start; check(i); i += TInt(st) {
		if f(i) {
			return
		}
	}
}
