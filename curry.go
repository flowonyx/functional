package functional

// Curry accepts a function that receives a parameter and the value of the parameter
// and returns a function that accepts no parameters but returns the result of
// applying the function to the given parameter.
func Curry[T, R any](f func(T) R, input T) func() R {
	return func() R {
		return f(input)
	}
}

// Curry_0 accepts a function that accepts one parameter and the value of the parameter
// and returns a function that accepts no parameters. The function must have no return value.
func Curry_0[T any](f func(T), input T) func() {
	return func() {
		f(input)
	}
}

// Curry_2 accepts a function that accepts one parameter and the value of the parameter
// and returns a function that accepts no parameters. The function must have 2 return values.
func Curry_2[T, R1, R2 any](f func(T) (R1, R2), input T) func() (R1, R2) {
	return func() (R1, R2) {
		return f(input)
	}
}

// Curry2 accepts a function that receives two parameters and the values of those parameters
// and returns a function that accepts no parameters but returns the result of
// applying the function to the given parameters.
func Curry2[T1, T2, R any](f func(T1, T2) R, input1 T1, input2 T2) func() R {
	return func() R {
		return f(input1, input2)
	}
}

// Curry2_0 accepts a function that receives two parameters and the values of those parameters
// and returns a function that accepts no parameters. The function must have no return value.
func Curry2_0[T1, T2 any](f func(T1, T2), input1 T1, input2 T2) func() {
	return func() {
		f(input1, input2)
	}
}

// Curry2_2 accepts a function that receives two parameters and the values of those parameters
// and returns a function that accepts no parameters but returns the result of
// applying the function to the given parameters. The function must have 2 return values.
func Curry2_2[T1, T2, R1, R2 any](f func(T1, T2) (R1, R2), input1 T1, input2 T2) func() (R1, R2) {
	return func() (R1, R2) {
		return f(input1, input2)
	}
}

// Curry3 accepts a function that receives three parameters and the values of those parameters
// and returns a function that accepts no parameters but returns the result of
// applying the function to the given parameters.
func Curry3[T1, T2, T3, R any](f func(T1, T2, T3) R, input1 T1, input2 T2, input3 T3) func() R {
	return func() R {
		return f(input1, input2, input3)
	}
}

// Curry3_0 accepts a function that receives three parameters and the values of those parameters
// and returns a function that accepts no parameters. The function must have no return value.
func Curry3_0[T1, T2, T3 any](f func(T1, T2, T3), input1 T1, input2 T2, input3 T3) func() {
	return func() {
		f(input1, input2, input3)
	}
}

// Curry3_2 accepts a function that receives three parameters and the values of those parameters
// and returns a function that accepts no parameters but returns the result of
// applying the function to the given parameters. The function must have 2 return values.
func Curry3_2[T1, T2, T3, R1, R2 any](f func(T1, T2, T3) (R1, R2), input1 T1, input2 T2, input3 T3) func() (R1, R2) {
	return func() (R1, R2) {
		return f(input1, input2, input3)
	}
}

// Curry2To1 accepts a function that receives two parameters and the value of the first parameter
// and returns a function that accepts one parameter which is the second parameter in the original function
// and returns the result of applying the function with input1 as the first parameter and
// whatever is given to the curried function as the second parameter.
func Curry2To1[T1, T2, R any](f func(T1, T2) R, input1 T1) func(T2) R {
	return func(input2 T2) R {
		return f(input1, input2)
	}
}

// Curry2To1_0 accepts a function that receives two parameters and the value of the first parameter
// and returns a function that accepts one parameter which is the second parameter in the original function.
// The curried function returns the result of applying the original function to input1 as the first parameter and
// whatever is given to the last returned function as the second parameter.
// The function must have no return value.
func Curry2To1_0[T1, T2 any](f func(T1, T2), input1 T1) func(T2) {
	return func(input2 T2) {
		f(input1, input2)
	}
}

// Curry2To1_2 accepts a function that receives two parameters and the value of the first parameter
// and returns a function that accepts one parameter which is the second parameter in the original function
// and returns the result of applying the function with input1 as the first parameter and
// whatever is given to the curried function as the second parameter. The function must 2 return values.
func Curry2To1_2[T1, T2, R1, R2 any](f func(T1, T2) (R1, R2), input1 T1) func(T2) (R1, R2) {
	return func(input2 T2) (R1, R2) {
		return f(input1, input2)
	}
}

// Curry2To1F accepts a function that receives two parameters
// and returns a function that accepts one parameter which is the first parameter in the original function
// and returns another function that accepts one parameter, which is the second parameter in the original function.
// The last function returns the result of applying the original function with the input to the first returned function as the first parameter and
// whatever is given to the last returned function as the second parameter.
func Curry2To1F[T1, T2, R any](f func(T1, T2) R) func(T1) func(T2) R {
	return func(input1 T1) func(T2) R {
		return func(input2 T2) R {
			return f(input1, input2)
		}
	}
}

// Curry3To2 accepts a function that receives three parameters
// and returns a function that accepts two parameters which are the last two parameters in the original function.
// The curried function returns the result of applying the original function to input1 as the first parameter and
// whatever is given to the last returned function as the second and third parameters.
func Curry3To2[T1, T2, T3, R any](f func(T1, T2, T3) R, input1 T1) func(T2, T3) R {
	return func(input2 T2, input3 T3) R {
		return f(input1, input2, input3)
	}
}

// Curry3To2_0 accepts a function that receives three parameters
// and returns a function that accepts two parameters which are the last two parameters in the original function.
// The curried function applies the original function to input1 as the first parameter and
// whatever is given to the last returned function as the second and third parameters.
// The function must have no return value.
func Curry3To2_0[T1, T2, T3 any](f func(T1, T2, T3), input1 T1) func(T2, T3) {
	return func(input2 T2, input3 T3) {
		f(input1, input2, input3)
	}
}

// Curry3To2_2 accepts a function that receives three parameters
// and returns a function that accepts two parameters which are the last two parameters in the original function.
// The curried function returns the result of applying the original function to input1 as the first parameter and
// whatever is given to the last returned function as the second and third parameters. The function must have 2 return values.
func Curry3To2_2[T1, T2, T3, R1, R2 any](f func(T1, T2, T3) (R1, R2), input1 T1) func(T2, T3) (R1, R2) {
	return func(input2 T2, input3 T3) (R1, R2) {
		return f(input1, input2, input3)
	}
}

// Curry3To1 accepts a function that receives three parameters
// and returns a function that accepts one parameter which is the last parameter in the original function.
// The curried function returns the result of applying the original function to input1 and input2 as the first two parameters and
// whatever is given to the last returned function as the third parameter.
func Curry3To1[T1, T2, T3, R any](f func(T1, T2, T3) R, input1 T1, input2 T2) func(T3) R {
	return func(input3 T3) R {
		return f(input1, input2, input3)
	}
}

// Curry3To1_0 accepts a function that receives three parameters
// and returns a function that accepts one parameter which is the last parameter in the original function.
// The curried function applies the original function to input1 and input2 as the first two parameters and
// whatever is given to the last returned function as the third parameter.
// The function must have no return value.
func Curry3To1_0[T1, T2, T3 any](f func(T1, T2, T3), input1 T1, input2 T2) func(T3) {
	return func(input3 T3) {
		f(input1, input2, input3)
	}
}

// Curry3To1_2 accepts a function that receives three parameters
// and returns a function that accepts one parameter which is the last parameter in the original function.
// The curried function returns the result of applying the original function to input1 and input2 as the first two parameters and
// whatever is given to the last returned function as the third parameter. The function must have 2 return values.
func Curry3To1_2[T1, T2, T3, R1, R2 any](f func(T1, T2, T3) (R1, R2), input1 T1, input2 T2) func(T3) (R1, R2) {
	return func(input3 T3) (R1, R2) {
		return f(input1, input2, input3)
	}
}

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
