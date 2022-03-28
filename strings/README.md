███████╗██╗   ██╗███╗   ██╗ ██████╗████████╗██╗ ██████╗ ███╗   ██╗ █████╗ ██╗         ███████╗████████╗██████╗ ██╗███╗   ██╗ ██████╗ ███████╗
██╔════╝██║   ██║████╗  ██║██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║██╔══██╗██║         ██╔════╝╚══██╔══╝██╔══██╗██║████╗  ██║██╔════╝ ██╔════╝
█████╗  ██║   ██║██╔██╗ ██║██║        ██║   ██║██║   ██║██╔██╗ ██║███████║██║         ███████╗   ██║   ██████╔╝██║██╔██╗ ██║██║  ███╗███████╗
██╔══╝  ██║   ██║██║╚██╗██║██║        ██║   ██║██║   ██║██║╚██╗██║██╔══██║██║         ╚════██║   ██║   ██╔══██╗██║██║╚██╗██║██║   ██║╚════██║
██║     ╚██████╔╝██║ ╚████║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║██║  ██║███████╗    ███████║   ██║   ██║  ██║██║██║ ╚████║╚██████╔╝███████║
╚═╝      ╚═════╝ ╚═╝  ╚═══╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝    ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝
                                                                                                                                             
This package has functions for working with `string`s, `rune`s, and types that are based on them. Many of them are just wrapping functions from the standard `strings` package, `strconv`, or `golang.org/x/text/cases` but uses generics so casting is unnecessary for types based on `string`s or `rune`s and some can work on either `string`s or `rune`s.

# Get it

```sh
go get -u github.com/flowonyx/functional/strings
```

# Use it

```go
import "github.com/flowonyx/functional/strings"
```

# Types

The types from the standard library's `strings` package are aliased in this package so we can just import the one package.

# Functions

Because most of the functions are just wrapping the functions from the standard `strings` package, I will just note some of the functions here.

* `Collect` accepts a function which maps each rune in the given string to a string, then concatenates them together into one string.
* `Collecti`: the difference between `Collecti` and `Collect` is that the mapping function receives the index of the rune.
* `Concat` concatenates a list of strings or runes into one string.
* `Exists` checks for the existence of a rune within a string that matches a predicate function.
* `Filter` returns a string built from the given string that only contains runes matching a predicate function.
* `ForAll` checks whether every rune in a given string matches a predicate function.
* `InitString` creates a string of a given length where each rune is initialized by an initializer function.
* `Iter` performs an action for each rune in a string, passing the rune to the action function.
* `Iteri` performs an action for each rune in a string, passing the rune and its index to the action function.
* `Join` joins items of any type into a string of values separated by whatever string or rune you choose.
  * It uses fmt.Sprint to represent each elem as a string.
  * If you are passing runes or strings, it is slightly faster to call `JoinRunes` or `JoinStrings` instead.
* `LastIndexRune` returns the index of the last instance of a rune in a string, or -1 if it is not present.
* `Lines` splits a string on newline boundaries into a slice of strings. The results do not include the newlines.
* `GetLine` gets the value of a specific line in a string indicated by index. An error is returned if the index is out of range.
  * Use `GetLineOpt` when you want an `option.Option` instead of an error.
* `Lower` returns a string with all lowercase letters based on the English language.
  * If you need to use another language, call `LowerSpecialCase` directly.
* `Mapi` does the same as `Map` except that the mapping function is supplied with the index of the rune within the string along with the rune.
  * If mapping returns a negative value, the character is dropped from the string with no replacement.
* `Range` creates a slice of runes that between the start and end runes (inclusive).
  * `Range('A', 'C')` would create `[]rune{'A', 'B', 'C'}`.
* `Title` returns a string with English title casing. It uses an approximation of the default Unicode Word Break algorithm.
  * If you want to have the title casing specific to another language, use `TitleSpecial` instead.
* `Upper` returns a string with English upper casing. It uses an approximation of the default Unicode Word Break algorithm.
  * If you want to have the upper casing specific to another language, use `UpperSpecial` instead.
* `Unquote` wraps `strconv.Unquote` and just returns the original string in case of an error.