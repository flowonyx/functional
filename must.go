package functional

// Must takes the output of a function that returns a value and an error,
// panics if the error is not nil, or otherwise returns the value.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// Must2 takes the output of a function that returns two values and an error,
// panics if the error is not nil, or otherwise returns the values.
func Must2[T1, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return t1, t2
}

// Must_0 takes a function that returns an error and returns a functions returns nothing
// and panics if there is an error.
func Must_0[T any](f func(T) error) func(T) {
	return func(t T) {
		err := f(t)
		if err != nil {
			panic(err)
		}
	}
}

// Must_1 takes a function that returns a value and an error and returns a functions returns only the value
// and panics if there is an error.
func Must_1[T, R any](f func(T) (R, error)) func(T) R {
	return func(t T) R {
		return Must(f(t))
	}
}

// Must_2 takes a function that returns two values and an error and returns a functions returns only the two values
// and panics if there is an error.
func Must_2[T, R1, R2 any](f func(T) (R1, R2, error)) func(T) (R1, R2) {
	return func(t T) (R1, R2) {
		return Must2(f(t))
	}
}
