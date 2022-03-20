package list

// Contains tests whether search is within values.
func Contains[T comparable](search T, values ...T) bool {
	return IndexOf(search, values) >= 0
}
