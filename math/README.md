# Functional Math
                                                                                                                          
It is doubtful that this package will be useful to many. It is mostly wrapping some of the functions from the standard library `math` and `strconv` packages to make them generic and take care of ugly casting for the caller (you can use `int` and not have to cast it to `float64` and back). Below are the functions that are here. As I am pretty much just wrapping standard functions, I did not write any tests.

# Get it

```sh
go get -u github.com/flowonyx/functional/math
```

# Use it

```go
import "github.com/flowonyx/functional/math"
```

# Functions

* `Abs` returns the absolute value of x.
* `RoundInt` returns the nearest integer as an int, rounding half away from zero.
* `Round` returns the nearest integer as the float type of x, rounding half away from zero.
* `RoundToEven` returns the nearest integer as the float type of x, rounding ties to even.
* `RoundToEvenInt` returns the nearest integer as an int, rounding ties to even.
* `Cbrt` returns the cube root of x.
* `CopySign` returns a value with the magnitude of x and the sign of y.
* `Dim` returns the maximum of x-y or 0.
* `Pow` returns x**y, the base-x exponential of y.
* `Pow10` returns 10**n, the base-10 exponential of n.
* `Remainder` returns the IEEE 754 floating-point remainder of x/y.
* `Sqrt` returns the square root of a number.
* `Max` returns the maximum value of two numbers.
* `Min` returns the minimum value of two numbers.
* `TryParseInt` returns the Integer parsed from a string as an `option.Option`. If parsing fails, it returns `None`.