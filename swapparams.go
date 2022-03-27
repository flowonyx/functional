package functional

// SwapParams0 adapts a function to take the second parameter as the first and the first parameter as the second.
// The supplied function must have no return value.
func SwapParams0[T1, T2 any](f func(T1, T2)) func(T2, T1) {
	return func(t1 T2, t2 T1) {
		f(t2, t1)
	}
}

// SwapParams1 adapts a function to take the second parameter as the first and the first parameter as the second.
// The supplied function must have one return value.
func SwapParams1[T1, T2, R any](f func(T1, T2) R) func(T2, T1) R {
	return func(t1 T2, t2 T1) R {
		return f(t2, t1)
	}
}

// SwapParams2 adapts a function to take the second parameter as the first and the first parameter as the second.
// The supplied function must have one return value.
func SwapParams2[T1, T2, R1, R2 any](f func(T1, T2) (R1, R2)) func(T2, T1) (R1, R2) {
	return func(t1 T2, t2 T1) (R1, R2) {
		return f(t2, t1)
	}
}
