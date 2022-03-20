package list

import . "github.com/flowonyx/functional"

// Zip puts the two slices of values into one slice of Pairs.
// It will only returns as many items as the smallest length of the two slices.
func Zip[T, T2 any](values1 []T, values2 []T2) []Pair[T, T2] {
	return Map2(func(t1 T, t2 T2) Pair[T, T2] {
		return PairOf(t1, t2)
	}, values1, values2)
}

// Zip3 puts the three slices of values into one slice of Triples.
// It will only returns as many items as the smallest length of the slices.
func Zip3[T, T2, T3 any](values1 []T, values2 []T2, values3 []T3) []Triple[T, T2, T3] {
	return Map3(func(t1 T, t2 T2, t3 T3) Triple[T, T2, T3] {
		return TripleOf(t1, t2, t3)
	}, values1, values2, values3)
}

// Unzip takes a slice of Pairs and returns a slice of all values in the first position
// and another slice with all values in the second position.
func Unzip[T, T2 any](values []Pair[T, T2]) ([]T, []T2) {
	output1 := make([]T, len(values))
	output2 := make([]T2, len(values))
	for i := range values {
		output1[i], output2[i] = FromPair(values[i])
	}
	return output1, output2
}

// Unzip3 takes a slice of Triples and returns a slice of all values in the first position,
// another slice with all values in the second position, and another slice with all values in the third position.
func Unzip3[T, T2, T3 any](values []Triple[T, T2, T3]) ([]T, []T2, []T3) {
	output1 := make([]T, len(values))
	output2 := make([]T2, len(values))
	output3 := make([]T3, len(values))
	for i := range values {
		output1[i], output2[i], output3[i] = FromTriple(values[i])
	}
	return output1, output2, output3
}
