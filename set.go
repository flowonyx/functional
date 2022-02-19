package functional

import "golang.org/x/exp/slices"

func SetItem[T any](input []T, index int, value T) error {
	if index > LastIndexOf(input) {
		return BadArgumentErr
	}
	input[index] = value
	return nil
}

func SetItem2D[T any](input [][]T, index1, index2 int, value T) error {
	if index1 > LastIndexOf(input) {
		return BadArgumentErr
	}
	return SetItem(input[index1], index2, value)
}

func SetItem3D[T any](input [][][]T, index1, index2, index3 int, value T) error {
	if index1 > LastIndexOf(input) {
		return BadArgumentErr
	}
	return SetItem2D(input[index1], index2, index3, value)
}

func SetItem4D[T any](input [][][][]T, index1, index2, index3, index4 int, value T) error {
	if index1 > LastIndexOf(input) {
		return BadArgumentErr
	}
	return SetItem3D(input[index1], index2, index3, index4, value)
}

func UpdateAt[T any](index int, value T, input []T) ([]T, error) {
	output := slices.Clone(input)
	err := SetItem(input, index, value)
	return output, err
}

func UpdateAt2D[T any](index1, index2 int, value T, input [][]T) ([][]T, error) {
	output := slices.Clone(input)
	err := SetItem2D(input, index1, index2, value)
	return output, err
}

func UpdateAt3D[T any](index1, index2, index3 int, value T, input [][][]T) ([][][]T, error) {
	output := slices.Clone(input)
	err := SetItem3D(input, index1, index2, index3, value)
	return output, err
}

func UpdateAt4D[T any](index1, index2, index3, index4 int, value T, input [][][][]T) ([][][][]T, error) {
	output := slices.Clone(input)
	err := SetItem4D(input, index1, index2, index3, index4, value)
	return output, err
}
