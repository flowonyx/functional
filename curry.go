package functional

func Curry[T1, T2 any](f func(T1) T2, input T1) func() T2 {
	return func() T2 {
		return f(input)
	}
}

func Curry2[T1, T2, T3 any](f func(T1, T2) T3, input1 T1, input2 T2) func() T3 {
	return func() T3 {
		return f(input1, input2)
	}
}

func Curry3[T1, T2, T3, T4 any](f func(T1, T2, T3) T4, input1 T1, input2 T2, input3 T3) func() T4 {
	return func() T4 {
		return f(input1, input2, input3)
	}
}

func Curry2To1[T1, T2, T3 any](f func(T1, T2) T3, input1 T1) func(T2) T3 {
	return func(input2 T2) T3 {
		return f(input1, input2)
	}
}

func Curry2To1F[T1, T2, T3 any](f func(T1, T2) T3) func(T1) func(T2) T3 {
	return func(input1 T1) func(T2) T3 {
		return func(input2 T2) T3 {
			return f(input1, input2)
		}
	}
}

func Curry3To2[T1, T2, T3, T4 any](f func(T1, T2, T3) T4, input1 T1) func(T2, T3) T4 {
	return func(input2 T2, input3 T3) T4 {
		return f(input1, input2, input3)
	}
}

func Curry3To1[T1, T2, T3, T4 any](f func(T1, T2, T3) T4, input1 T1, input2 T2) func(T3) T4 {
	return func(input3 T3) T4 {
		return f(input1, input2, input3)
	}
}
