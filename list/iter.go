package list

// Iter iterates over the values, applying action to each value.
func Iter[T any](action func(T), values []T) {
	for i := range values {
		action(values[i])
	}
}

// IterRev iterates over the values in reverse, applying action to each value.
func IterRev[T any](action func(T), values []T) {
	DoRangeToRev(func(i int) { action(values[i]) }, LastIndexOf(values))
}

// Iter2 iterates over two slices of values, applying action to each pair of values.
// It only iterates until the end of the shortest of the value slices.
func Iter2[T, T2 any](action func(T, T2), values1 []T, values2 []T2) {
	min, _ := Min(LastIndexOf(values1), LastIndexOf(values2))
	DoRangeTo(func(i int) { action(values1[i], values2[i]) }, min)
}

// Iter2Rev iterates over two slices of values in reverse, applying action to each pair of values.
// It only iterates until the end of the shortest of the value slices.
func Iter2Rev[T, T2 any](action func(T, T2), values1 []T, values2 []T2) {
	min, _ := Min(LastIndexOf(values1), LastIndexOf(values2))
	DoRangeToRev(func(i int) { action(values1[i], values2[i]) }, min)
}

// Iter3 iterates over three slices of values, applying action to each series of values.
// It only iterates until the end of the shortest of the value slices.
func Iter3[T, T2, T3 any](action func(T, T2, T3), values1 []T, values2 []T2, values3 []T3) {
	min, _ := Min(LastIndexOf(values1), LastIndexOf(values2), LastIndexOf(values3))
	DoRangeTo(func(i int) { action(values1[i], values2[i], values3[i]) }, min)
}

// Iter3Rev iterates over three slices of values in reverse, applying action to each series of values.
// It only iterates until the end of the shortest of the value slices.
func Iter3Rev[T, T2, T3 any](action func(T, T2, T3), values1 []T, values2 []T2, values3 []T3) {
	min, _ := Min(LastIndexOf(values1), LastIndexOf(values2), LastIndexOf(values3))
	DoRangeToRev(func(i int) { action(values1[i], values2[i], values3[i]) }, min)
}

// Iteri iterates over the values, applying action to each value with the index of the value.
func Iteri[T any](action func(int, T), values []T) {
	for i := range values {
		action(i, values[i])
	}
}

// IteriRev iterates over the values in reverse, applying action to each value with the index of the value.
func IteriRev[T any](action func(int, T), values []T) {
	DoRangeToRev(func(i int) { action(i, values[i]) }, LastIndexOf(values))
}

// Iteri2 iterates over the two slices of values, applying action to each pair of values with the index of the values.
// It only iterates until the end of the shortest of the value slices.
func Iteri2[T, T2 any](action func(int, T, T2), values1 []T, values2 []T2) {
	min, _ := Min(LastIndexOf(values1), LastIndexOf(values2))
	DoRangeTo(func(i int) { action(i, values1[i], values2[i]) }, min)
}

// Iteri2Rev iterates over the two slices of values in reverse, applying action to each pair of values with the index of the values.
// It only iterates until the end of the shortest of the value slices.
func Iteri2Rev[T, T2 any](action func(int, T, T2), values1 []T, values2 []T2) {
	min, _ := Min(LastIndexOf(values1), LastIndexOf(values2))
	DoRangeToRev(func(i int) { action(i, values1[i], values2[i]) }, min)
}

// Iteri3 iterates over the three slices of values, applying action to each series of values with the index of the values.
// It only iterates until the end of the shortest of the value slices.
func Iteri3[T, T2, T3 any](action func(int, T, T2, T3), values1 []T, values2 []T2, values3 []T3) {
	min, _ := Min(LastIndexOf(values1), LastIndexOf(values2), LastIndexOf(values3))
	DoRangeTo(func(i int) { action(i, values1[i], values2[i], values3[i]) }, min)
}

// Iteri3Rev iterates over the three slices of values in reverse, applying action to each series of values with the index of the values.
// It only iterates until the end of the shortest of the value slices.
func Iteri3Rev[T, T2, T3 any](action func(int, T, T2, T3), values1 []T, values2 []T2, values3 []T3) {
	min, _ := Min(LastIndexOf(values1), LastIndexOf(values2), LastIndexOf(values3))
	DoRangeToRev(func(i int) { action(i, values1[i], values2[i], values3[i]) }, min)
}

// Iter2D iterates over a two dimensional slice of values, applying action to each value.
func Iter2D[T any](action func(T), values [][]T) {
	for i := range values {
		Iter(action, values[i])
	}
}

// Iter2DRev iterates over a two dimensional slice of values in reverse, applying action to each value.
func Iter2DRev[T any](action func(T), values [][]T) {
	IterRev(func(v []T) { IterRev(action, v) }, values)
}

// Iteri2D iterates over a two dimensional slice of values, applying action to each value with the indexes of the value.
func Iteri2D[T any](action func(int, int, T), values [][]T) {
	for i := range values {
		Iteri(func(j int, elem T) { action(i, j, elem) }, values[i])
	}
}

// Iteri2DRev iterates over a two dimensional slice of values in reverse, applying action to each value with the indexes of the value.
func Iteri2DRev[T any](action func(int, int, T), values [][]T) {
	IteriRev(func(i int, v []T) {
		IteriRev(func(j int, elem T) { action(i, j, elem) }, v)
	}, values)
}

// Iter3D iterates over a three dimensional slice of values, applying action to each value.
func Iter3D[T any](action func(T), values [][][]T) {
	for i := range values {
		Iter2D(action, values[i])
	}
}

// Iter3DRev iterates over a three dimensional slice of values in reverse, applying action to each value.
func Iter3DRev[T any](action func(T), values [][][]T) {
	IterRev(func(v [][]T) {
		Iter2DRev(action, v)
	}, values)
}

// Iteri3D iterates over a three dimensional slice of values, applying action to each value with the indexes of each value.
func Iteri3D[T any](action func(int, int, int, T), values [][][]T) {
	for i := range values {
		Iteri2D(func(j, k int, elem T) { action(i, j, k, elem) }, values[i])
	}
}

// Iteri3DRev iterates over a three dimensional slice of values in reverse, applying action to each value with the indexes of each value.
func Iteri3DRev[T any](action func(int, int, int, T), values [][][]T) {
	IteriRev(func(i int, v [][]T) {
		Iteri2DRev(func(j, k int, elem T) { action(i, j, k, elem) }, v)
	}, values)
}

// Iter4D iterates over a four dimensional slice of values, applying action to each value.
func Iter4D[T any](action func(T), values [][][][]T) {
	for i := range values {
		Iter3D(action, values[i])
	}
}

// Iter4DRev iterates over a four dimensional slice of values in reverse, applying action to each value.
func Iter4DRev[T any](action func(T), values [][][][]T) {
	IterRev(func(v [][][]T) {
		Iter3DRev(action, v)
	}, values)
}

// Iteri4D iterates over a four dimensional slice of values, applying action to each value with the indexes of each value.
func Iteri4D[T any](action func(int, int, int, int, T), values [][][][]T) {
	for i := range values {
		Iteri3D(func(j, k, l int, elem T) { action(i, j, k, l, elem) }, values[i])
	}
}

// Iteri4DRev iterates over a four dimensional slice of values in reverse, applying action to each value with the indexes of each value.
func Iteri4DRev[T any](action func(int, int, int, int, T), values [][][][]T) {
	IteriRev(func(i int, v [][][]T) {
		Iteri3DRev(func(j, k, l int, elem T) { action(i, j, k, l, elem) }, v)
	}, values)
}

// IterUntil iterates over a slice of values, applying action to each value until action returns true.
func IterUntil[T any](action func(T) bool, values []T) {
	for i := range values {
		if action(values[i]) {
			return
		}
	}
}

// IterRevUntil iterates over a slice of values in reverse, applying action to each value until action returns true.
func IterRevUntil[T any](action func(T) bool, values []T) {
	for i := LastIndexOf(values); i >= 0; i-- {
		if action(values[i]) {
			return
		}
	}
}
