package option

import (
	"fmt"
	"strconv"
)

// Option is a type that allows optional values. It is similar to a Sum type
// and can either be Some(value) or None.
type Option[T any] struct {
	value *some[T]
}

func (o Option[T]) String() string {
	if o.IsNone() {
		return "None"
	}
	val := printv(o.Value())
	return fmt.Sprintf("Some(%v)", val)
}

type some[T any] struct {
	value T
}

// Some creates an option with a value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: &some[T]{value: value}}
}

// None creates an option with no value.
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

// IsSome tests if the option contains a value.
func (o Option[T]) IsSome() bool {
	return o.value != nil
}

// IsNone tests whether the option does not contain a value.
func (o Option[T]) IsNone() bool {
	return o.value == nil
}

// Value returns the value in the option.
// If the option is None, it returns the zero value for the type.
func (o Option[T]) Value() T {
	if o.IsNone() {
		r := new(T)
		return *r
	}
	return o.value.value
}

// Bind applies f to input if input.IsSome() and otherwise returns None.
func Bind[T, R any](f func(T) Option[R], input Option[T]) Option[R] {
	if input.IsNone() {
		return None[R]()
	}
	return f(input.Value())
}

// Filter retuns o if the value in o matches the predicate.
// Otherwise, it returns None.
func Filter[T any](predicate func(T) bool, o Option[T]) Option[T] {
	if Exists(predicate, o) {
		return o
	}
	return None[T]()
}

// Flatten takes a nested option and returns the inner option.
func Flatten[T any](oo Option[Option[T]]) Option[T] {
	if oo.IsNone() {
		return None[T]()
	}
	return oo.Value()
}

// Map applies f to the value of o and returns the result as an Option.
// If o is None, it returns None.
func Map[T, R any](f func(T) R, o Option[T]) Option[R] {
	if o.IsNone() {
		return None[R]()
	}
	return Some(f(o.Value()))
}

// Map2 applies f to the values in both o1 and o2 as the first and second parameters and returns the result as an Option.
// If either option is None, it returns None.
func Map2[T1, T2, R any](f func(T1, T2) R, o1 Option[T1], o2 Option[T2]) Option[R] {
	if o1.IsNone() || o2.IsNone() {
		return None[R]()
	}
	return Some(f(o1.Value(), o2.Value()))
}

// Map3 applies f to the values in o1, o2, and o3 as the first, second, and third parameters and returns the result as an Option.
// If any of the options are None, it returns None.
func Map3[T1, T2, T3, R any](f func(T1, T2, T3) R, o1 Option[T1], o2 Option[T2], o3 Option[T3]) Option[R] {
	if o1.IsNone() || o2.IsNone() || o3.IsNone() {
		return None[R]()
	}
	return Some(f(o1.Value(), o2.Value(), o3.Value()))
}

// OfNullable returns None if value is nil.
// Otherwise it returns Some of the value (after dereferencing the pointer).
func OfNullable[T any](value *T) Option[T] {
	if value == nil {
		return None[T]()
	}
	return Some(*value)
}

// Lift converts the function f that returns a value and an error
// to a function that returns an Option.
func Lift[T any](f func() (T, error)) func() Option[T] {
	return func() Option[T] {
		r, err := f()
		if err != nil {
			return None[T]()
		}
		return Some(r)
	}
}

// Lift1 converts the function f that accepts a single input and returns a value and an error
// to a function that accepts a single input and returns an Option.
func Lift1[TInput, T any](f func(TInput) (T, error)) func(TInput) Option[T] {
	return func(input TInput) Option[T] {
		r, err := f(input)
		if err != nil {
			return None[T]()
		}
		return Some(r)
	}
}

// Lift2 converts the function f that accepts two inputs and returns a value and an error
// to a function that accepts two inputs and returns an Option.
func Lift2[TInput1, TInput2, T any](f func(TInput1, TInput2) (T, error)) func(TInput1, TInput2) Option[T] {
	return func(input1 TInput1, input2 TInput2) Option[T] {
		r, err := f(input1, input2)
		if err != nil {
			return None[T]()
		}
		return Some(r)
	}
}

// helper function for quoting the internal values
func printv(v any) string {
	switch o := v.(type) {
	case string:
		return strconv.Quote(o)
	case rune:
		return strconv.QuoteRune(o)
	default:
		return fmt.Sprint(o)
	}
}
