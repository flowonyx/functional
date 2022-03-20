package list

import (
	"fmt"

	"github.com/flowonyx/functional/errors"
)

// InsertAt inserts newValue into existing at the given index.
// If the index is not in the range of indexes for values, it will return a nil slice and a IndexOutOfRangeErr.
func InsertAt[T any](index int, newValue T, existing []T) ([]T, error) {
	if index < 0 || index > len(existing) {
		return nil, fmt.Errorf("%w: InsertAt(%d, _, [%d]%T)", errors.IndexOutOfRangeErr, index, len(existing), existing)
	}
	return InsertManyAt(index, []T{newValue}, existing)
}

// InsertManyAt inserts newValues into existing at the given index.
// If the index is not in the range of indexes for existing, it will return a nil slice and a IndexOutOfRangeErr.
func InsertManyAt[T any](index int, newValues []T, existing []T) ([]T, error) {
	if index < 0 || index > len(existing) {
		return nil, fmt.Errorf("%w: InsertManyAt(%d, _, [%d]%T)", errors.IndexOutOfRangeErr, index, len(existing), existing)
	}
	if index == len(existing) {
		return append(existing, newValues...), nil
	}
	return append(existing[:index], append(newValues, existing[index:]...)...), nil
}

// RemoveAt removes the item at index from values.
// If index is not in the range of indexes for values, it will return a nil slice and a IndexOutOfRangeErr.
func RemoveAt[T any](index int, values []T) ([]T, error) {
	if index < 0 || index >= len(values) {
		return nil, fmt.Errorf("%w: RemoveAt(%d, [%d]%T)", errors.IndexOutOfRangeErr, index, len(values), values)
	}
	return RemoveManyAt(index, 1, values)
}

// RemoveManyAt removes count number of items starting at index from values.
// If index is not in the range of indexes for values, it will return a nil slice and a IndexOutOfRangeErr.
// If count is larger the the number of items in values starting at index, it will only remove as many items as is in the slice.
func RemoveManyAt[T any](index int, count int, values []T) ([]T, error) {
	if index < 0 || index >= len(values) {
		return nil, fmt.Errorf("%w: RemoveManyAt(%d, [%d]%T)", errors.IndexOutOfRangeErr, index, len(values), values)
	}
	count = Min(count, len(values)-index)
	return append(values[0:index], values[index+count:]...), nil
}
