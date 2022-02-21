package functional

type ifthenelse[T any] struct {
	conds []bool
	thens []func() T
}

func If[T any](cond bool, then func() T) *ifthenelse[T] {
	return &ifthenelse[T]{conds: []bool{cond}, thens: []func() T{then}}
}

func (ife *ifthenelse[T]) Elif(cond bool, then func() T) *ifthenelse[T] {
	ife.conds = append(ife.conds, cond)
	ife.thens = append(ife.thens, then)
	return ife
}

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

func IfV[T any](cond bool, then T) *ifthenelsev[T] {
	return &ifthenelsev[T]{conds: []bool{cond}, thens: []T{then}}
}

func (ife *ifthenelsev[T]) Elif(cond bool, then T) *ifthenelsev[T] {
	ife.conds = append(ife.conds, cond)
	ife.thens = append(ife.thens, then)
	return ife
}

func (ife *ifthenelsev[T]) Else(then T) T {
	for i := 0; i < len(ife.conds); i++ {
		if ife.conds[i] {
			return ife.thens[i]
		}
	}
	return then
}
