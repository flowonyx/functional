package list

import (
	"fmt"

	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/option"
)

// Singleton creates a slice with one value.
func Singleton[T any](value T) []T {
	return []T{value}
}

// ExactlyOne returns the single item in a slice.
// If the slice is empty or has more than one item
// it will return a BadArgumentErr.
func ExactlyOne[T any](values []T) (T, error) {
	if len(values) != 1 {
		return *(new(T)), fmt.Errorf("%w: ExactlyOne([%d]%T)", errors.BadArgumentErr, len(values), values)
	}
	return values[0], nil
}

// TryExactlyOne returns the single item in a slice as an Option.
// If the slice is empty or has more than one item
// it will return None.
func TryExactlyOne[T any](input []T) option.Option[T] {
	if item, err := ExactlyOne(input); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}
