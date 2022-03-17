package option

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/flowonyx/functional"
)

type Optional[T any] interface {
	IsSome() bool
	IsNone() bool
	Value() T
}

type OptionalCheckOnly interface {
	IsSome() bool
	IsNone() bool
}

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

// HandleOption accepts a input that is an Optional value and two functions that returns errors.
// whenSome is the function to use when input.IsSome() is true.
// whenNone is the function to use when input.IsNone() is true.
// The error returned by the function that is used will be returned from HandleOption.
func HandleOption[T any, TOptional Optional[T], F func(T) error, FN func() error](input TOptional, whenSome F, whenNone FN) error {
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

// HandleOptionIgnoreNone accepts a input that is an Optional value and one function that returns errors.
// whenSome is the function to use when input.IsSome() is true.
// The error returned by whenSome will be returned from HandleOption or nil will be returned when input.IsNone().
func HandleOptionIgnoreNone[T any, TOptional Optional[T], F func(T) error](input TOptional, whenSome F) error {
	return HandleOption(input, whenSome, func() error { return nil })
}

// DefaultValue tests whether input.IsNone() and returns value if true.
// If input.IsSome(), input.Value() is returned.
func DefaultValue[T any, TOptional Optional[T]](value T, input TOptional) T {
	if input.IsNone() {
		return value
	}
	return input.Value()
}

// DefaultValue tests whether input.IsNone() and returns the result of defThunk if true.
// If input.IsSome(), input.Value() is returned.
func DefaultWith[T any, TOptional Optional[T]](defThunk func() T, input TOptional) T {
	if input.IsNone() {
		return defThunk()
	}
	return input.Value()
}

// Bind applies f to input if input.IsSome() and otherwise returns None.
func Bind[T, R any](f func(T) Option[R], input Option[T]) Option[R] {
	if input.IsNone() {
		return None[R]()
	}
	return f(input.Value())
}

// Contains tests whether the value in the Optional value o is equal to value.
func Contains[T comparable, TOptional Optional[T]](value T, o TOptional) bool {
	return o.IsSome() && o.Value() == value
}

// Count returns 0 if o.IsNone() and 1 if o.IsSome().
func Count(o OptionalCheckOnly) int {
	if o.IsNone() {
		return 0
	}
	return 1
}

// Exists tests is the value in the Optional o matches the predicate.
// If o.IsNone(), it will return false.
func Exists[T any, TOptional Optional[T]](predicate func(T) bool, o TOptional) bool {
	if o.IsNone() {
		return false
	}
	return predicate(o.Value())
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

// Fold applies
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

func ForAll[T any](predicate func(T) bool, o Option[T]) bool {
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

func Lift[T any](f func() (T, error)) func() Option[T] {
	return func() Option[T] {
		r, err := f()
		if err != nil {
			return None[T]()
		}
		return Some(r)
	}
}

func Lift1[TInput, T any](f func(TInput) (T, error)) func(TInput) Option[T] {
	return func(input TInput) Option[T] {
		r, err := f(input)
		if err != nil {
			return None[T]()
		}
		return Some(r)
	}
}

func Lift2[TInput1, TInput2, T any](f func(TInput1, TInput2) (T, error)) func(TInput1, TInput2) Option[T] {
	return func(input1 TInput1, input2 TInput2) Option[T] {
		r, err := f(input1, input2)
		if err != nil {
			return None[T]()
		}
		return Some(r)
	}
}

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
