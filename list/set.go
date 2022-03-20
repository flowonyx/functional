package list

import (
	"fmt"

	"github.com/flowonyx/functional/errors"
	"golang.org/x/exp/slices"
)

// SetItem sets values[index] to value or returns an IndexOutOfRange error
// if the index is outside the range of indexes in values.
func SetItem[T any](values []T, index int, value T) error {
	if index < 0 || index > LastIndexOf(values) {
		return fmt.Errorf("%w: SetItem([%d]%T, %d, _)", errors.IndexOutOfRangeErr, len(values), values, index)
	}
	values[index] = value
	return nil
}

// SetItem2D sets values[index1][index2] to value or returns an IndexOutOfRange error
// if any of the indexes is outside the range of indexes in values.
func SetItem2D[T any](values [][]T, index1, index2 int, value T) error {
	if index1 > LastIndexOf(values) {
		return fmt.Errorf("%w: SetItem2D([%d][_]%T, %d, %d, _)", errors.IndexOutOfRangeErr, len(values), values, index1, index2)
	}
	return SetItem(values[index1], index2, value)
}

func SetItem3D[T any](values [][][]T, index1, index2, index3 int, value T) error {
	if index1 > LastIndexOf(values) {
		return fmt.Errorf("%w: SetItem3D([%d][_][_]%T, %d, %d, %d _)", errors.IndexOutOfRangeErr, len(values), values, index1, index2, index3)
	}
	return SetItem2D(values[index1], index2, index3, value)
}

func SetItem4D[T any](values [][][][]T, index1, index2, index3, index4 int, value T) error {
	if index1 > LastIndexOf(values) {
		return fmt.Errorf("%w: SetItem3D([%d][_][_][_]%T, %d, %d, %d, %d _)", errors.IndexOutOfRangeErr, len(values), values, index1, index2, index3, index4)
	}
	return SetItem3D(values[index1], index2, index3, index4, value)
}

func UpdateAt[T any](index int, value T, values []T) ([]T, error) {
	output := slices.Clone(values)
	err := SetItem(values, index, value)
	return output, err
}

func UpdateAt2D[T any](index1, index2 int, value T, values [][]T) ([][]T, error) {
	output := slices.Clone(values)
	err := SetItem2D(values, index1, index2, value)
	return output, err
}

func UpdateAt3D[T any](index1, index2, index3 int, value T, values [][][]T) ([][][]T, error) {
	output := slices.Clone(values)
	err := SetItem3D(values, index1, index2, index3, value)
	return output, err
}

func UpdateAt4D[T any](index1, index2, index3, index4 int, value T, values [][][][]T) ([][][][]T, error) {
	output := slices.Clone(values)
	err := SetItem4D(values, index1, index2, index3, index4, value)
	return output, err
}
