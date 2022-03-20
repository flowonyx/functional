package list

import (
	"golang.org/x/exp/slices"
)

// Partition splits the collection into two collections, containing the elements for which the given predicate returns True and False respectively.
// Element order is preserved in both of the created lists.
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
