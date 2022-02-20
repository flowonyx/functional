package errors

type FunctionalError string

const (
	KeyNotFoundErr = FunctionalError("key not found")
	BadArgumentErr = FunctionalError("bad argument")
)

func (fe FunctionalError) Error() string {
	return string(fe)
}
