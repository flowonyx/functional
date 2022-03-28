[![Go Reference](https://pkg.go.dev/badge/github.com/flowonyx/functional/orderedMap.svg)](https://pkg.go.dev/github.com/flowonyx/functional/orderedMap)

# Functional Ordered Map
                                                                                                                                                                              
There are probably better implementations of an ordered map. This one is very simple, based on a slice of `functional.Pair`s to store the items in order. This order will be in order that items were added or if a comparing function is supplied, in sorted order. It is not safe for concurrent use any more than the standard `map` type.

# Get it

```sh
go get -u github.com/flowonyx/functional/orderedMap
```

# Use it

```go
import "github.com/flowonyx/functional/orderedMap"
```

# Type

There is single generic type in this package: `OrderedMap[KeyType, ValueType]`. It has no public members and is meant to be interacted with from the functions in this package.

# Functions

* `NewOrderedMap` creates a new `OrderedMap`. If a function for comparing key:value pairs is provided, it is used to keep the items in sorted order rather than added order.
* `FromSlice` creates an `OrderedMap` from a slice of key:value `functional.Pair`s. If a function for comparing key:value pairs is provided, the items are sorted.
  * If a key is repeated, the last value for the key is used.
* `ToSlice` exports the map as a slice of key:value `functional.Pair`s.
* `Len` returns the length of the map.
* `Equal` tests whether two `OrderedMaps` are equal. Equality is based on the keys and values all being the same. Order is not considered.
* `EqualBy` tests whether two `OrderedMaps` are equal by applying a predicate function to each value in both maps.
* `Contains` tests whether the given key is present in the map.
* `Exists` tests whether a key:value pair that matches the predicate is present in the map.
* `Set` either adds the key and value to the map or updates the value of the key that is already present.
* `Get` either gets the value associated with the key or returns the zero value of the value type if the key is not present.
* `Remove` removes a key from the map.
* `TryGet` returns an `option.Optional` value where if the key exists, it will be `Some(value)`, otherwise it will be `None`.
* `Filter` returns a new `OrderedMap` with only the values that match a predicate function.
* `Find` is the same as `Get` except that it returns an error if the key is not found.
* `FindKey` returns a key that matches a predicate function or a `KeyNotFoundErr` error.
* `TryFindKey` is the same as `FindKey` but it returns an `option.Option` with `None` if no key matches instead of returning an error.
* `ForAll` tests whether all key:value pairs match the predicate.
* `IsEmpty` tests whether the map is empty.
* `Iter` applies the action to each key:value pair in the map.
* `Iteri` applies the action to each key:value pair in the map, with the index as the first parameter to the action.
* `Keys` returns all the keys in the map.
* `Values` returns all the values in the map.
* `Partition` returns two `OrderedMaps` where the first contains all key:value pairs that match the predicate and the second contains all key:value pairs that do not match the predicate.
* `Fold` applies a folder function to each key:value pair until arriving at the final state.
* `FoldBack` applies a folder function in reverse order to each key:value pair until arriving at the final state.
* `MapTo` creates a new map with the keys and values changed through a projection function.
* `MapValuesTo` creates a new map with the same keys, but the values are changed through a projection function.
* `Pick` searches the map looking for the first element where the given function returns a `option.Some` value. Returns a `KeyNotFoundErr` if no such element exists.
* `TryPick` searches the map looking for the first element where the given function returns a `option.Some` value and returns the `option.Some` value. Returns `option.None` if no such element exists.
* `Set` returns a copy of a map with the given key set to the given value.
* `Remove` returns a copy of a map with the given key removed.