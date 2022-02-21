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

func Map[T, R any](mapping functional.Projection[T, R], result Result[T]) Result[R] {
	if result.IsError() {
		return Error[R](result.Err())
	}
	return OK(mapping(result.Value()))
}

func MapError[T any](mapping functional.Projection[error, error], result Result[T]) Result[T] {
	if result.IsError() {
		return Error[T](mapping(result.Err()))
	}
	return result
}

func DefaultValue[T any](value T, input Result[T]) T {
	if input.IsError() {
		return value
	}
	return input.Value()
}

func DefaultWith[T any](defThunk func() T, input Result[T]) T {
	if input.IsError() {
		return defThunk()
	}
	return input.Value()
}

func Contains[T comparable](value T, o Result[T]) bool {
	return o.IsOK() && o.Value() == value
}

func Count[T any](o Result[T]) int {
	if o.IsError() {
		return 0
	}
	return 1
}

func Exists[T any](predicate functional.Predicate[T], o Result[T]) bool {
	if o.IsError() {
		return false
	}
	r := predicate(o.Value())
	return r
}

func Filter[T any](predicate functional.Predicate[T], o Result[T]) Result[T] {
	if Exists(predicate, o) {
		return o
	}
	if o.IsError() {
		return o
	}
	return Error[T](errors.New("does not exist"))
}

func Flatten[T any](oo Result[Result[T]]) Result[T] {
	if oo.IsError() {
		return Error[T](oo.Err())
	}
	return oo.Value()
}

func Fold[T, State any](folder func(State, T) State, s State, o Result[T]) State {
	r := Map(functional.Curry2To1(folder, s), o)
	return r.Value()
}

func FoldBack[T, State any](folder func(T, State) State, o Result[T], s State) State {
	if o.IsError() {
		return s
	}
	return folder(o.Value(), s)
}

func ForAll[T any](predicate functional.Predicate[T], o Result[T]) bool {
	if o.IsError() {
		return true
	}
	return predicate(o.Value())
}

func Get[T any](o Result[T]) T {
	if o.IsError() {
		panic("cannot get value of None")
	}
	return o.Value()
}

func IsNone[T any](o Result[T]) bool {
	return o.IsError()
}

func IsSome[T any](o Result[T]) bool {
	return o.IsOK()
}

func Iter[T any](action func(T), o Result[T]) {
	if o.IsError() {
		return
	}
	action(o.Value())
}

func Map2[T1, T2, R any](f func(T1, T2) R, o1 Result[T1], o2 Result[T2]) Result[R] {
	if o1.IsError() {
		return Error[R](o1.Err())
	}
	if o2.IsError() {
		return Error[R](o2.Err())
	}
	return OK(f(o1.Value(), o2.Value()))
}

func Map3[T1, T2, T3, R any](f func(T1, T2, T3) R, o1 Result[T1], o2 Result[T2], o3 Result[T3]) Result[R] {
	if o1.IsError() {
		return Error[R](o1.Err())
	}
	if o2.IsError() {
		return Error[R](o2.Err())
	}
	if o3.IsError() {
		return Error[R](o3.Err())
	}
	return OK(f(o1.Value(), o2.Value(), o3.Value()))
}

func OfNullable[T any](value *T) Result[T] {
	if value == nil {
		return Error[T](errors.New("nil"))
	}
	return OK(*value)
}

func OrElse[T any](ifNone Result[T], o Result[T]) Result[T] {
	if o.IsError() {
		return ifNone
	}
	return o
}

func OrElseWith[T any](ifNoneThunk func() Result[T], o Result[T]) Result[T] {
	if o.IsError() {
		return ifNoneThunk()
	}
	return o
}

func ToSlice[T any](o Result[T]) []T {
	if o.IsError() {
		return []T{}
	}
	return []T{o.Value()}
}

func ToNullable[T any](o Result[T]) *T {
	if o.IsError() {
		return nil
	}
	v := o.Value()
	return &v
}
