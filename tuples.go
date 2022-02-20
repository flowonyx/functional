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

type Triple[T, T2, T3 any] struct {
	First  T
	Second T2
	Third  T3
}

func TripleOf[T, T2, T3 any](a T, b T2, c T3) Triple[T, T2, T3] {
	return Triple[T, T2, T3]{a, b, c}
}

func FromTriple[T, T2, T3 any](t Triple[T, T2, T3]) (a T, b T2, c T3) {
	return t.First, t.Second, t.Third
}
