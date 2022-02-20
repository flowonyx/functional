package list

func InitSlice[T any](count uint64, initializer func(int) T) []T {
	output := make([]T, count)
	for i := range output {
		output[i] = initializer(i)
	}
	return output
}

func InitSlice2D[T any](initializer func(int, int) T, length1, length2 int) [][]T {
	output := ZeroCreate2D[T](length1, length2)
	Iteri2D(func(i1, i2 int, _ T) { output[i1][i2] = initializer(i1, i2) }, output)
	return output
}

func InitSlice3D[T any](initializer func(int, int, int) T, length1, length2, length3 int) [][][]T {
	output := ZeroCreate3D[T](length1, length2, length3)
	Iteri3D(func(i1, i2, i3 int, _ T) { output[i1][i2][i3] = initializer(i1, i2, i3) }, output)
	return output
}

func InitSlice4D[T any](initializer func(int, int, int, int) T, length1, length2, length3, length4 int) [][][][]T {
	output := ZeroCreate4D[T](length1, length2, length3, length4)
	Iteri4D(func(i1, i2, i3, i4 int, _ T) { output[i1][i2][i3][i4] = initializer(i1, i2, i3, i4) }, output)
	return output
}
