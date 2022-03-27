package list

import (
	"fmt"

	"github.com/flowonyx/functional"
	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/option"
)

// Head returns the first item from values.
// If values contains no items, it returns the zero value for the type
// and a IndexOutOfRangeErr.
func Head[T any](values []T) (T, error) {
	if len(values) > 0 {
		return values[0], nil
	}
	return *(new(T)), fmt.Errorf("%w: Head([0]%T)", errors.IndexOutOfRangeErr, values)
}

// Tail returns all but the first item from values.
// If values contains no items, it returns the an empty slice.
func Tail[T any](values []T) []T {
	if len(values) > 0 {
		return values[1:]
	}
	return make([]T, 0)
}

// Last returns the last item from values.
// If values contains no items, it returns the zero value for the type
// and a IndexOutOfRangeErr.
func Last[T any](values []T) (T, error) {
	if len(values) > 0 {
		return values[len(values)-1], nil
	}
	return *(new(T)), fmt.Errorf("%w: Last([0]%T)", errors.IndexOutOfRangeErr, values)
}

// TryHead returns the first item from values as an Option.
// If values contains no items, it returns None.
func TryHead[T any](values []T) option.Option[T] {
	if h, err := Head(values); err != nil {
		return option.None[T]()
	} else {
		return option.Some(h)
	}
}

// TryLast returns the last item from values as an Option.
// If values contains no items, it returns None.
func TryLast[T any](values []T) option.Option[T] {
	if l, err := Last(values); err != nil {
		return option.None[T]()
	} else {
		return option.Some(l)
	}
}

// MustHead returns the first item from values.
// If values contains no items, it will panic with an index out of range.
// This should only be used when you are certain there is at least one item in values.
func MustHead[T any](values []T) T {
	if len(values) == 0 {
		panic("MustHead called with empty slice")
	}
	return values[0]
}

// MustLast returns the last item from values.
// If values contains no items, it will panic with an index out of range.
// This should only be used when you are certain there is at least one item in values.
func MustLast[T any](values []T) T {
	if len(values) == 0 {
		panic("MustLast called with empty slice")
	}
	return values[len(values)-1]
}

// Cons makes a new list with head at the beginning and the items in tail after that.
func Cons[T any](head T, tail []T) []T {
	return append([]T{head}, tail...)
}

// ConsPair makes a new list with the first item in the Pair at the beginning
// and the items in the second item in the Pair after that.
func ConsPair[T any](p functional.Pair[T, []T]) []T {
	head, tail := functional.FromPair(p)
	return append([]T{head}, tail...)
}
