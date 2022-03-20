package list

import "golang.org/x/exp/constraints"

// InitSlice makes a new slice of the given length and initializes the values
// with the return value of initializer when given the index of the item.
func InitSlice[T any, TCount, TIndex constraints.Integer](length TCount, initializer func(TIndex) T) []T {
	output := make([]T, length)
	for i := range output {
		output[i] = initializer(TIndex(i))
	}
	return output
}

// InitSlice2D makes a new two dimensional slice of length1 for the first dimension and length2 for the second and initializes the values
// with the return value of initializer when given the indexes of the item.
func InitSlice2D[T any](initializer func(int, int) T, length1, length2 int) [][]T {
	output := ZeroCreate2D[T](length1, length2)
	Iteri2D(func(i1, i2 int, _ T) { output[i1][i2] = initializer(i1, i2) }, output)
	return output
}

// InitSlice3D makes a new three dimensional slice of length1 for the first dimension, length2 for the second, and length3 for the third
// and initializes the values with the return value of initializer when given the indexes of the item.
func InitSlice3D[T any](initializer func(int, int, int) T, length1, length2, length3 int) [][][]T {
	output := ZeroCreate3D[T](length1, length2, length3)
	Iteri3D(func(i1, i2, i3 int, _ T) { output[i1][i2][i3] = initializer(i1, i2, i3) }, output)
	return output
}

// InitSlice4D makes a new four dimensional slice of length1 for the first dimension, length2 for the second, length3 for the third, and length4 for the fourth
// and initializes the values with the return value of initializer when given the indexes of the item.
func InitSlice4D[T any](initializer func(int, int, int, int) T, length1, length2, length3, length4 int) [][][][]T {
	output := ZeroCreate4D[T](length1, length2, length3, length4)
	Iteri4D(func(i1, i2, i3, i4 int, _ T) { output[i1][i2][i3][i4] = initializer(i1, i2, i3, i4) }, output)
	return output
}
