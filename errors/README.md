███████╗██╗   ██╗███╗   ██╗ ██████╗████████╗██╗ ██████╗ ███╗   ██╗ █████╗ ██╗       
██╔════╝██║   ██║████╗  ██║██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║██╔══██╗██║       
█████╗  ██║   ██║██╔██╗ ██║██║        ██║   ██║██║   ██║██╔██╗ ██║███████║██║       
██╔══╝  ██║   ██║██║╚██╗██║██║        ██║   ██║██║   ██║██║╚██╗██║██╔══██║██║       
██║     ╚██████╔╝██║ ╚████║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║██║  ██║███████╗  
╚═╝      ╚═════╝ ╚═╝  ╚═══╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝  
                                                                                                                                       
███████╗██████╗ ██████╗  ██████╗ ██████╗ ███████╗
██╔════╝██╔══██╗██╔══██╗██╔═══██╗██╔══██╗██╔════╝
█████╗  ██████╔╝██████╔╝██║   ██║██████╔╝███████╗
██╔══╝  ██╔══██╗██╔══██╗██║   ██║██╔══██╗╚════██║
███████╗██║  ██║██║  ██║╚██████╔╝██║  ██║███████║
╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝

This package provides just a few error constants that are used in different sub packages under the `functional` package. You probably do not need to install this package directly.

# Error Type

All of these errors are of the type `FunctionalError` and are constants (possible because `FunctionalError` is an alias for `string`).

# Errors

* `KeyNotFoundErr` is used by functions that check for a key when that key is not found.
* `NotFoundErr` is used by functions that check for a value when that value is not found.
* `BadArgumentErr` is used by functions that receive parameters (arguments) that are not valid for the function to use.
* `IndexOutOfRangeErr` is used by functions that take an index as a parameter when it is out of range for its use.

# Error Functions

This package aliases these functions from the standard `errors` package so it does not need to be imported separately.

* `New`
* `As`
* `Is`
* `Unwrap`

# Checking for specific errors

Generally, the functions that return these errors wrap them in another error that has more information about the error. This means that while these errors are constants, it will often not work to test by:

```go
if err == KeyNotFoundErr {...}
```

Instead, use the `Is` function to test:

```go
if errors.Is(err, KeyNotFoundErr) {...}
```