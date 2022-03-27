package list

// Map applies mapping to values and returns the results as a new slice.
func Map[T, R any](mapping func(T) R, values []T) []R {
	output := make([]R, len(values))
	iter := func(i int, t T) { output[i] = mapping(t) }
	Iteri(iter, values)
	return output
}

// Map2 applies mapping to pairs of values from the two slices and returns the results as a new slice.
func Map2[T, T2, R any](mapping func(T, T2) R, values1 []T, values2 []T2) []R {
	min, _ := Min(len(values1), len(values2))
	t3 := make([]R, min)
	iter := func(i int, t T, t2 T2) { t3[i] = mapping(t, t2) }
	Iteri2(iter, values1, values2)
	return t3
}

// Map3 applies mapping to three values from the three slices and returns the results as a new slice.
func Map3[T, T2, T3, R any](action func(T, T2, T3) R, values1 []T, values2 []T2, values3 []T3) []R {
	min, _ := Min(len(values1), len(values2), len(values3))
	output := make([]R, min)
	iter := func(i int, t T, t2 T2, t3 T3) { output[i] = action(t, t2, t3) }
	Iteri3(iter, values1, values2, values3)
	return output
}

// Mapi applies mapping to values with the index of each value and returns the results as a new slice.
func Mapi[T, R any](mapping func(int, T) R, values []T) []R {
	output := make([]R, len(values))
	iter := func(i int, t T) { output[i] = mapping(i, t) }
	Iteri(iter, values)
	return output
}

// Mapi2 applies mapping to pairs of values with the index of each value from the two slices and returns the results as a new slice.
func Mapi2[T, T2, R any](mapping func(int, T, T2) R, values1 []T, values2 []T2) []R {
	min, _ := Min(len(values1), len(values2))
	t3 := make([]R, min)
	iter := func(i int, t T, t2 T2) { t3[i] = mapping(i, t, t2) }
	Iteri2(iter, values1, values2)
	return t3
}

// Mapi3 applies mapping to three values with the index of each value from the three slices and returns the results as a new slice.
func Mapi3[T, T2, T3, R any](mapping func(int, T, T2, T3) R, values1 []T, values2 []T2, values3 []T3) []R {
	min, _ := Min(len(values1), len(values2), len(values3))
	output := make([]R, min)
	iter := func(i int, t T, t2 T2, t3 T3) { output[i] = mapping(i, t, t2, t3) }
	Iteri3(iter, values1, values2, values3)
	return output
}

// Map2D applies mapping to each value in the two dimensional slice and returns the results as a new two dimensional slice.
func Map2D[T, R any](mapping func(T) R, values [][]T) [][]R {
	output := CreateFromStructure2D[T, R](values)
	iter := func(i, j int, t T) { output[i][j] = mapping(t) }
	Iteri2D(iter, values)
	return output
}

// Mapi2D applies mapping to each value in the two dimensional slice with the indexes and returns the results as a new two dimensional slice.
func Mapi2D[T, R any](mapping func(int, int, T) R, values [][]T) [][]R {
	output := CreateFromStructure2D[T, R](values)
	iter := func(i, j int, t T) { output[i][j] = mapping(i, j, t) }
	Iteri2D(iter, values)
	return output
}

// Map3D applies mapping to each value in the three dimensional slice and returns the results as a new three dimensional slice.
func Map3D[T, R any](mapping func(T) R, values [][][]T) [][][]R {
	output := CreateFromStructure3D[T, R](values)
	iter := func(i, j, k int, t T) { output[i][j][k] = mapping(t) }
	Iteri3D(iter, values)
	return output
}

// Mapi3D applies mapping to each value in the three dimensional slice with the indexes and returns the results as a new three dimensional slice.
func Mapi3D[T, R any](mapping func(int, int, int, T) R, values [][][]T) [][][]R {
	output := CreateFromStructure3D[T, R](values)
	iter := func(i, j, k int, t T) { output[i][j][k] = mapping(i, j, k, t) }
	Iteri3D(iter, values)
	return output
}
