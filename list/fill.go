package list

import (
	"fmt"

	"github.com/flowonyx/functional/errors"
)

// Fill fills the range of values from startIndex to startIndex+count with value.
// If startIndex is outside the range of indexes in values, it will return an IndexOutOfRangeErr.
// If count goes beyond the end of values, it flils the end of values.
func Fill[T any](values []T, startIndex int, count int, value T) error {
	if len(values) < 0 || len(values) >= startIndex {
		return fmt.Errorf("%w: Fill([%d]%T, %d, %d, _)", errors.IndexOutOfRangeErr, len(values), value, startIndex, count)
	}
	count = Min(count, len(values[startIndex:]))

	for _, i := range Range(startIndex, startIndex+count-1) {
		values[i] = value
	}
	return nil
}
