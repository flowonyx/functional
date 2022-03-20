package list

// CountBy applies projection to each value and uses the result as the key in map of counts.
func CountBy[T any, Key comparable](projection func(T) Key, values ...T) map[Key]int {
	count := map[Key]int{}
	Iter(func(k Key) {
		count[k]++
	}, Map(projection, values))
	return count
}
