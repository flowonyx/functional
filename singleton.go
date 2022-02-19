package functional

import "github.com/flowonyx/functional/option"

func Singleton[T any](value T) []T {
	return []T{value}
}

func ExactlyOne[T any](input []T) (T, error) {
	if len(input) != 1 {
		return *(new(T)), BadArgumentErr
	}
	return input[0], nil
}

func TryExactlyOne[T any](input []T) option.Option[T] {
	if item, err := ExactlyOne(input); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}
