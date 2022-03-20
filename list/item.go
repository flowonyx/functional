package list

import (
	"fmt"

	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/option"
)

// Item is the same as values[index] but instead of a panic
// it returns a IndexOutOfRangeErr if index is outside of the range of indexes in values.
func Item[T any](index int, values []T) (T, error) {
	if index >= len(values) {
		return *(new(T)), fmt.Errorf("%w: Item(%d, [%d]%T)", errors.IndexOutOfRangeErr, index, len(values), values)
	}
	return values[index], nil
}

// TryItem is the same as Some(values[index]) but instead of a panic
// it returns None if index is outside of the range of indexes in values.
func TryItem[T any](index int, values []T) option.Option[T] {
	if item, err := Item(index, values); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}

// Item2D is the same as values[index1][index2] but instead of a panic
// it returns a IndexOutOfRangeErr if index is outside of the range of indexes in values.
func Item2D[T any](values [][]T, index1, index2 int) (T, error) {
	if index1 >= len(values) {
		return *(new(T)), fmt.Errorf("%w: Item2D(%d, _, [%d][_]%T)", errors.IndexOutOfRangeErr, index1, len(values), values)
	}
	if index2 >= len(values[index1]) {
		return *(new(T)), fmt.Errorf("%w: Item2D(%d, %d, [_][%d]%T)", errors.IndexOutOfRangeErr, index1, index2, len(values[index1]), values)
	}
	return values[index1][index2], nil
}

// TryItem2D is the same as Some(values[index1][index2]) but instead of a panic
// it returns None if index is outside of the range of indexes in values.
func TryItem2D[T any](values [][]T, index1, index2 int) option.Option[T] {
	if item, err := Item2D(values, index1, index2); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}

// Item3D is the same as values[index1][index2][index3] but instead of a panic
// it returns a IndexOutOfRangeErr if index is outside of the range of indexes in values.
func Item3D[T any](values [][][]T, index1, index2, index3 int) (T, error) {
	if index1 >= len(values) {
		return *(new(T)), fmt.Errorf("%w: Item3D(%d, _, _, [%d][_][_]%T)", errors.IndexOutOfRangeErr, index1, len(values), values)
	}
	if index2 >= len(values[index1]) {
		return *(new(T)), fmt.Errorf("%w: Item3D(%d, %d, _, [_][%d][_]%T)", errors.IndexOutOfRangeErr, index1, index2, len(values[index1]), values)
	}
	if index3 >= len(values[index1][index2]) {
		return *(new(T)), fmt.Errorf("%w: Item3D(%d, %d, %d, [_][_][%d]%T)", errors.IndexOutOfRangeErr, index1, index2, index3, len(values[index1][index2]), values)
	}
	return values[index1][index2][index3], nil
}

// TryItem3D is the same as Some(values[index1][index2][index3]) but instead of a panic
// it returns None if index is outside of the range of indexes in values.
func TryItem3D[T any](values [][][]T, index1, index2, index3 int) option.Option[T] {
	if item, err := Item3D(values, index1, index2, index3); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}

// Item4D is the same as values[index1][index2][index3][index4] but instead of a panic
// it returns a IndexOutOfRangeErr if index is outside of the range of indexes in values.
func Item4D[T any](values [][][][]T, index1, index2, index3, index4 int) (T, error) {
	if index1 >= len(values) {
		return *(new(T)), fmt.Errorf("%w: Item4D(%d, _, _, _, [%d][_][_][_]%T)", errors.IndexOutOfRangeErr, index1, len(values), values)
	}
	if index2 >= len(values[index1]) {
		return *(new(T)), fmt.Errorf("%w: Item4D(%d, %d, _, _, [_][%d][_][_]%T)", errors.IndexOutOfRangeErr, index1, index2, len(values[index1]), values)
	}
	if index3 >= len(values[index1][index2]) {
		return *(new(T)), fmt.Errorf("%w: Item4D(%d, %d, %d, _, [_][_][%d][_]%T)", errors.IndexOutOfRangeErr, index1, index2, index3, len(values[index1][index2]), values)
	}
	if index4 >= len(values[index1][index2][index3]) {
		return *(new(T)), fmt.Errorf("%w: Item4D(%d, %d, %d, %d, [_][_][_][%d]%T)", errors.IndexOutOfRangeErr, index1, index2, index3, index4, len(values[index1][index2][index3]), values)
	}
	return Item3D(values[index1], index2, index3, index4)
}

// TryItem4D is the same as Some(values[index1][index2][index3][index4]) but instead of a panic
// it returns None if index is outside of the range of indexes in values.
func TryItem4D[T any](values [][][][]T, index1, index2, index3, index4 int) option.Option[T] {
	if item, err := Item4D(values, index1, index2, index3, index4); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}
