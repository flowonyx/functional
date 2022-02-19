package functional

func InsertAt[T any](index int, value T, input []T) ([]T, error) {
	return InsertManyAt(index, []T{value}, input)
}

func InsertManyAt[T any](index int, value []T, input []T) ([]T, error) {
	if index < 0 || index >= len(input) {
		return nil, BadArgumentErr
	}
	return append(input[:index], append(value, input[index:]...)...), nil
}

func RemoveAt[T any](index int, input []T) ([]T, error) {
	if index < 0 || index >= len(input) {
		return nil, BadArgumentErr
	}
	return RemoveManyAt(index, 1, input)
}

func RemoveManyAt[T any](index int, count int, input []T) ([]T, error) {
	if index < 0 || index >= len(input) {
		return nil, BadArgumentErr
	}
	return append(input[0:index], input[index+count:]...), nil
}
