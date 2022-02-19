package functional

import "github.com/flowonyx/functional/option"

func Choose[T any, R any](chooser Projection[T, option.Option[R]], input []T) []R {
	return Map(option.Option[R].Value, Filter(option.Option[R].IsSome, Map(chooser, input)))
}

func Pick[T any, R any](chooser Projection[T, option.Option[R]], input []T) (R, error) {
	var val R
	var found bool

	IterUntil(func(t T) bool {
		if o := chooser(t); o.IsSome() {
			val = o.Value()
			found = true
			return true
		}
		return false
	}, input)

	return val, IfV[error](found, nil).Else(KeyNotFoundErr)
}

func TryPick[T any, R any](chooser Projection[T, option.Option[R]], input []T) option.Option[R] {
	if p, err := Pick(chooser, input); err != nil {
		return option.None[R]()
	} else {
		return option.Some(p)
	}
}
