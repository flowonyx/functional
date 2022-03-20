package list

// Exists tests whether any value in values matches predicate.
func Exists[T any](predicate func(T) bool, values ...T) bool {
	for i := range values {
		if predicate(values[i]) {
			return true
		}
	}
	return false
}

// Exists2 tests whether any pair of values from values1 and values2 matches predicate.
// It will only use pairs of items until the minimum length of values1 and values2.
func Exists2[T any](predicate func(T, T) bool, values1 []T, values2 []T) bool {
	min := MinLen(values1, values2)
	for _, i := range RangeTo(min - 1) {
		if predicate(values1[i], values2[i]) {
			return true
		}
	}
	return false
}
