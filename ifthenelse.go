package functional

type ifthenelse[T any] struct {
	conds []bool
	thens []func() T
}

// If is essentially like a ternary expression. It allows for an if statement on one line.
// cond is the boolean condition to check against.
// then is the function to execute if cond is true. The given function must return a value.
// It will only be called once there is an Else invoked.
func If[T any](cond bool, then func() T) *ifthenelse[T] {
	return &ifthenelse[T]{conds: []bool{cond}, thens: []func() T{then}}
}

// Elif is the "else if" clause. It is evaluated if the If condition is false.
// cond is the boolean condition to check against.
// then is the function to execute if cond is true.
// The given function must return a value of the same type that the then function in If clause returns.
// It will only be called once there is an Else invoked.
func (ife *ifthenelse[T]) Elif(cond bool, then func() T) *ifthenelse[T] {
	ife.conds = append(ife.conds, cond)
	ife.thens = append(ife.thens, then)
	return ife
}

// Else accepts a function that returns a value when If and Elif clauses fail.
// then is the function to execute if none of the previous conditions are true.
// The given function must return a value of the same type that the then function in If clause returns.
func (ife *ifthenelse[T]) Else(then func() T) T {
	for i := 0; i < len(ife.conds); i++ {
		if ife.conds[i] {
			return ife.thens[i]()
		}
	}
	return then()
}

type ifthenelsev[T any] struct {
	conds []bool
	thens []T
}

// IfV (for "if value") is essentially like a ternary expression. It allows for an if statement on one line.
// cond is the boolean condition to check against.
// then is the value to return if cond is true.
// It will only be returned once there is an Else invoked.
func IfV[T any](cond bool, then T) *ifthenelsev[T] {
	return &ifthenelsev[T]{conds: []bool{cond}, thens: []T{then}}
}

// Elif is the "else if" clause. It is evaluated if the If condition is false.
// cond is the boolean condition to check against.
// then is the value to return if cond is true.
// The given value must be of the same type that the If clause has in its then parameter.
// It will only be called once there is an Else invoked.
func (ife *ifthenelsev[T]) Elif(cond bool, then T) *ifthenelsev[T] {
	ife.conds = append(ife.conds, cond)
	ife.thens = append(ife.thens, then)
	return ife
}

// Else accepts a value to return when If and Elif clauses fail.
// then is the value to return if none of the previous conditions are true.
// The value must be of the same type as the then parameter in the IfV clause.
func (ife *ifthenelsev[T]) Else(then T) T {
	for i := 0; i < len(ife.conds); i++ {
		if ife.conds[i] {
			return ife.thens[i]
		}
	}
	return then
}
