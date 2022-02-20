package list

func Len2[T any](input [][]T) int {
	if len(input) == 0 {
		return 0
	}
	return MinLen(input...)
}

func Len3[T any](input [][][]T) int {
	if len(input) == 0 {
		return 0
	}
	min := Len2(input[0])
	Iter(func(t [][]T) { min = Min(Len2(t), min) }, input)
	return min
}

func Len4[T any](input [][][][]T) int {
	if len(input) == 0 {
		return 0
	}
	min := Len3(input[0])
	Iter(func(t [][][]T) { min = Min(Len3(t), min) }, input)
	return min
}

func MinLen[T any](input ...[]T) int {
	if len(input) == 0 {
		return 0
	}
	min := len(input[0])
	Iter(func(t []T) { min = Min(len(t), min) }, input)
	return min
}

func MinSlice[T any](input ...[]T) []T {
	min := MinLen(input...)
	minIndex := 0
	DoRangeUntil(func(i int) bool {
		if len(input[i]) == min {
			minIndex = i
			return true
		}
		return false
	}, 0, LastIndexOf(input))
	return input[minIndex]
}

func LastIndexOf[T any](input []T) int {
	return len(input) - 1
}
