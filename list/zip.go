package list

import . "github.com/flowonyx/functional"

func Zip[T, T2 any](input1 []T, input2 []T2) []Pair[T, T2] {
	return Map2(func(t1 T, t2 T2) Pair[T, T2] {
		return PairOf(t1, t2)
	}, input1, input2)
}

func Zip3[T, T2, T3 any](input1 []T, input2 []T2, input3 []T3) []Triple[T, T2, T3] {
	return Map3(func(t1 T, t2 T2, t3 T3) Triple[T, T2, T3] {
		return TripleOf(t1, t2, t3)
	}, input1, input2, input3)
}

func Unzip[T, T2 any](input []Pair[T, T2]) ([]T, []T2) {
	output1 := make([]T, len(input))
	output2 := make([]T2, len(input))
	for i := range input {
		output1[i], output2[i] = FromPair(input[i])
	}
	return output1, output2
}

func Unzip3[T, T2, T3 any](input []Triple[T, T2, T3]) ([]T, []T2, []T3) {
	output1 := make([]T, len(input))
	output2 := make([]T2, len(input))
	output3 := make([]T3, len(input))
	for i := range input {
		output1[i], output2[i], output3[i] = FromTriple(input[i])
	}
	return output1, output2, output3
}
