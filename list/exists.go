package list

import . "github.com/flowonyx/functional"

func Exists[T any](predicate Predicate[T], input []T) bool {
	for i := range input {
		if predicate(input[i]) {
			return true
		}
	}
	return false
}

func Exists2[T any](predicate Predicate2[T, T], input1 []T, input2 []T) bool {
	min := MinLen(input1, input2)
	for _, i := range RangeTo(min - 1) {
		if predicate(input1[i], input2[i]) {
			return true
		}
	}
	return false
}
