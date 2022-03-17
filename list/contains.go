package list

func Contains[T comparable](search T, input []T) bool {
	return IndexOf(search, input) >= 0
}
