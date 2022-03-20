package errors

import "errors"

var (
	New    = errors.New
	As     = errors.As
	Is     = errors.Is
	Unwrap = errors.Unwrap
)

type FunctionalError string

const (
	KeyNotFoundErr     = FunctionalError("key not found")
	NotFoundErr        = FunctionalError("not found")
	BadArgumentErr     = FunctionalError("bad argument")
	IndexOutOfRangeErr = FunctionalError("index out of range")
)

func (fe FunctionalError) Error() string {
	return string(fe)
}
