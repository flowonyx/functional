package functional

import "github.com/flowonyx/functional/option"

func Item[T any](index int, input []T) (T, error) {
	if index >= len(input) {
		return *(new(T)), BadArgumentErr
	}
	return input[index], nil
}

func TryItem[T any](index int, input []T) option.Option[T] {
	if item, err := Item(index, input); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}

func Item2D[T any](input [][]T, index1, index2 int) (T, error) {
	if index1 >= len(input) {
		return *(new(T)), BadArgumentErr
	}
	return Item(index2, input[index1])
}

func TryItem2D[T any](input [][]T, index1, index2 int) option.Option[T] {
	if item, err := Item2D(input, index1, index2); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}

func Item3D[T any](input [][][]T, index1, index2, index3 int) (T, error) {
	if index1 >= len(input) {
		return *(new(T)), BadArgumentErr
	}
	return Item2D(input[index1], index2, index3)
}

func TryItem3D[T any](input [][][]T, index1, index2, index3 int) option.Option[T] {
	if item, err := Item3D(input, index1, index2, index3); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}

func Item4D[T any](input [][][][]T, index1, index2, index3, index4 int) (T, error) {
	if index1 >= len(input) {
		return *(new(T)), BadArgumentErr
	}
	return Item3D(input[index1], index2, index3, index4)
}

func TryItem4D[T any](input [][][][]T, index1, index2, index3, index4 int) option.Option[T] {
	if item, err := Item4D(input, index1, index2, index3, index4); err != nil {
		return option.None[T]()
	} else {
		return option.Some(item)
	}
}
