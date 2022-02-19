package option

import "errors"

type optional[T any] interface {
	IsSome() bool
	IsNone() bool
	Value() T
}

type Option[T any] struct {
	value *some[T]
}

type some[T any] struct {
	value T
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &some[T]{value: value}}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

func (o Option[T]) IsSome() bool {
	return o.value != nil
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) Value() T {
	if o.IsNone() {
		r := new(T)
		return *r
	}
	return o.value.value
}

func HandleOption[T any, F func(T) error, FN func() error](input Option[T], whenSome F, whenNone FN) error {
	if whenSome == nil {
		return errors.New("whenSome function must be supplied to HandleOption")
	}
	if whenNone == nil {
		return errors.New("whenNone function must be supplied to HandleOption")
	}
	if input.IsNone() {
		return whenNone()
	}
	return whenSome(input.Value())
}

func HandleOptionIgnoreNone[T any, F func(T) error](input Option[T], whenSome F) error {
	return HandleOption(input, whenSome, func() error { return nil })
}

func ValueOrDefault[T any](input Option[T], value T) T {
	if input.IsNone() {
		return value
	}
	return input.Value()
}
