// Package set provides a generic Set type.
package set

import (
	"github.com/flowonyx/functional"
	"github.com/flowonyx/functional/list"
	"github.com/flowonyx/functional/orderedMap"
	"golang.org/x/exp/constraints"
)

// Set is a type keeps a set of values and allows set operations on them.
type Set[T comparable] struct {
	m orderedMap.OrderedMap[T, struct{}]
}

// NewSet creates a new Set.
// If lessFunc is provided, it is used in ordering the set.
func NewSet[T comparable](lessFunc ...func(T, T) bool) Set[T] {
	var lf func(functional.Pair[T, struct{}], functional.Pair[T, struct{}]) bool
	if len(lessFunc) > 0 {
		lf = func(p functional.Pair[T, struct{}], p2 functional.Pair[T, struct{}]) bool {
			return lessFunc[0](p.First, p2.First)
		}
	}
	return Set[T]{m: orderedMap.NewOrderedMap(lf)}
}

// Singleton creates a set of exactly one item.
// This does not actually prevent you from adding more items later,
// so this is a way to creat a set and add the first item in one call.
func Singleton[T comparable](item T) Set[T] {
	s := NewSet[T]()
	s.Add(item)
	return s
}

func (s Set[T]) clone() Set[T] {
	n := NewSet[T]()
	for _, i := range s.Items() {
		n.Add(i)
	}
	return n
}

// FromSlice creates a new Set from the items in the given slice.
func FromSlice[T comparable](input []T, lessFunc ...func(T, T) bool) Set[T] {
	s := NewSet(lessFunc...)
	for _, i := range input {
		s.Add(i)
	}
	return s
}

// ToSlice returns the values in the Set as a slice of items.
func (s Set[T]) ToSlice() []T {
	return s.Items()
}

// Add adds an item to the Set. If it already exists in the set, nothing changes.
func (s *Set[T]) Add(item T) {
	s.m.Set(item, struct{}{})
}

// Remove removes an item from the Set.
func (s *Set[T]) Remove(item T) {
	s.m.Remove(item)
}

// Equal tests if two sets are equal.
func (s Set[T]) Equal(s2 Set[T]) bool {
	return orderedMap.Equal(s.m, s2.m)
}

// Contains tests whether item is present in the Set.
func (s Set[T]) Contains(item T) bool {
	return s.m.Contains(item)
}

// Exists tests whether any item in the Set matches the predicate.
func (s Set[T]) Exists(predicate func(T) bool) bool {
	return list.Exists(predicate, s.Items()...)
}

// Count returns the number of items in the Set.
func (s Set[T]) Count() int {
	return s.m.Len()
}

// IsEmpty test whether this is an empty set.
func (s Set[T]) IsEmpty() bool {
	return s.Count() == 0
}

// IsSubsetOf tests whether this Set is a subset of the Set passed as a parameter.
// A Set is a subset of another Set if it only contains items that are also present in the other Set.
// An empty Set is considered a subset of any other Set. This also considers equal sets to be subsets of each other.
// If you do not want these cases to be considered a subset, use IsProperSubsetOf instead.
func (s Set[T]) IsSubsetOf(potentialSuperset Set[T]) bool {
	if potentialSuperset.Count() < s.Count() {
		return false
	}
	for _, i := range s.Items() {
		if !potentialSuperset.Contains(i) {
			return false
		}
	}
	return true
}

// IsProperSubsetOf tests whether this Set is a subset of the Set passed as a parameter.
// A Set is a subset of another Set if it only contains items that are also present in the other Set.
// An empty Set is not considered a subset of any other Set. Also, if the Sets are equal, neither is considered
// to a subset of the other. If you do want these cases to be considered a subset, use IsSubsetOf instead.
func (s Set[T]) IsProperSubsetOf(potentialSuperset Set[T]) bool {
	if potentialSuperset.IsEmpty() || s.IsEmpty() {
		return false
	}
	if potentialSuperset.Count() <= s.Count() {
		return false
	}
	for _, i := range s.Items() {
		if !potentialSuperset.Contains(i) {
			return false
		}
	}
	return true
}

// IsProperSupersetOf tests whether this Set is a superset of the Set passed as a parameter.
// A Set is a superset of another Set if it contains all items that are present in the other Set.
// If either Set is empty or if they are equal, this will not consider the Set to be a superset.
// If you want these cases to be considered as supersets, use IsSupersetOf instead.
func (s Set[T]) IsProperSupersetOf(potentialSubset Set[T]) bool {
	return potentialSubset.IsProperSubsetOf(s)
}

// IsSupersetOf tests whether this Set is a superset of the Set passed as a parameter.
// A Set is a superset of another Set if it contains all items that are present in the other Set.
// If the other Set is empty or if they are equal, this will consider the Set to be a superset.
// If you do not want these cases to be considered as supersets, use IsProperSupersetOf instead.
func (s Set[T]) IsSupersetOf(potentialSubset Set[T]) bool {
	return potentialSubset.IsSubsetOf(s)
}

// IndexOf finds the index of the item within the Set.
// The order of items within the Set is the order in which the items were added.
func (s Set[T]) IndexOf(item T) int {
	index := -1
	s.m.Iteri(func(i int, t T, _ struct{}) {
		if t == item {
			index = i
		}
	})
	return index
}

// Items returns the items in the Set as a slice.
func (s Set[T]) Items() []T {
	return s.m.Keys()
}

// Difference returns a Set containing the items that are only present in one Set or the other.
func (s Set[T]) Difference(set2 Set[T]) Set[T] {
	output := NewSet[T]()
	for _, i := range s.Items() {
		if !set2.Contains(i) {
			output.Add(i)
		}
	}
	for _, i := range set2.Items() {
		if !s.Contains(i) {
			output.Add(i)
		}
	}
	return output
}

// Intersect returns a Set containing the items that are present in both Sets.
func (s Set[T]) Intersect(set2 Set[T]) Set[T] {
	output := NewSet[T]()
	for _, i := range s.Items() {
		if set2.Contains(i) {
			output.Add(i)
		}
	}
	for _, i := range set2.Items() {
		if s.Contains(i) {
			output.Add(i)
		}
	}
	return output
}

// Union returns a Set that contains all items that are present in either Set.
func (s Set[T]) Union(set2 Set[T]) Set[T] {
	output := NewSet[T]()
	for _, i := range s.Items() {
		output.Add(i)
	}
	for _, i := range set2.Items() {
		output.Add(i)
	}
	return output
}

// Filter returns a Set that contains only items from this Set that match the predicate.
func (s Set[T]) Filter(predicate func(T) bool) Set[T] {
	output := NewSet[T]()
	for _, i := range s.Items() {
		if predicate(i) {
			output.Add(i)
		}
	}
	return output
}

// ForAll tests whether all items in the set match the predicate.
func (s Set[T]) ForAll(predicate func(T) bool) bool {
	return list.ForAll(predicate, s.Items())
}

// Iter applies the action to each item in the Set.
func (s Set[T]) Iter(action func(T)) {
	for _, i := range s.Items() {
		action(i)
	}
}

// Iteri applies the action to each item in the Set,
// also passing the index to the action.
func (s Set[T]) Iteri(action func(int, T)) {
	for i, v := range s.Items() {
		action(i, v)
	}
}

// Partition returns two Sets. The first Set are those items in this Set
// that match the predicate. The second Set are those items in this Set
// that do not match the predicate.
func (s Set[T]) Partition(predicate func(T) bool) (trueSet Set[T], falseSet Set[T]) {
	trueSet = NewSet[T]()
	falseSet = NewSet[T]()
	s.Iter(func(t T) {
		if predicate(t) {
			trueSet.Add(t)
		} else {
			falseSet.Add(t)
		}
	})
	return trueSet, falseSet
}

// Fold applies the folder function to each item in the Set until it reaches the final state.
func Fold[T comparable, State any](folder func(State, T) State, initial State, s Set[T]) State {
	return list.Fold(folder, initial, s.Items())
}

// FoldBack applies the folder function in reverse order to each item in the Set until it reaches the final state.
func FoldBack[T comparable, State any](folder func(T, State) State, s Set[T], initial State) State {
	return list.FoldBack(folder, s.Items(), initial)
}

// DifferenceMany finds the difference between all the sets.
func DifferenceMany[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}
	output := sets[0].clone()
	for _, s := range sets[1:] {
		output = output.Difference(s)
	}
	return output
}

// ItersectMany finds the intersection between all the sets.
func IntersectMany[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}
	output := sets[0].clone()
	for _, s := range sets[1:] {
		output = output.Intersect(s)
	}
	return output
}

// UnionMany finds the union between all the sets.
func UnionMany[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}
	output := sets[0].clone()
	for _, s := range sets[1:] {
		output = output.Union(s)
	}
	return output
}

// Map applies the mapping function to each item in the Set s.
func Map[T comparable, R comparable](mapping func(T) R, s Set[T]) Set[R] {
	output := NewSet[R]()
	s.Iter(func(t T) { output.Add(mapping(t)) })
	return output
}

// MaxElement finds the largest item in the Set s.
func MaxElement[T constraints.Ordered](s Set[T]) (T, error) {
	return list.Max(s.Items()...)
}

// MinElement finds the smallest item in the Set s.
func MinElement[T constraints.Ordered](s Set[T]) (T, error) {
	return list.Min(s.Items()...)
}

// MaxElementBy finds the largest item in the Set s using the return values from projection for comparison.
func MaxElementBy[T comparable, R constraints.Ordered](projection func(T) R, s Set[T]) (T, error) {
	return list.MaxBy(projection, s.Items()...)
}

// MinElementBy finds the smallest item in the Set s using the return values from projection for comparison.
func MinElementBy[T comparable, R constraints.Ordered](projection func(T) R, s Set[T]) (T, error) {
	return list.MinBy(projection, s.Items()...)
}
