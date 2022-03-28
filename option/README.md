# Functional Options
                                                                                                                                              
Optional types are common in a number of programming languages today. The API for this `Option` type is based on the standard library for F# although it is not completely the same.

# Get it

```sh
go get -u github.com/flowonyx/functional/option
```

# Use it

```go
import "github.com/flowonyx/functional/option"
```

# Types

## Optional interface

There may be a reason to use a different type for optional values but still want to use these functions to work with it. The interface `Optional` is pretty simple:

```go
type Optional[T any] interface {
    IsSome() bool
	IsNone() bool
	Value() T
}
```

There is a smaller interface the leaves out `Value`, and it is used in any functions that only need to test for `IsSome` or `IsNone`.

```go
type OptionalCheckOnly interface {
	IsSome() bool
	IsNone() bool
}
```

## Option type

The `Option` type implements the Optional interface and has some additional functions associated with it.

# `Optional` Functions

* `HandleOption` accepts a input that is an `Optional` value and two functions that return errors.
  * `whenSome` is the function to use when `o.IsSome()` is true.
  * `whenNone` is the function to use when `o.IsNone()` is true.
  * The error returned by the function that is used will be returned from `HandleOption`.
* `HandleOptionIgnoreNone` accepts an `Optional` value and one function that returns an error.
  * `whenSome` is the function to use when `o.IsSome()` is true.
  * The error returned by `whenSome` will be returned from `HandleOption` or `nil` will be returned when `o.IsNone()`.
* `DefaultValue` tests whether `o.IsNone()` and returns the given default value if true.
  * If `o.IsSome()`, `o.Value()` is returned.
* `DefaultWith` tests whether `o.IsNone()` and returns the result of a function if true.
  * If `o.IsSome()`, `o.Value()` is returned.
* `Contains` tests whether the value in the `Optional` is equal to the test value.
* `Count` returns 0 if `o.IsNone()` and 1 if `o.IsSome()`.
* `Exists` tests is the value in the Optional o matches the predicate.
  * If `o.IsNone()`, it will return false.
* `Fold` applies the folder function to the value in an `Optional`.
  * If `o.IsNone()`, it will return the initial state.
  * If `o.IsSome()`, it will return the result of the function.
* `FoldBack` applies the folder function to the value in an `Optional`. If `o.IsNone()`, it will return the initial state.
  * If `o.IsSome()`, it will return the result of the function.
  * It is the same as `Fold`, but with the parameters swapped.
* `ForAll` returns true if either `o.IsNone()` or the predicate function returns true when applied to the value of the `Optional`.
  * It returns false only if the predicate function returns false.
* `Get` retrieves the value in an `Optional` and panics if `o.IsNone()`.
* `IsNone` checks if an `Optional` is `None`.
* `IsSome` checks if `Optional` is `Some`.
* `Iter` applies an action function to the value of an `Optional`. If `o.IsNone()`, this does nothing.
* `OrElse` returns a given `Optional` if `o.IsNone()`. Otherwise, it returns o.
* `OrElseWith` returns the result of a function if `o.IsNone()`. Otherwise it returns o.
* `ToSlice` creates a single item slice from the value in an `Optional`.
  * If `o.IsNone()`, it returns an empty slice.
* `ToNullable` returns `nil` if `o.IsNone()`.
  * Otherwise, it returns a pointer the value of the `Optional`.

# Option Functions

These functions need to have access to the actual type, not just an interface.

* `Some` creates an `Option` with a value.
* `None` creates an `Option` with no value.
* `IsSome` tests if the `Option` contains a value.
* `IsNone` tests whether the `Option` does not contain a value.
* `Value` returns the value in the `Option`. If the `Option` is `None`, it returns the zero value for the type.
* `Bind` applies a function to an `Option` `o` if `o.IsSome()` and otherwise returns `None`.
* `Filter` returns an `Option` if the value in it matches the predicate. Otherwise, it returns `None`.
* `Flatten` takes a nested `Option` and returns the inner `Option`.
* `Map` applies a function to the value of an `Option` and returns the result as an `Option`. If the given `Option` is `None`, it returns `None`.
* `Map2` applies a function to the values in two `Option`s as the first and second parameters and returns the result as an `Option`. If either `Option` is `None`, it returns `None`.
* `Map3` applies a function to the values in three `Option`s as the first, second, and third parameters and returns the result as an `Option`. If any of the `Option`s are `None`, it returns `None`.
* `OfNullable` returns `None` if the supplied pointer is `nil`. Otherwise it returns `Some` of the value (after dereferencing the pointer).
* `Lift` converts a function that returns a value and an error to a function that returns an `Option`.
* `Lift1` converts a function that accepts a single input and returns a value and an error to a function that accepts a single input and returns an `Option`.
* `Lift2` converts a function that accepts two inputs and returns a value and an error to a function that accepts two inputs and returns an `Option`.