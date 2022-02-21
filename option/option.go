package option

import (
	"errors"
	"fmt"

	"github.com/flowonyx/functional"
)

type optional[T any] interface {
	IsSome() bool
	IsNone() bool
	Value() T
}

type Option[T any] struct {
	value *some[T]
}

func (o Option[T]) String() string {
	if o.IsNone() {
		return "None"
	}
	return fmt.Sprintf("Some(%v)", o.Value())
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

func DefaultValue[T any](value T, input Option[T]) T {
	if input.IsNone() {
		return value
	}
	return input.Value()
}

func DefaultWith[T any](defThunk func() T, input Option[T]) T {
	if input.IsNone() {
		return defThunk()
	}
	return input.Value()
}

func Bind[T, R any](f func(T) Option[R], input Option[T]) Option[R] {
	if input.IsNone() {
		return None[R]()
	}
	return f(input.Value())
}

func Contains[T comparable](value T, o Option[T]) bool {
	return o.IsSome() && o.Value() == value
}

func Count[T any](o Option[T]) int {
	if o.IsNone() {
		return 0
	}
	return 1
}

func Exists[T any](predicate functional.Predicate[T], o Option[T]) bool {
	r := Map(predicate, o)
	return r.IsSome() && r.Value()
}

func Filter[T any](predicate functional.Predicate[T], o Option[T]) Option[T] {
	if Exists(predicate, o) {
		return o
	}
	return None[T]()
}

func Flatten[T any](oo Option[Option[T]]) Option[T] {
	if oo.IsNone() {
		return None[T]()
	}
	return oo.Value()
}

func Fold[T, State any](folder func(State, T) State, s State, o Option[T]) State {
	r := Map(functional.Curry2To1(folder, s), o)
	return r.Value()
}

func FoldBack[T, State any](folder func(T, State) State, o Option[T], s State) State {
	if o.IsNone() {
		return s
	}
	return folder(o.Value(), s)
}

func ForAll[T any](predicate functional.Predicate[T], o Option[T]) bool {
	if o.IsNone() {
		return true
	}
	return predicate(o.Value())
}

func Get[T any](o Option[T]) T {
	if o.IsNone() {
		panic("cannot get value of None")
	}
	return o.Value()
}

func IsNone[T any](o Option[T]) bool {
	return o.IsNone()
}

func IsSome[T any](o Option[T]) bool {
	return o.IsSome()
}

func Iter[T any](action func(T), o Option[T]) {
	if o.IsNone() {
		return
	}
	action(o.Value())
}

func Map[T, R any](f func(T) R, o Option[T]) Option[R] {
	if o.IsNone() {
		return None[R]()
	}
	return Some(f(o.Value()))
}

func Map2[T1, T2, R any](f func(T1, T2) R, o1 Option[T1], o2 Option[T2]) Option[R] {
	if o1.IsNone() || o2.IsNone() {
		return None[R]()
	}
	return Some(f(o1.Value(), o2.Value()))
}

func Map3[T1, T2, T3, R any](f func(T1, T2, T3) R, o1 Option[T1], o2 Option[T2], o3 Option[T3]) Option[R] {
	if o1.IsNone() || o2.IsNone() || o3.IsNone() {
		return None[R]()
	}
	return Some(f(o1.Value(), o2.Value(), o3.Value()))
}

func OfNullable[T any](value *T) Option[T] {
	if value == nil {
		return None[T]()
	}
	return Some(*value)
}

func OrElse[T any](ifNone Option[T], o Option[T]) Option[T] {
	if o.IsNone() {
		return ifNone
	}
	return o
}

func OrElseWith[T any](ifNoneThunk func() Option[T], o Option[T]) Option[T] {
	if o.IsNone() {
		return ifNoneThunk()
	}
	return o
}

func ToSlice[T any](o Option[T]) []T {
	if o.IsNone() {
		return []T{}
	}
	return []T{o.Value()}
}

func ToNullable[T any](o Option[T]) *T {
	if o.IsNone() {
		return nil
	}
	v := o.Value()
	return &v
}
