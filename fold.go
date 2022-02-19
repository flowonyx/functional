package functional

import "github.com/flowonyx/functional/option"

func Fold[State any, T any](folder func(State, T) State, initialState State, input []T) State {
	output := initialState
	Iter(func(t T) {
		output = folder(output, t)
	}, input)
	return output
}

func Fold2[State any, T any, T2 any](folder func(State, T, T2) State, initialState State, input1 []T, input2 []T2) State {
	output := initialState
	Iter2(func(t1 T, t2 T2) {
		output = folder(output, t1, t2)
	}, input1, input2)
	return output
}

func FoldBack[State any, T any](folder func(T, State) State, input []T, initialState State) State {
	output := initialState
	IterRev(func(t T) {
		output = folder(t, output)
	}, input)
	return output
}

func FoldBack2[State any, T any, T2 any](folder func(T, T2, State) State, input1 []T, input2 []T2, initialState State) State {
	output := initialState
	Iter2Rev(func(t1 T, t2 T2) { output = folder(t1, t2, output) }, input1, input2)
	return output
}

func Unfold[T any, State any](generator func(State) option.Option[Pair[T, State]], state State) []T {
	output := []T{}
	s := generator(state)

	for s.IsSome() && len(output) < 1000000 {
		output = append(output, s.Value().First)
		s = generator(s.Value().Second)
	}

	return output
}
