package list

// Scan applies a function to each element of the collection, threading an accumulator argument through the computation.
// Take the second argument, and apply the function to it and the first element of the list.
// Then feed this result into the function along with the second element and so on. Returns the list of intermediate results and the final result.
func Scan[State any, T any](folder func(State, T) State, initialState State, values []T) []State {
	return scan(folder, initialState, values, Iter[T])
}

func scan[State any, T any](folder func(State, T) State, initialState State, values []T, iterFunc func(func(T), []T)) []State {
	output := []State{initialState}
	intermediate := initialState

	iterFunc(func(t T) {
		intermediate = folder(intermediate, t)
		output = append(output, intermediate)
	}, values)

	return output
}

// ScanBack applies a function to each element of the collection in reverse, threading an accumulator argument through the computation.
// Take the second argument, and apply the function to it and the first element of the list.
// Then feed this result into the function along with the second element and so on. Returns the list of intermediate results and the final result.
func ScanBack[State any, T any](folder func(State, T) State, initialState State, values []T) []State {
	return scan(folder, initialState, values, IterRev[T])
}
