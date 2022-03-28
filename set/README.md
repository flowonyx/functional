# Functional Sets
                                                                                                               
The package provides a generic set type and functions that work with it. It is based on the `orderedMap.OrderedMap` type which means it is not safe for concurrent usage.

# Get it

```sh
go get -u github.com/flowonyx/functional/set
```

# Use it

```go
import "github.com/flowonyx/functional/set"
```

# Types

The only type here is `Set[T comparable]`. The interesting part is in the functions that interact with it.

# Functions

* `NewSet` creates a new `Set`. If a comparison function is provided, it is used in ordering the set.
* `Singleton` creates a set of exactly one item.
  * This does not actually prevent you from adding more items later, so this is a way to create a set and add the first item in one call.
* `FromSlice` creates a new `Set` from the items in the given slice.

# Methods

* `ToSlice` returns the values in the `Set` as a slice of items.
* `Add` adds an item to the `Set`. If it already exists in the set, nothing changes.
* `Remove` removes an item from the `Set`.
* `Equal` tests if two sets are equal.
* `Contains` tests whether item is present in the `Set`.
* `Exists` tests whether any item in the `Set` matches the predicate.
* `Count` returns the number of items in the `Set`.
* `IsEmpty` test whether this is an empty set.
* `IsSubsetOf` tests whether this `Set` is a subset of the `Set` passed as a parameter.
  * A `Set` is a subset of another `Set` if it only contains items that are also present in the other `Set`.
  * An empty `Set` is considered a subset of any other `Set`. This also considers equal sets to be subsets of each other.
* `IsProperSubsetOf` tests whether this `Set` is a subset of the `Set` passed as a parameter.
  * A `Set` is a subset of another `Set` if it only contains items that are also present in the other `Set`.
  * An empty `Set` is not considered a subset of any other `Set`. Also, if the `Set`s are equal, neither is considered to a subset of the other.
* `IsProperSupersetOf` tests whether this `Set` is a superset of the `Set` passed as a parameter.
  * A `Set` is a superset of another `Set` if it contains all items that are present in the other `Set`.
  * If either `Set` is empty or if they are equal, this will not consider the `Set` to be a superset.
* `IsSupersetOf` tests whether this `Set` is a superset of the `Set` passed as a parameter.
  * A `Set` is a superset of another `Set` if it contains all items that are present in the other `Set`.
  * If the other `Set` is empty or if they are equal, this will consider the `Set` to be a superset.
* `IndexOf` finds the index of the item within the `Set`.
  * The order of items within the `Set` is the order in which the items were added or sorted order if a comparison function was supplied when the `Set` was created.
* `Items` returns the items in the `Set` as a slice.
* `Difference` returns a `Set` containing the items that are only present in one `Set` or the other.
* `Intersect` returns a `Set` containing the items that are present in both `Set`s.
* `Union` returns a `Set` that contains all items that are present in either `Set`.
* `Filter` returns a `Set` that contains only items from this `Set` that match a predicate function.
* `ForAll` tests whether all items in the set match a predicate function.
* `Iter` applies an action function to each item in the `Set`.
* `Iteri` applies an action function to each item in the `Set`, also passing the index to the action.
* `Partition` returns two `Set`s. The first `Set` are those items in this `Set` that match a predicate function. The second `Set` are those items in this `Set` that do not match the predicate.
* `Fold` applies the folder function to each item in the `Set` until it reaches the final state.
* `FoldBack` applies the folder function in reverse order to each item in the `Set` until it reaches the final state.
* `DifferenceMany` finds the difference between all the sets.
* `ItersectMany` finds the intersection between all the sets.
* `UnionMany` finds the union between all the sets.
* `Map` applies the mapping function to each item in the `Set` s.
* `MaxElement` finds the largest item in the `Set` s.
* `MinElement` finds the smallest item in the `Set` s.
* `MaxElementBy` finds the largest item in the `Set` s using the return values from a projection function for comparison.
* `MinElementBy` finds the smallest item in the `Set` s using the return values from a projection function for comparison.