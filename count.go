package functional

func CountBy[T any, Key comparable](projection Projection[T, Key], input []T) map[Key]int {
	count := map[Key]int{}
	Iter(func(k Key) {
		count[k]++
	}, Map(projection, input))
	return count
}
