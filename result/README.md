[![Go Reference](https://pkg.go.dev/badge/github.com/flowonyx/functional/result.svg)](https://pkg.go.dev/github.com/flowonyx/functional/result)

# Functional Results

A `Result` type is similar to an `Option` but instead of `Some` and `None`, we need it to have `Success` and `Failure` where both have values associated with them.

# Get it

```sh
go get -u github.com/flowonyx/functional/result
```

# Use it

```go
import "github.com/flowonyx/functional/result"
```

# Types

`Result[SuccessType, FailureType]` is a generic type, where the types of the values associated with `Success` and `Failure` must be specified. This means that with a `Failure`, you can use `error` types or `string`s (or any other type you want).

# Functions

* `IsSuccess` tests if this `Result` has a value.
* `IsFailure` tests if this `Result` has a failure.
* `IsSome` is an alias for `IsSuccess` to satisfy the `option.Optional` interface.
* `IsNone` is an alias for `IsFailure` to satisfy the `option.Optional` interface.
* `SuccessValue` returns the value if this `Result` is `Success`. If the result is an error, `Value` returns the zero value of the value type.
* `FailureValue` returns the error if this `Result` is `Error`. Otherwise, it returns `nil`.
* `Value` is an alias for `SuccessValue` to satisfy the `option.Optional` interface.
* `Success` creates a `Result` with a value.
* `Failure` creates a `Result` with an error.
* `HandleResult` accepts functions to handle a `Result` when it has a success or when it has an failure.
* `Bind` applies a projection function from `SuccessType`->`Result[AnyOtherType, _]` when a `Result` is `Success` and otherwise returns the `Failure`.
* `Map` applies a projection function from `SuccessType`->`AnyOtherType` when `Result` is `Success` and otherwise returns the `Failure`.
* `MapError` applies mapping to the error when the `Result` is `Failure` and otherwise returns the `Success`.
* `DefaultValue` returns the value of of a `Result` if it is `Success`. Otherwise, it returns the supplied default value.
* `DefaultWith` returns the value of a `Result` if it is `Success`. Otherwise, it returns the output of a supplied function.
* `Contains` tests whether the `Result` contains value.
* `Count` returns 0 if this `Result` is `Failure`. Otherwise returns 1.
* `Exists` tests whether the value of a `Result` matches a predicate function. If the `Result` is an error, it returns false.
* `Flatten` returns the inner `Result` when `Result`s are nested.
* `Fold` applies a folder function to a `Result` with an initial state for the folder.  If the `Result` is a `Failure`, the initial state is returned.
* `FoldBack` applies the folder function to a `Result` with s being in the initial state for the folder. If the `Result` is a `Failure`, the initial state is returned.
* `ForAll` tests whether the value contained in a `Result` matches a predicate function. It will always return true if the `Result` is a `Failure`.
* `Get` returns the value of the `Result`. If `Result` is a `Failure`, it panics.
* `IsNone` returns true if the `Result` is a `Failure`.
* `IsSome` returns true if the `Result` is `Success`.
* `Iter` applies an action function to the `Result`.
* `Map2` applies a function to two `Result`s and returns the function's return value as a `Result`. If either `Result` is a `Failure`, it returns the error as the `Result`.
* `Map3` applies a function to three `Result`s and returns the function's return value as a `Result`. If any of the `Result`s is a `Failure`, it returns the error as the `Result`.
* `OfNullable` creates a `Result` from a pointer.
  * If the pointer is `nil`, the `Result` will be a `Failure` with the the message "nil".
  * If the pointer is not `nil`, the `Result` will be `Succeess` of the value the pointer points to.
* `OrElse` returns the given `Result` if it is a `Success` or the supplied value if it is a `Failure`.
* `OrElseWith` returns the given `Result` if it is `Success` or the `Result` returned from a supplied function if it is a `Failure`.
* `ToSlice` returns the value in `Result` as a single item slice.
* If the `Result` is a `Failure`, it returns an empty slice.
* `ToNullable` returns a pointer to the value in the `Result` if it is `Success`. If the `Result` is an `Failure`, it returns `nil`.
* `Lift` adapts a function that returns a value and an error into a function that returns a `Result` that will be `Success` if there is no error and `Failure` if there is an error.
* `Lift1` adapts a function that accepts one input and returns a value and an error into a function that returns a `Result` that will be `Success` if there is no error and `Failure` if there is an error.
* `Lift2` adapts a function that accepts two inputs and returns a value and an error into a function that returns a `Result` that will be `Success` if there is no error and `Failure` if there is an error.