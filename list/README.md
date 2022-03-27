███████╗██╗   ██╗███╗   ██╗ ██████╗████████╗██╗ ██████╗ ███╗   ██╗ █████╗ ██╗         ██╗     ██╗███████╗████████╗███████╗
██╔════╝██║   ██║████╗  ██║██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║██╔══██╗██║         ██║     ██║██╔════╝╚══██╔══╝██╔════╝
█████╗  ██║   ██║██╔██╗ ██║██║        ██║   ██║██║   ██║██╔██╗ ██║███████║██║         ██║     ██║███████╗   ██║   ███████╗
██╔══╝  ██║   ██║██║╚██╗██║██║        ██║   ██║██║   ██║██║╚██╗██║██╔══██║██║         ██║     ██║╚════██║   ██║   ╚════██║
██║     ╚██████╔╝██║ ╚████║╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║██║  ██║███████╗    ███████╗██║███████║   ██║   ███████║
╚═╝      ╚═════╝ ╚═╝  ╚═══╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝    ╚══════╝╚═╝╚══════╝   ╚═╝   ╚══════╝
                                                                                                                          

This package provides many generic functions for working with slices. The API here is mostly inspired by the `list` library in the F# standard library. There are some minor changes and the implementation is just how I thought it should be done. Several of these functions may not be very useful in Go but I have included them for the sake of having a complete set of functions that can be passed around when needed. You can generally just copy the function definition into your own code. That will also protect against changes in the API which may happen in the future. This is a very early version and may or may not be changed in the future as I see a need.

# Installation

```sh
go get -u github.com/flowonyx/functional/list
```

# Terminology

* A `projection` just means a function that takes a certain type and returns another type.
  * This is often used to return one field from the type for use in different functions.
  * It can also return the whole value wrapped in another type.
  * It can also return something completely different, such as a projection of a string where it returns the length.
* A `predicate` is a function that returns a boolean (`true` or `false`), generally based on one or two parameters.
* Many functions have a variation which begin with `Try`.
  * These functions return `option.Option` values which are `None` if they fail to find something.
  * You can see the `option` package for more information about how to work with this. 
* Some functions have a variation which begin with `Must`.
  * These functions panic instead of returning an error.    
* Some functions have a variation which ends with `By`.
  * These functions take a `projection` function to get the values to use.

# Functions

* `Average` returns the average of the items in any numeric slice.
* `CountBy` applies a projection to each value in a slice and uses the result as the key in map of counts.
* `Sum` returns the result of adding all values in a numeric slice together. (It can also concatenate strings.)
* `SumBy` returns the result of adding all results from a given projection function to each value in a slice.
* `Min` returns the minimum value of all items. The only time an error will be returned is when no values are passed. If you know you are passing it values, you can either ignore the error value or use the `MustMin` variation instead.
* `Max` returns the maximum value of all items. Everything said about `Min` applies to this as well.

* `Average` returns the average of the items in any numeric slice.
* `CountBy` applies a projection to each value in a slice and uses the result as the key in map of counts.
* `Sum` returns the result of adding all values in a numeric slice together. (It can also concatenate strings.)
* `SumBy` returns the result of adding all results from a given projection function to each value in a slice.
* `Min` returns the minimum value of all items. The only time an error will be returned is when no values are passed. If you know you are passing it values, you can either ignore the error value or use the `MustMin` variation instead.
* `Max` returns the maximum value of all items. Everything said about `Min` applies to this as well.

* `Average` returns the average of the items in any numeric slice.
* `CountBy` applies a projection to each value in a slice and uses the result as the key in map of counts.
* `Sum` returns the result of adding all values in a numeric slice together. (It can also concatenate strings.)
* `SumBy` returns the result of adding all results from a given projection function to each value in a slice.
* `Min` returns the minimum value of all items. The only time an error will be returned is when no values are passed. If you know you are passing it values, you can either ignore the error value or use the `MustMin` variation instead.
* `Max` returns the maximum value of all items. Everything said about `Min` applies to this as well.

## Selecting and checking for existence

* `Choose` returns a slice of values from an input slice that match the given predicate function (it does not return `None`).
* `Pick` returns the first value from an input slice that matches the given predicate function (it does not return `None`).
* `Contains` tests whether a value is within a slice.
* `Distinct` returns a copy of a slice with any duplicate values removed.
* `Except` returns values that are not in a slice of values to exclude.
* `Exists` tests whether any value in a slice matches a predicate.
* `Filter` returns the values in a slice that match a predicate.
* `Find` and its variants returns the first value that matches a predicate. (It is the same as `Pick` except that the predicate does not return an `Option`.)
* `FindBack` and and its variants returns the last value that matches a predicate.
* `ForAll` tests whether all values in a slice match a predicate.
* `IndexOf` returns the first index within a slice that matches a value. There are several variants of this, including `IndexOfBack` which searches backwards from the end of the slice.
* `Skip` returns a clone of a slice after skipping the number of items requested. It will panic if count is greater than the length of the slice.
* `SkipWhile` returns a clone of a slice starting from where a provided predicate returns false.
* `Take` returns an `option.Option` of the first values in a slice up to a given count. It returns an `option.Option` so it will not need to return an error if count is out of bounds. (See `Truncate` for ignoring the error.)
  * `TakeWhile` returns the first values in a slice until a given predicate function returns false.
  * `Truncate` returns the first values in a slice up to a given count. If the count exceeds the number of values in the slice, it will return all the values.

## Rearranging slices

* `AllPairs` returns a new slice that contains all pairings of elements from two slices.
* `ChunkBySize` returns a two dimensional slice where each sub slice is a "chunk" of the input slice of the specified size.
* `Collect` takes a projection function that returns a slice for each item in a given slice. After applying the projection to each item, `Collect` then returns all the generated slices concatenated together.
* `Concat` accepts any number of slices and concatenates them into a single slice.
* `GroupBy` applies a projection to each value in a slice and returns a slice of `functional.Pair`s where each `Pair` is a key and a slice of values for which the projection returned that key.
  * You can get the results returned as a map by using `GroupByAsMap`.
* `Indexed` converts a slice into a a slice of `functional.Pair`s where the first part is the index and the second part is the item.
* `Pairwise` returns a slice of each item in the input slice and the one before it, except for the first item which is only returned as the one before the second element.
* `Partition` splits the slice into two slices. The first slice contains the items for which the given predicate returns `true`, and the second slice contains the items for which the given predicate returns `false` respectively.
* `Reverse` returns a clone of a slice with the order of the items reversed.
* `Sort` returns a clone of a slice of any ordered type in sorted order.
  * `SortDescending` returns a clone of a slice of any ordered type in reverse sorted order.
  * `SortBy` returns a clone of a slice of any type (it does not have to be `constraints.Ordered`) sorted in ascending order based on the key returned from a projection function.
  * `SortByDescending` is obviously the same as `SortBy` but in reverse order.
  * `SortWith` returns a clone of a slice of any type (it does not have to be `constraints.Ordered`) sorted in the order as determined by the `less` function you provide.
* `SplitAt` splits a slice at a given index into two separate slices.
* `SplitInto` splits a slice into a series of slices of whatever length is necessary to have the given number of slices.
* `Windowed` returns the values in sliding windows of a specified size.
* `Zip` puts two slices into one slice of `functional.Pair`s.
  * `Unzip` takes a slice of `functional.Pair`s and returns two slices as they would have been before a `Zip` operation.
  * `Zip3` and `Unzip3` are the same except that they work on three slices and `functional.Triple`s.

## Generating slices and setting indexes

* `Cons` takes a Head and a Tail and puts them together into one slice.
* `Create` and its variations creates a slice with all values set to the supplied value.
* `CreateZero` and its variations creates a slice with all values set to the zero value of the slice.
* `CreateFromStructure2D` and `CreateFromStructure3D` creates a 2 dimensional or 3 dimensional slice of the same dimensions as the slice passed in.
* `Empty` makes a slice with the given cap but length of 0. This is helpful for when you need to use `append` to fill out the slice because it is difficult to figure out the index but you know the size ahead of time.
* `Fill` fills the range of items in a slice from a start index for the specified number of items with a given value.
* `InitSlice` makes a new slice of a specified length and initializes the values with the return value of an initializer function which is given the index of the item.
  * There are several variants for different dimensions of slices.
* `InsertAt` and `RemoveAt` and their variants just simplify the syntax for placing items or removing them at specific indexes in a slice.
* `SetItem` and its variants (`SetItem2D` etc.) are equivalent to simply setting an index in a slice to the given value except that instead of panicking, it returns an `errors.IndexOutOfRangeErr` error.
* `UpdateItem` and its variants are the same as `SetItem` but they return a clone of the slice with the change updated instead of changing the original slice.
* `Singleton` creates a slice with a single value in it.
  * `ExactlyOne` checks if a slice has exactly one value in it and it does, returns the slice. If it does not, it returns a `errors.BadArgumentErr` error.

## Comparing slices

* `Equal` compares two slices and returns true if they have the same values in the same order.
* `EqualUnordered` compares two slices and returns true if they have the same values in any order.
* `MinLen` returns the minimum length of any number of slices.
* `MinSlice` returns the slice of any number of slices that has the minimum length.

## Properties of a slice

* `Len2`, `Len3`, and `Len4` gets the minimum length of the sub slices in the two, three, or four dimensional slice.
* `LastIndexOf` is basically just `len(slice)-1`, but I like the way it reads better and I don't have to worry about forgetting the `-1`.
* `Head` gets the first item in a slice. It will return an error if the slice is empty. If you want to ignore the error, you can use `MustHead` instead.
* `Last` gets the last item in a slice. It will return an error if the slice is empty. If you want to ignore the error, you can use `MustLast` instead.
* `Tail` gets the items in the slice after the first item (Head). If the slice is empty, it will just return an empty slice.

# Operations over a slice

* `Average` returns the average of the items in any numeric slice.
* `CountBy` applies a projection to each value in a slice and uses the result as the key in map of counts.
* `Fold` and its variants apply a folding function to each item in a slice and keep changing the state each time. It then returns the final state.
* `Iter` iterates over all items in a slice, applying a function to each value that returns nothing. There are several variants, including `IterRev` which iterates backwards over the slice.
* `IterUntil` iterates over each item in a slice until the function applied to it returns `true`.
* `Map` and its variants applies a mapping function to each item in a slice and returns the results as a new slice. The returned slice can be of a completely different type from the original slice.
* `Permute` returns a slice with all items rearranged by function that accepts the original index and returns the new index.
* `Range` creates a slice of Integers from a specified start to a specified end. If a step value is specified, the values will be spaced by that amount.
* `RangeTo` is the same as `Range` except you do not need to specify the start value--it will start at 0 and the step will always be 1.
  * The variants that begin with `DoRange` execute a function which is passed each number in the range.
* `Sum` returns the result of adding all values in a numeric slice together. (It can also concatenate strings.)
* `SumBy` returns the result of adding all results from a given projection function to each value in a slice.
* `Transpose` returns the transpose of the sequence of slices.
* `Min` returns the minimum value of all items. The only time an error will be returned is when no values are passed. If you know you are passing it values, you can either ignore the error value or use the `MustMin` variation instead.
* `Max` returns the maximum value of all items. Everything said about `Min` applies to this as well.