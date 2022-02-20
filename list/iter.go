package list

func Iter[T any](action func(T), input []T) {
	for i := range input {
		action(input[i])
	}
}

func IterRev[T any](action func(T), input []T) {
	DoRangeToRev(func(i int) { action(input[i]) }, LastIndexOf(input))
}

func Iter2[T, T2 any](action func(T, T2), input1 []T, input2 []T2) {
	min := Min(LastIndexOf(input1), LastIndexOf(input2))
	DoRangeTo(func(i int) { action(input1[i], input2[i]) }, min)
}

func Iter2Rev[T, T2 any](action func(T, T2), input1 []T, input2 []T2) {
	min := Min(LastIndexOf(input1), LastIndexOf(input2))
	DoRangeToRev(func(i int) { action(input1[i], input2[i]) }, min)
}

func Iter3[T, T2, T3 any](action func(T, T2, T3), input1 []T, input2 []T2, input3 []T3) {
	min := Min(LastIndexOf(input1), LastIndexOf(input2), LastIndexOf(input3))
	DoRangeTo(func(i int) { action(input1[i], input2[i], input3[i]) }, min)
}

func Iteri[T any](action func(int, T), input []T) {
	for i := range input {
		action(i, input[i])
	}
}

func IteriRev[T any](action func(int, T), input []T) {
	DoRangeToRev(func(i int) { action(i, input[i]) }, LastIndexOf(input))
}

func Iteri2[T, T2 any](action func(int, T, T2), input1 []T, input2 []T2) {
	min := Min(LastIndexOf(input1), LastIndexOf(input2))
	DoRangeTo(func(i int) { action(i, input1[i], input2[i]) }, min)
}

func Iteri3[T, T2, T3 any](action func(int, T, T2, T3), input1 []T, input2 []T2, input3 []T3) {
	min := Min(LastIndexOf(input1), LastIndexOf(input2), LastIndexOf(input3))
	DoRangeTo(func(i int) { action(i, input1[i], input2[i], input3[i]) }, min)
}

func Iter2D[T any](action func(T), input [][]T) {
	for i := range input {
		Iter(action, input[i])
	}
}

func Iteri2D[T any](action func(int, int, T), input [][]T) {
	for i := range input {
		Iteri(func(j int, elem T) { action(i, j, elem) }, input[i])
	}
}

func Iter3D[T any](action func(T), input [][][]T) {
	for i := range input {
		Iter2D(action, input[i])
	}
}

func Iteri3D[T any](action func(int, int, int, T), input [][][]T) {
	for i := range input {
		Iteri2D(func(j, k int, elem T) { action(i, j, k, elem) }, input[i])
	}
}

func Iter4D[T any](action func(T), input [][][][]T) {
	for i := range input {
		Iter3D(action, input[i])
	}
}

func Iteri4D[T any](action func(int, int, int, int, T), input [][][][]T) {
	for i := range input {
		Iteri3D(func(j, k, l int, elem T) { action(i, j, k, l, elem) }, input[i])
	}
}

func IterUntil[T any](action func(T) bool, input []T) {
	for i := range input {
		if action(input[i]) {
			return
		}
	}
}
