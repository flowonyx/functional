[![Go Reference](https://pkg.go.dev/badge/github.com/flowonyx/functional/maps.svg)](https://pkg.go.dev/github.com/flowonyx/functional/maps)

# Functional Map
                                                                                                                         
This package provides some functions for working Go maps, using generics (at least Go 1.18).

# Get it

```sh
go get -u github.com/flowonyx/functional/maps
```

# Use it

```go
import "github.com/flowonyx/functional/maps"
```

# Terminology

* A `predicate` is a function that returns a boolean (`true` or `false`), generally based on one or two parameters.
* A `projection` just means a function that takes a certain type and returns another type.

# Functions

## Creating Maps

* `Clone` creates a copy of a map.
* `FromSlice` creates a map from a slice of key, value pairs.
* `FromSlices` creates a map by combining a slice of keys with a slice of values.
* `MapTo` creates a new map by mapping each key:value pair in a map with a projection function.
* `Partition` returns two maps from a given map: 
  * The first contains key:value pairs from the map that match a predicate function.
  * The second contains key:value pairs from the map that do not match a predicate function.

## Maps -> Slices

* `ToSlice` converts the map into a slice of key, value `functional.Pair`s.
* `Keys` returns a slice of all the keys in a map.
* `Values` returns a slice of all the values in a map.

## Removal

* `Clear` clears all keys in a map.
* `Remove` returns a copy of a map with a given key deleted.
* `RemoveBy` returns a copy of a map with all keys matched by a predicate function removed.
* `Filter` returns a copy of a map with only key:value pairs that match a predicate function.

## Action

* `CopyFrom` returns a copy of a map with key:values added or replaced from another map.
* `Iter` performs the given action for each key, value pair in the map.
* `Set` returns a copy of the map with the key set to the new value.
  * Of course, if you wanted to update the current map, you would just use map[key]=value.
* `FoldMap` applies a folder function to each key:value pair in a map until it reaches the finished state. The key:value pairs in the map are sorted by key before processing.
* `FoldBackMap` is the same as `FoldMap` but in reverse order.

## Retrieval

* `Get` returns the value of the given key. If the key does not exist, it will be the zero value of the value type.
* `TryGet` returns an `option.Option` value of the given key. If the key does not exist, the returned value will be `None`.
* `Find` either returns the value in a map belonging to a given key or returns a `errors.KeyNotFoundErr` error if the key is not present.
* `FindKey` finds the first key in a map that is matched by a predicate function. Remember that no order can be assumed. If no key is matched by the predicate, it returns a `errors.KeyNotFoundErr` error.
* `TryFindKey` is just like `FindKey` but it returns an `option.Option` with the value of `None` if the key is not found.

## Properties / Tests

* `Contains` tests if a map contains a given key.
* `Exists` tests if a map contains a key:value pair that matches a predicate function.
* `ForAll` tests if all key:value pairs in a map match a predicate function.
* `IsEmpty` tests if the map contains nothing. This is only useful if you need to pass the function around because it just calls `len(m) == 0`.
* `Len` returns the count of items in a map. It just calls the builtin `len` so there is no advantage to using this one. It is only if you need to pass the function around.