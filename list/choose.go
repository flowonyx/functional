package list

import (
	"fmt"

	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/option"
)

// Choose applies chooser to each value in values and if chooser returns Some,
// it includes the returned value in a new slice.
func Choose[T any, R any](chooser func(T) option.Option[R], values []T) []R {
	return Map(option.Option[R].Value, Filter(option.Option[R].IsSome, Map(chooser, values)...))
}

// Pick applies chooser to each value in values until chooser returns Some.
// It then returns the value within Some.
// If no values are matched by chooser, a NotFoundErr is returned.
func Pick[T any, R any](chooser func(T) option.Option[R], values []T) (R, error) {
	var val R
	var found bool

	IterUntil(func(t T) bool {
		if o := chooser(t); o.IsSome() {
			val = o.Value()
			found = true
			return true
		}
		return false
	}, values)

	return val, IfV[error](found, nil).Else(fmt.Errorf("Pick(%v): %w", values, errors.NotFoundErr))
}

// TryPick is the same as Pick but returns None in place of an error.
func TryPick[T any, R any](chooser func(T) option.Option[R], values []T) option.Option[R] {
	if p, err := Pick(chooser, values); err != nil {
		return option.None[R]()
	} else {
		return option.Some(p)
	}
}
