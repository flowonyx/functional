package list

import (
	"fmt"

	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/math"
)

// SpitAt splits values at index into two separate slices.
// If index is out of range, it will return an IndexOutOfRangeErr.
func SplitAt[T any](index int, values []T) ([]T, []T, error) {
	if index < 0 || index >= len(values) {
		return nil, nil, fmt.Errorf("%w: SplitAt(%d, [%d]%T)", errors.IndexOutOfRangeErr, index, len(values), values)
	}
	return values[0:index], values[index:], nil
}

// SplitInto splits values into a series of slices of whatever length is necessary to have count slices.
func SplitInto[T any](count int, values []T) [][]T {
	if len(values) < 2 {
		return [][]T{values}
	}
	if len(values) < count {
		count = len(values)
	}
	if count <= 0 {
		count = 1
	}
	output := Empty[[]T](count)
	c := math.RoundToEvenInt(float64(len(values)) / float64(count))
	lastIndex := 0
	DoRange(func(i int) {
		output = append(output, values[lastIndex:i])
		lastIndex = i
	}, c, LastIndexOf(values), c)
	if len(output) < count {
		output = append(output, values[lastIndex:])
	} else {
		output[LastIndexOf(output)] = append(output[LastIndexOf(output)], values[lastIndex:]...)
	}
	return output
}
