package list

import (
	"github.com/flowonyx/functional"
	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/option"
)

func Head[T any](input []T) (T, error) {
	if len(input) > 0 {
		return input[0], nil
	}
	return *(new(T)), errors.BadArgumentErr
}

func Tail[T any](input []T) []T {
	if len(input) > 0 {
		return input[1:]
	}
	return make([]T, 0)
}

func Last[T any](input []T) (T, error) {
	if len(input) == 0 {
		return *(new(T)), errors.BadArgumentErr
	}
	return input[len(input)-1], nil
}

func TryHead[T any](input []T) option.Option[T] {
	if h, err := Head(input); err != nil {
		return option.None[T]()
	} else {
		return option.Some(h)
	}
}

func TryLast[T any](input []T) option.Option[T] {
	if l, err := Last(input); err != nil {
		return option.None[T]()
	} else {
		return option.Some(l)
	}
}

func MustHead[T any](input []T) T {
	return input[0]
}

func Cons[T any](head T, tail []T) []T {
	return append([]T{head}, tail...)
}

func ConsPair[T any](p functional.Pair[T, []T]) []T {
	head, tail := functional.FromPair(p)
	return append([]T{head}, tail...)
}
