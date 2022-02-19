package functional

type Projection[T, R any] func(T) R
type Predicate[T any] func(T) bool
type Predicate2[T, T2 any] func(T, T2) bool

func Map[T, R any](f Projection[T, R], input []T) []R {
	output := make([]R, len(input))
	iter := func(i int, t T) { output[i] = f(t) }
	Iteri(iter, input)
	return output
}

func Map2[T, T2, R any](f func(T, T2) R, input1 []T, input2 []T2) []R {
	min := Min(len(input1), len(input2))
	t3 := make([]R, min)
	iter := func(i int, t T, t2 T2) { t3[i] = f(t, t2) }
	Iteri2(iter, input1, input2)
	return t3
}

func Map3[T, T2, T3, R any](action func(T, T2, T3) R, input1 []T, input2 []T2, input3 []T3) []R {
	min := Min(len(input1), len(input2), len(input3))
	output := make([]R, min)
	iter := func(i int, t T, t2 T2, t3 T3) { output[i] = action(t, t2, t3) }
	Iteri3(iter, input1, input2, input3)
	return output
}

func Mapi[T, R any](mapping func(int, T) R, input []T) []R {
	output := make([]R, len(input))
	iter := func(i int, t T) { output[i] = mapping(i, t) }
	Iteri(iter, input)
	return output
}

func Mapi2[T, T2, R any](f func(int, T, T2) R, input1 []T, input2 []T2) []R {
	min := Min(len(input1), len(input2))
	t3 := make([]R, min)
	iter := func(i int, t T, t2 T2) { t3[i] = f(i, t, t2) }
	Iteri2(iter, input1, input2)
	return t3
}

func Mapi3[T, T2, T3, R any](action func(int, T, T2, T3) R, input1 []T, input2 []T2, input3 []T3) []R {
	min := Min(len(input1), len(input2), len(input3))
	output := make([]R, min)
	iter := func(i int, t T, t2 T2, t3 T3) { output[i] = action(i, t, t2, t3) }
	Iteri3(iter, input1, input2, input3)
	return output
}

func Map2D[T, R any](mapping Projection[T, R], input [][]T) [][]R {
	output := CreateFromStructure2D[T, R](input)
	iter := func(i, j int, t T) { output[i][j] = mapping(t) }
	Iteri2D(iter, input)
	return output
}

func Mapi2D[T, R any](mapping func(int, int, T) R, input [][]T) [][]R {
	output := CreateFromStructure2D[T, R](input)
	iter := func(i, j int, t T) { output[i][j] = mapping(i, j, t) }
	Iteri2D(iter, input)
	return output
}

func Map3D[T, R any](mapping Projection[T, R], input [][][]T) [][][]R {
	output := CreateFromStructure3D[T, R](input)
	iter := func(i, j, k int, t T) { output[i][j][k] = mapping(t) }
	Iteri3D(iter, input)
	return output
}

func Mapi3D[T, R any](mapping func(int, int, int, T) R, input [][][]T) [][][]R {
	output := CreateFromStructure3D[T, R](input)
	iter := func(i, j, k int, t T) { output[i][j][k] = mapping(i, j, k, t) }
	Iteri3D(iter, input)
	return output
}
