package list

import "github.com/flowonyx/functional/errors"

func Fill[T any](input []T, startIndex int, count int, value T) error {
	if len(input) < startIndex+count {
		return errors.BadArgumentErr
	}

	for _, i := range Range(startIndex, startIndex+count-1) {
		input[i] = value
	}
	return nil
}
