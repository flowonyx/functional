package list

import (
	"golang.org/x/exp/slices"
)

// Partition splits the slice into two slices. The first slice contains the items for which the given predicate returns True,
// and the second slice contains the items for which the given predicate returns False respectively.
func Partition[T any](predicate func(T) bool, values []T) (trueValues []T, falseValues []T) {
	trueValues = Empty[T](len(values))
	falseValues = Empty[T](len(values))
	for i := range values {
		if predicate(values[i]) {
			trueValues = append(trueValues, values[i])
		} else {
			falseValues = append(falseValues, values[i])
		}
	}
	slices.Clip(trueValues)
	slices.Clip(falseValues)
	return trueValues, falseValues
}
