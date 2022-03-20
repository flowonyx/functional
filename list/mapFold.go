package list

// MapFold combines map and fold. Builds a new list whose elements are the results of applying the given function to each of the elements of the input list.
// The function is also used to accumulate a final value.
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

// MapFoldBack combines map and foldBack. Builds a new list whose elements are the results of applying the given function to each of the elements of the input list in reverse.
// The function is also used to accumulate a final value.
func MapFoldBack[T any, State any, R any](mapping func(State, T) (R, State), initial State, input []T) ([]R, State) {
	return mapFold(mapping, initial, input, IterRev[T])
}
