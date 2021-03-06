package list

import . "github.com/flowonyx/functional"

// AllPairs returns a slice of all pairs of items from two slices.
func AllPairs[T, T2 any](input1 []T, input2 []T2) []Pair[T, T2] {
	output := Empty[Pair[T, T2]](len(input1) * len(input2))

	Iter(func(t T) {
		Iter(func(t2 T2) {
			output = append(output, PairOf(t, t2))
		}, input2)
	}, input1)

	return output
}

// Pairwise returns a slice of each item in the input slice and the one before it, except for the first item
// which is only returned as the one before the second element.
func Pairwise[T any](input []T) []Pair[T, T] {
	if len(input) == 0 {
		return []Pair[T, T]{}
	}
	output := make([]Pair[T, T], len(input)-1)
	Iteri(func(i int, t T) {
		if i == 0 {
			return
		}
		output[i-1] = PairOf(input[i-1], t)
	}, input)

	return output
}
