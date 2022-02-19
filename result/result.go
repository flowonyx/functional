package result

import (
	"errors"

	"github.com/flowonyx/functional"
)

type ok[T any] struct {
	value T
}

type Result[T any] struct {
	value *ok[T]
	err   error
}

func (r Result[T]) IsOK() bool {
	return r.value != nil
}

func (r Result[T]) IsError() bool {
	return r.err != nil
}

func (r Result[T]) Value() T {
	if r.value == nil {
		v := new(T)
		return *v
	}
	return r.value.value
}

func (r Result[_]) Err() error {
	return r.err
}

func OK[T any](v T) Result[T] {
	return Result[T]{
		value: &ok[T]{v},
		err:   nil,
	}
}

func Error[T any](err error) Result[T] {
	return Result[T]{
		value: nil,
		err:   err,
	}
}

func HandleResult[T any](input Result[T], whenOK func(r T) error, whenError func(error) error) error {
	if whenOK == nil {
		return errors.New("whenOK function must be supplied to HandleOption")
	}
	if whenError == nil {
		return errors.New("whenError function must be supplied to HandleOption")
	}
	if input.IsError() {
		return whenError(input.Err())
	}
	return whenOK(input.Value())
}

func Bind[T, R any](binder func(T) Result[R], result Result[T]) Result[R] {
	if result.IsError() {
		return Error[R](result.Err())
	}
	return binder(result.Value())
}

func MapResult[T, R any](mapping functional.Projection[T, R], result Result[T]) Result[R] {
	if result.IsError() {
		return Error[R](result.Err())
	}
	return OK(mapping(result.Value()))
}
