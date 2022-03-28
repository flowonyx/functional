package option

import "errors"

// Optional is an interface for Option-like types.
type Optional[T any] interface {
	IsSome() bool
	IsNone() bool
	Value() T
}

// OptionalCheckOnly is an interface for Option-like types when Value is not required.
type OptionalCheckOnly interface {
	IsSome() bool
	IsNone() bool
}

// HandleOption accepts an Optional value and two functions that return errors.
// whenSome is the function to use when o.IsSome() is true.
// whenNone is the function to use when o.IsNone() is true.
// The error returned by the function that is used will be returned from HandleOption.
func HandleOption[T any, TOptional Optional[T], F func(T) error, FN func() error](o TOptional, whenSome F, whenNone FN) error {
	if whenSome == nil {
		return errors.New("whenSome function must be supplied to HandleOption")
	}
	if whenNone == nil {
		return errors.New("whenNone function must be supplied to HandleOption")
	}
	if o.IsNone() {
		return whenNone()
	}
	return whenSome(o.Value())
}

// HandleOptionIgnoreNone accepts an Optional value and one function that returns an error.
// whenSome is the function to use when o.IsSome() is true.
// The error returned by whenSome will be returned from HandleOptionIgnoreNone or nil will be returned when o.IsNone().
func HandleOptionIgnoreNone[T any, TOptional Optional[T], F func(T) error](o TOptional, whenSome F) error {
	return HandleOption(o, whenSome, func() error { return nil })
}

// DefaultValue tests whether o.IsNone() and returns value if true.
// If o.IsSome(), o.Value() is returned.
func DefaultValue[T any, TOptional Optional[T]](value T, o TOptional) T {
	if o.IsNone() {
		return value
	}
	return o.Value()
}

// DefaultWith tests whether o.IsNone() and returns the result of defThunk if true.
// If o.IsSome(), o.Value() is returned.
func DefaultWith[T any, TOptional Optional[T]](defThunk func() T, o TOptional) T {
	if o.IsNone() {
		return defThunk()
	}
	return o.Value()
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

// Fold applies the folder function to the value in o. If o is None, it will return s.
// If o is Some, it will return the result of the function, with s being passed as the first parameter
// and the value of o being the second parameter.
func Fold[T, State any, TOptional Optional[T]](folder func(State, T) State, s State, o TOptional) State {
	if o.IsNone() {
		return s
	}
	return folder(s, o.Value())
}

// FoldBack applies the folder function to the value in o. If o is None, it will return s.
// If o is Some, it will return the result of the function, with the value of o being passed as the first parameter
// and s being the second parameter.
// It is the same as Fold, but with the parameters swapped.
func FoldBack[T, State any, TOptional Optional[T]](folder func(T, State) State, o TOptional, s State) State {
	if o.IsNone() {
		return s
	}
	return folder(o.Value(), s)
}

// ForAll returns true if either o is None or the predicate returns true when applied to the value of o.
// It returns false only if the predicate returns false.
func ForAll[T any, TOptional Optional[T]](predicate func(T) bool, o TOptional) bool {
	if o.IsNone() {
		return true
	}
	return predicate(o.Value())
}

// Get retrieves the value in o and panics if o is None.
func Get[T any, TOptional Optional[T]](o TOptional) T {
	if o.IsNone() {
		panic("cannot get value of None")
	}
	return o.Value()
}

// IsNone checks if o is None.
func IsNone[T any, TOptional Optional[T]](o TOptional) bool {
	return o.IsNone()
}

// IsSome checks if o is Some.
func IsSome[T any, TOptional Optional[T]](o TOptional) bool {
	return o.IsSome()
}

// Iter applies the action to the value of o. If o is None, this does nothing.
func Iter[T any, TOptional Optional[T]](action func(T), o TOptional) {
	if o.IsNone() {
		return
	}
	action(o.Value())
}

// OrElse returns ifNone if o is None.
// Otherwise, it returns o.
func OrElse[T any, TOptional Optional[T]](ifNone TOptional, o TOptional) TOptional {
	if o.IsNone() {
		return ifNone
	}
	return o
}

// OrElseWith returns the return value of ifNoneThunk if o is None.
// Otherwise it returns o.
func OrElseWith[T any, TOptional Optional[T]](ifNoneThunk func() TOptional, o TOptional) TOptional {
	if o.IsNone() {
		return ifNoneThunk()
	}
	return o
}

// ToSlice creates a single item slice from the value in o.
// If o is None, it returns an empty slice.
func ToSlice[T any, TOptional Optional[T]](o TOptional) []T {
	if o.IsNone() {
		return []T{}
	}
	return []T{o.Value()}
}

// ToNullable returns nil if o is None.
// Otherwise, it returns a pointer the value of o.
func ToNullable[T any, TOptional Optional[T]](o TOptional) *T {
	if o.IsNone() {
		return nil
	}
	v := o.Value()
	return &v
}
