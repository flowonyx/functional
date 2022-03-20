package list

import (
	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/option"
)

// Fold applies folder to each value in order starting with initialState.
func Fold[State any, T any](folder func(State, T) State, initialState State, values []T) State {
	output := initialState
	Iter(func(t T) {
		output = folder(output, t)
	}, values)
	return output
}

// Fold2 applies folder to each pair of values from values1 and values2 in order starting with initialState.
func Fold2[State any, T any, T2 any](folder func(State, T, T2) State, initialState State, values1 []T, values2 []T2) State {
	output := initialState
	Iter2(func(t1 T, t2 T2) {
		output = folder(output, t1, t2)
	}, values1, values2)
	return output
}

// FoldBack applies folder to each value in reverse order starting with initialState.
func FoldBack[State any, T any](folder func(T, State) State, values []T, initialState State) State {
	output := initialState
	IterRev(func(t T) {
		output = folder(t, output)
	}, values)
	return output
}

// FoldBack2 applies folder to each pair of values from values1 and values2 in reverse order starting with initialState.
func FoldBack2[State any, T any, T2 any](folder func(T, T2, State) State, values1 []T, values2 []T2, initialState State) State {
	output := initialState
	Iter2Rev(func(t1 T, t2 T2) { output = folder(t1, t2, output) }, values1, values2)
	return output
}

// Unfold applies generator to state until generator returns None (or it is run 1000000 times).
// It returns all the generated values.
func Unfold[T any, State any](generator func(State) option.Option[Pair[T, State]], state State) []T {
	output := []T{}
	s := generator(state)

	for s.IsSome() && len(output) <= 1000000 {
		output = append(output, s.Value().First)
		s = generator(s.Value().Second)
	}

	return output
}
