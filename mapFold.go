package functional

func MapFold[T any, State any, R any](mapping func(State, T) (R, State), initial State, input []T) ([]R, State) {
	return mapFold(mapping, initial, input, Iter[T])
}

func mapFold[T any, State any, R any](mapping func(State, T) (R, State), initial State, input []T, rangFunc func(func(T), []T)) ([]R, State) {
	results := []R{}
	s := initial
	r := *(new(R))
	rangFunc(func(t T) {
		r, s = mapping(s, t)
		results = append(results, r)
	}, input)
	return results, s
}

func MapFoldBack[T any, State any, R any](mapping func(State, T) (R, State), initial State, input []T) ([]R, State) {
	return mapFold(mapping, initial, input, IterRev[T])
}
