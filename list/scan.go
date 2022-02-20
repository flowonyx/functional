package list

func Scan[State any, T any](folder func(State, T) State, initialState State, input []T) []State {
	return scan(folder, initialState, input, Iter[T])
}

func scan[State any, T any](folder func(State, T) State, initialState State, input []T, iterFunc func(func(T), []T)) []State {
	output := []State{initialState}
	intermediate := initialState

	iterFunc(func(t T) {
		intermediate = folder(intermediate, t)
		output = append(output, intermediate)
	}, input)

	return output
}

func ScanBack[State any, T any](folder func(State, T) State, initialState State, input []T) []State {
	return scan(folder, initialState, input, IterRev[T])
}
