// Package result provides a helper type for error handling without returning multiple values.
// It also provides methodes for dealing with the Result type.
package result

import (
	"fmt"

	"github.com/flowonyx/functional/option"
)

type container[T any] struct {
	value T
}

// Result is a helper type for error handling without returning multiple values.
// Any Result will either contain a value or an error.
type Result[S, F any] struct {
	value   *container[S]
	failure *container[F]
}

func (r *Result[_, _]) String() string {
	if r.IsFailure() {
		var v any = r.FailureValue()
		switch f := v.(type) {
		case error:
			return f.Error()
		case string:
			return f
		case fmt.Stringer:
			return f.String()
		}
		return fmt.Sprint(v)
	}

	var v any = r.SuccessValue()
	switch f := v.(type) {
	case error:
		return f.Error()
	case fmt.Stringer:
		return f.String()
	}
	return fmt.Sprint(v)
}

// IsSuccess tests if this Result has a value.
func (r Result[_, _]) IsSuccess() bool {
	return r.value != nil
}

// IsFailure tests if this Result has a failure.
func (r Result[_, _]) IsFailure() bool {
	return r.failure != nil
}

// IsSome is an alias for IsSuccess to satisfy the option.Optional interface.
func (r Result[_, _]) IsSome() bool {
	return r.value != nil
}

// IsNone is an alias for IsFailure to satisfy the option.Optional interface.
func (r Result[_, _]) IsNone() bool {
	return r.failure != nil
}

// SuccessValue returns the value if this Result is Success. If the result is an error, Value returns the zero value of the value type.
func (r Result[S, _]) SuccessValue() S {
	if r.value == nil {
		v := new(S)
		return *v
	}
	return r.value.value
}

// FailureValue returns the error if this Result is Error. Otherwise, it returns nil.
func (r Result[_, F]) FailureValue() F {
	if r.failure == nil {
		v := new(F)
		return *v
	}
	return r.failure.value
}

// Value is an alias for SuccessValue to satisfy the option.Optional interface.
func (r Result[S, _]) Value() S {
	return r.SuccessValue()
}

// Success creates a Result with a value.
func Success[S, F any](v S) Result[S, F] {
	return Result[S, F]{
		value:   &container[S]{v},
		failure: nil,
	}
}

// Failure creates a Result with an error.
func Failure[S, F any](v F) Result[S, F] {
	return Result[S, F]{
		value:   nil,
		failure: &container[F]{v},
	}
}

// HandleResult accepts functions to handle a Result when it has a success or when it has an failure.
// This will panic if either of the functions are nil.
func HandleResult[S, F, R any](r Result[S, F], whenSuccess func(S) R, whenFailure func(F) R) R {
	if whenSuccess == nil {
		panic("whenSuccess function must be supplied to HandleResult")
	}
	if whenFailure == nil {
		panic("whenFailure function must be supplied to HandleResult")
	}
	if r.IsFailure() {
		return whenFailure(r.FailureValue())
	}
	return whenSuccess(r.SuccessValue())
}

// Bind applies binder when result is Success and otherwise returns the Failure.
func Bind[S, F, R any](binder func(S) Result[R, F], r Result[S, F]) Result[R, F] {
	if r.IsFailure() {
		return Failure[R](r.FailureValue())
	}
	return binder(r.SuccessValue())
}

// Map applies mapping when result is Success and otherwise returns the Failure.
func Map[S, F, R any](mapping func(S) R, r Result[S, F]) Result[R, F] {
	if r.IsFailure() {
		return Failure[R](r.FailureValue())
	}
	return Success[R, F](mapping(r.SuccessValue()))
}

// MapError applies mapping to the error when the result is Failure and otherwise returns the Success.
func MapError[S, F any](mapping func(F) F, r Result[S, F]) Result[S, F] {
	if r.IsFailure() {
		return Failure[S](mapping(r.FailureValue()))
	}
	return r
}

// DefaultValue returns the value of r if r is Success. Otherwise, it returns success.
func DefaultValue[S, F any](success S, r Result[S, F]) S {
	if r.IsFailure() {
		return success
	}
	return r.SuccessValue()
}

// DefaultWith returns the value of r if r is Success. Otherwise, it returns the output of defThunk.
func DefaultWith[S, F any](defThunk func() S, r Result[S, F]) S {
	return option.DefaultWith(defThunk, r)
}

// Contains tests whether the result contains value.
func Contains[S comparable, F any](value S, r Result[S, F]) bool {
	return option.Contains(value, r)
}

// Count returns 0 if this result is Failure. Otherwise returns 1.
func Count[S, F any](r Result[S, F]) int {
	return option.Count(r)
}

// Exists tests whether the value of r matches the predicate. If the Result is an error, it returns false.
func Exists[S, F any](predicate func(S) bool, r Result[S, F]) bool {
	return option.Exists(predicate, r)
}

// Flatten returns the inner Result when Results are nested.
func Flatten[S, F any](rr Result[Result[S, F], F]) Result[S, F] {
	if rr.IsFailure() {
		return Failure[S](rr.FailureValue())
	}
	return rr.SuccessValue()
}

// Fold applies the folder function to a Result with s being the initial state for the folder.
// If the Result is an Failure, the initial state is returned.
func Fold[S, F, State any](folder func(State, S) State, s State, r Result[S, F]) State {
	return option.Fold(folder, s, r)
}

// FoldBack applies the folder function to a Result with s being in the initial state for the folder.
// If the Result is an Failure, the initial state is returned.
func FoldBack[S, F, State any](folder func(S, State) State, r Result[S, F], s State) State {
	return option.FoldBack(folder, r, s)
}

// ForAll tests whether the value contained in the Result matches the predicate.
// It will always return true if the Result is a Failure.
func ForAll[S, F any](predicate func(S) bool, r Result[S, F]) bool {
	return option.ForAll(predicate, r)
}

// Get returns the value of the Result.
// If Result is a Failure, it panics.
func Get[S, F any](r Result[S, F]) S {
	return option.Get[S](r)
}

// IsNone returns true if the Result is a Failure.
func IsNone[S, F any](r Result[S, F]) bool {
	return r.IsFailure()
}

// IsSome returns true if the Result is Success.
func IsSome[S, F any](r Result[S, F]) bool {
	return r.IsSuccess()
}

// Iter applies the action to the result.
func Iter[S, F any](action func(S), r Result[S, F]) {
	option.Iter(action, r)
}

// Map2 applies function f to two Results and returns the function's return value as a Result.
// If either Result is an Error, it returns the error as the Result.
func Map2[S1, S2, F, R any](f func(S1, S2) R, r1 Result[S1, F], r2 Result[S2, F]) Result[R, F] {
	if r1.IsFailure() {
		return Failure[R](r1.FailureValue())
	}
	if r2.IsFailure() {
		return Failure[R](r2.FailureValue())
	}
	return Success[R, F](f(r1.SuccessValue(), r2.SuccessValue()))
}

// Map3 applies function f to three Results and returns the function's return value as a Result.
// If any of the Results is an Error, it returns the error as the Result.
func Map3[S1, S2, S3, F, R any](f func(S1, S2, S3) R, r1 Result[S1, F], r2 Result[S2, F], r3 Result[S3, F]) Result[R, F] {
	if r1.IsFailure() {
		return Failure[R](r1.FailureValue())
	}
	if r2.IsFailure() {
		return Failure[R](r2.FailureValue())
	}
	if r3.IsFailure() {
		return Failure[R](r3.FailureValue())
	}
	return Success[R, F](f(r1.SuccessValue(), r2.SuccessValue(), r3.SuccessValue()))
}

// OfNullable creates a result from a pointer.
// If the pointer is nil, the result will be a Failure with the the message "nil".
// If the pointer is not nil, the result will be Succeess of the value the pointer points to.
func OfNullable[S, F any](value *S) Result[S, F] {
	if value == nil {
		return Failure[S](*(new(F)))
	}
	return Success[S, F](*value)
}

// OrElse returns r if it is Success or ifNone if r is a Failure.
func OrElse[S, F any](ifNone Result[S, F], r Result[S, F]) Result[S, F] {
	return option.OrElse[S](ifNone, r)
}

// OrElseWith returns r if it is Success or the Result returned from ifNoneThunk if r is an Error.
func OrElseWith[S, F any](ifNoneThunk func() Result[S, F], r Result[S, F]) Result[S, F] {
	return option.OrElseWith[S](ifNoneThunk, r)
}

// ToSlice returns the value in Result as a single item slice.
// If the Result is an Failure, it returns an empty slice.
func ToSlice[S, F any](r Result[S, F]) []S {
	return option.ToSlice[S](r)
}

// ToNullable returns a pointer to the value in the Result if it is Success.
// If the Result is an Failure, it returns nil.
func ToNullable[S, F any](r Result[S, F]) *S {
	return option.ToNullable[S](r)
}

// Lift adapts a function that returns a value and an error into a function
// that returns a Result that will be Success if there is no error and Failure if there is an error.
func Lift[S, F any](f func() (S, error)) func() Result[S, F] {
	return func() Result[S, F] {
		s, err := f()
		if err != nil {
			var fv any = *(new(F))
			switch e := fv.(type) {
			case error:
				return Failure[S](e.(F))
			case string:
				var v any = err.Error()
				return Failure[S](v.(F))
			}
			return Failure[S](fv.(F))
		}
		return Success[S, F](s)
	}
}

// Lift1 adapts a function that accepts one input and returns a value and an error into a function
// that returns a Result that will be Success if there is no error and Failure if there is an error.
func Lift1[T, S, F any](f func(T) (S, error)) func(T) Result[S, F] {
	return func(input T) Result[S, F] {
		s, err := f(input)
		lifted := Lift[S, F](func() (S, error) { return s, err })
		return lifted()
	}
}

// Lift2 adapts a function that accepts two inputs and returns a value and an error into a function
// that returns a Result that will be Success if there is no error and Failure if there is an error.
func Lift2[T1, T2, S, F any](f func(T1, T2) (S, error)) func(T1, T2) Result[S, F] {
	return func(input1 T1, input2 T2) Result[S, F] {
		s, err := f(input1, input2)
		lifted := Lift[S, F](func() (S, error) { return s, err })
		return lifted()
	}
}
