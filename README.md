[![Go Reference](https://pkg.go.dev/badge/github.com/flowonyx/functional.svg)](https://pkg.go.dev/github.com/flowonyx/functional)

# Functional

`functional` is a Go package for functional programming with generics (it requires at least Go 1.18). Most of the functionality is found in the sub packages.

It is heavily inspired by the standard library API for F#. It is not one-to-one equivalent but most functions that are available in the F# standard library have an equivalent here. There are examples for most functions, which is also mostly how it is tested. You can see the [documentation here](https://pkg.go.dev/github.com/flowonyx/functional).

This type of programming is not necessarily considered "idiomatic" programming in the Go language, but I find it useful for coding the way I think and also for translating algorithms from functional languages into Go.

I make no claim to this being stable or useful. The API may change as I see a need for my own purposes. However, many of the types and functions can be copied easily into your own code and used without taking a dependency on this package.

# Get it

```sh
go get -u github.com/flowonyx/functional
```

# Use it

You will just need to import the packages you want to use. Below is the full set of packages in this repository.

```go
import (
    // this package: basic types and high level functions
    "github.com/flowonyx/functional"
    // standard errors that are used by different packages
    "github.com/flowonyx/functional/errors"
    // functions for working with slices
    "github.com/flowonyx/functional/list"
    // provides functions for working with the builtin map type
    "github.com/flowonyx/functional/maps"
    // wraps the standard math package functions to make them generic
    // and get rid of the need for casting (not sure how useful it is)
    "github.com/flowonyx/functional/math"
    // provides an Option type and functions that go with it
    "github.com/flowonyx/functional/option"
    // provides an OrderedMap type that works in a similar way to map but
    // keeps the entries in order (either order they are added or sorted order)
    "github.com/flowonyx/functional/orderedMap"
    // provides a Result type with Success or Failure and related functions
    "github.com/flowonyx/functional/result"
    // provides a generic Set based on the OrderedMap
    "github.com/flowonyx/functional/set"
    // strings provides generic functions for working with strings, runes, and types based on them
    "github.com/flowonyx/functional/strings"
)
```

This is the top-level package. Here you will find some basic types and high level functions.

# Tuple Types

Tuples allow you to pass around pairs or triples of values without creating special structs for them. It is a very common pattern in many programming languages.

* `Pair[T1, T2 any]` is a basic tuple type with two items.
* `Triple[T1, T2, T3 any]` is a basic tuple type with three items.

# Tuple functions

* `PairOf(T1, T2) Pair[T1, T2]` creates a Pair type.
* `TripleOf(T1, T2, T3) Triple[T1, T2, T3]` creates a Triple type.
* `FromPair(Pair[T1, T2]) (T1, T2)` returns the two values in the Pair.
* `FromTriple(Triple[T1, T2, T3]) (T1, T2, T3)` returns the three values in the Triple.

# Curry Functions

Currying is the process of turning a function that takes parameters into a function that already has some parameters set and takes fewer parameters. Many of the functions in these packages were designed with currying in mind. While it might make more sense at times for the parameters to be in a different order, I tried to put the parameters that would be more likely to be curried at the beginning of the parameter list.

* `Curry`, `Curry2`, and `Curry3` accept functions that have 1, 2, and 3 parameters respectively with the values for those parameters and return a function that has no parameters.
* `Curry2To1`, `Curry3to2`, and `Curry3to1` accept functions that have 2 or 3 parameters and return a function that has 1 or 2 parameters as the name implies.
* `Curry2To1F` accepts a function with 2 parameters and returns a function that has 1 parameter and returns a function with no parameters.
* Curry functions that end with `_0` accept functions with no return value.
* Curry functions that end with `_2` accept functions with 2 return values.

# Swap Parameters

* `SwapParams0` takes a function with 2 parameters and no return value (the 0 is for the number of values returned) and returns a function in which the parameters are swapped.
* `SwapParams1` and `SwapParams2` are the same but have 1 or 2 return values.

# Ternary function (If->ElIf->Else)

There are two different styles of ternary functions. Neither is probably a good idea if Go. It is always going to be faster to use the builtin `if` statements. However, there may be some cases, where this is useful to you.

* The first kind takes a boolean test and a function to call for the result if it is `true`.
  * `If(bool, resultIfFunc).ElIf(bool, resultElseIfFunc).Else(resultElseFunc)`
* The second kind takes a boolean test and value for the result if it is `true`.
  * `IfV(bool, resultIf).ElIf(bool, resultElseIf).Else(resultElse)`

# Similar Work

* https://github.com/samber/lo has a few overlapping functions but also some different ones that you might find useful.
* https://github.com/BooleanCat/go-functional also has several similar functions and types.
* https://github.com/rjNemo/underscore is a port of underscore.js so there are similar functions here.
* https://github.com/chyroc/go-lambda has a few functions that I think all have equivalents here.
* https://github.com/bullgare/funktional
* https://github.com/peterzeller/go-fun
* https://github.com/life4/genesis has some interesting Async functions.
* https://github.com/xyctruth/stream is a port of Java Streams.
* https://github.com/szmcdull/glinq is a port DotNet Linq.
* ... There are many others that are similar.

# Sub Packages

Most of the work is done by the sub packages.

* [errors](./errors)
  * Has very few error constants that are used (generally wrapped by other errors) by the other packages here.
* [list](./list)
  * This is where functions live for working with generic slices. I named it `list` to mirror the terminology in F# as most of these functions are inspired by the API in the builtin  list library for F#.
* [maps](./maps)
  * This provides some functions for working with generic maps.
* [math](./math)
  * This mostly wraps the functions from the standard library `math` package so that it can take numbers of different types and return numbers of different types without casting (on the part of the caller).
* [option](./option)
  * This provides a generic Option type where something can either be Some(value) or None.
  * It also provides many functions for interacting with Options and types that fit the same interface.
* [orderedMap](./orderedMap)
  * This provides a generic map type that keeps items in order: either the order in which they were added or a provided sorted order.
  * This is not as well implemented as it probably could be, but it works well enough for my purposes.
* [result](./result)
  * This provides a generic Result type where something can either be Success(value) or Error(error) where Success values and errors can be any type you want.
  * It also provides many functions for interacting with Results.
* [set](./set)
  * This provides a generic set type which is built on the `orderedMap`.
  * It also provides many functions for interacting with Sets.
* [strings](./strings)
  * This provides functions for working with strings, runes, and types that are aliases for them.
  * I believe it wraps all the functions in the builtin `strings` package and also several from `strconv`.