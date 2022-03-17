package list

func Equal[T comparable](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	return ForAll2(func(i1, i2 T) bool { return i1 == i2 }, s1, s2)
}

func EqualUnordered[T comparable](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	// this is to keep duplicate items in s1 from matching to the same item in s2 and returning a false positive
	indexCount := 0
	equal := ForAll(func(item T) bool {
		i := IndexOf(item, s2)
		indexCount++
		return i >= 0
	}, s1)

	return equal && indexCount == len(s1)
}
