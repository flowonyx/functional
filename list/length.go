package list

// Len2 gets the minimum length of the sub slices in the two dimensional slice.
func Len2[T any](values [][]T) int {
	if len(values) == 0 {
		return 0
	}
	return MinLen(values...)
}

// Len3 gets the minimum length of the sub slices in the three dimensional slice.
func Len3[T any](values [][][]T) int {
	if len(values) == 0 {
		return 0
	}
	min := Len2(values[0])
	Iter(func(t [][]T) { min, _ = Min(Len2(t), min) }, values)
	return min
}

// Len4 gets the minimum length of the sub slices in the four dimensional slice.
func Len4[T any](values [][][][]T) int {
	if len(values) == 0 {
		return 0
	}
	min := Len3(values[0])
	Iter(func(t [][][]T) { min, _ = Min(Len3(t), min) }, values)
	return min
}

// MinLen returns the minimum length of the slices.
func MinLen[T any](values ...[]T) int {
	if len(values) == 0 {
		return 0
	}
	min := len(values[0])
	Iter(func(t []T) { min, _ = Min(len(t), min) }, values)
	return min
}

// MinSlice returns the first slice within values that has the minimum length.
func MinSlice[T any](values ...[]T) []T {
	if len(values) == 0 {
		return make([]T, 0)
	}
	min := MinLen(values...)
	minIndex := 0
	DoRangeUntil(func(i int) bool {
		if len(values[i]) == min {
			minIndex = i
			return true
		}
		return false
	}, 0, LastIndexOf(values))
	return values[minIndex]
}

// LastIndexOf is just a replacement for len(values) - 1.
func LastIndexOf[T any](values []T) int {
	return len(values) - 1
}
