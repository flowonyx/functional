package functional

type Pair[T any, T2 any] struct {
	First  T
	Second T2
}

func PairOf[T, T2 any](a T, b T2) Pair[T, T2] {
	return Pair[T, T2]{First: a, Second: b}
}

func FromPair[T, T2 any](p Pair[T, T2]) (T, T2) {
	return p.First, p.Second
}

func AllPairs[T, T2 any](input1 []T, input2 []T2) []Pair[T, T2] {
	output := Empty[Pair[T, T2]](len(input1) * len(input2))

	Iter(func(t T) {
		Iter(func(t2 T2) {
			output = append(output, PairOf(t, t2))
		}, input2)
	}, input1)

	return output
}

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
