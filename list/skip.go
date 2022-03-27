package list

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Skip returns a clone of values starting from count.
// It is the same as values[count:] but returns it as a clone
// so that modifications to the returned slice do not affect to original slice.
func Skip[T any](count int, values []T) []T {
	if count < 0 {
		count = 0
	}
	if count > len(values) {
		panic(fmt.Sprintf("Skip cannot skip more items [%d] than exist in the slice [%d].", count, len(values)))
	}
	return slices.Clone(values[count:])
}

// SkipWhile returns a clone of values starting from where predicate returns false.
// If predicate never returns false, it returns an empty slice.
// The returned slice is a clone, so modifications to it do not affect the original slice.
func SkipWhile[T any](predicate func(T) bool, values []T) []T {
	for i := range values {
		if !predicate(values[i]) {
			return slices.Clone(values[i:])
		}
	}
	return []T{}
}
