package functional

func Fill[T any](input []T, startIndex int, count int, value T) error {
	if len(input) < startIndex+count {
		return BadArgumentErr
	}
	for _, i := range Range(startIndex, startIndex+count-1) {
		input[i] = value
	}
	return nil
}
