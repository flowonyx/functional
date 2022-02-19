package functional

import (
	"github.com/flowonyx/functional/math"
	"github.com/flowonyx/functional/option"
)

func checker(start, end, step int) (stepOut int, check func(i int) bool) {
	if start < end {
		return math.Abs(step), func(i int) bool { return i <= end }
	}
	return -math.Abs(step), func(i int) bool { return i >= end }
}

func Range(start, end int, step ...int) []int {
	st := option.ValueOrDefault(TryHead(step), 1)
	if st == 0 {
		st = 1
	}

	st, check := checker(start, end, st)

	ii := Empty[int](math.Abs(start - end))
	for i := start; check(i); i += st {
		ii = append(ii, i)
	}

	return ii
}

func RangeTo(end int) []int {
	return Range(0, end)
}

func RangeChan(start, end int, step ...int) <-chan int {
	output := make(chan int, 1)
	st := option.ValueOrDefault(TryHead(step), 1)
	if st == 0 {
		st = 1
	}

	st, check := checker(start, end, st)

	go func() {
		for i := start; check(i); i += st {
			output <- i
		}
		close(output)
	}()
	return output
}

func DoRange(f func(int), start, end int, step ...int) {
	st := option.ValueOrDefault(TryHead(step), 1)
	st, check := checker(start, end, st)
	for i := start; check(i); i += st {
		f(i)
	}
}

func DoRangeTo(f func(int), end int) {
	st, check := checker(0, end, 1)
	for i := 0; check(i); i += st {
		f(i)
	}
}

func DoRangeToRev(f func(int), end int) {
	st, check := checker(end, 0, 1)
	for i := end; check(i); i += st {
		f(i)
	}
}

func DoRangeUntil(f func(int) bool, start, end int, step ...int) {
	st := option.ValueOrDefault(TryHead(step), 1)
	st, check := checker(start, end, st)
	for i := start; check(i); i += st {
		if f(i) {
			return
		}
	}
}
