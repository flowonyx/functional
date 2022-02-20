package functional

type Projection[T, R any] func(T) R
type Predicate[T any] func(T) bool
type Predicate2[T, T2 any] func(T, T2) bool
