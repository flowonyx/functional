// Package orderedMap offers a map-like type that keeps it's items in either the order they are added or
// in sorted order if it is provided a less function at creation.
package orderedMap

import (
	"fmt"

	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/list"
	"github.com/flowonyx/functional/option"
	"golang.org/x/exp/slices"
)

// OrderedMap is a map-like structure which keeps items in order.
type OrderedMap[Key comparable, T any] struct {
	pairs []Pair[Key, T]
	less  func(Pair[Key, T], Pair[Key, T]) bool
}

// NewOrderedMap creates a new OrderedMap. If lessFunc is provided, it is used to keep the items in sorted order.
// If lessFunc is not provided, the items are kept in the order in which they are added.
func NewOrderedMap[Key comparable, T any](lessFunc ...func(Pair[Key, T], Pair[Key, T]) bool) OrderedMap[Key, T] {
	var lf func(Pair[Key, T], Pair[Key, T]) bool
	if len(lessFunc) > 0 {
		lf = lessFunc[0]
	}
	return OrderedMap[Key, T]{
		pairs: make([]Pair[Key, T], 0),
		less:  lf,
	}
}

// FromSlice creates an OrderedMap from a slice of Key, Value Pairs. If lessFunc is provided, the items are sorted.
// If a key is repeated, the last value for the key is used.
func FromSlice[Key comparable, T any](s []Pair[Key, T], lessFunc ...func(Pair[Key, T], Pair[Key, T]) bool) OrderedMap[Key, T] {
	var lf func(Pair[Key, T], Pair[Key, T]) bool
	if len(lessFunc) > 0 {
		lf = lessFunc[0]
	}
	// We reverse it to get the latest values for the keys, then reverse it again to get the right order if there is no lessFunc.
	pairs := list.DistinctBy(func(p Pair[Key, T]) Key { return p.First }, list.Reverse(s)...)
	if lf == nil {
		pairs = list.Reverse(pairs)
	} else {
		pairs = list.SortWith(lf, pairs)
	}
	m := OrderedMap[Key, T]{
		pairs: pairs,
		less:  lf,
	}
	return m
}

// ToSlice exports the map as a slice of Key, Value Pairs.
func (m OrderedMap[Key, T]) ToSlice() []Pair[Key, T] {
	return slices.Clone(m.pairs)
}

func (m OrderedMap[Key, T]) clone() OrderedMap[Key, T] {
	return OrderedMap[Key, T]{pairs: slices.Clone(m.pairs)}
}

// Len returns the length of the map.
func (m OrderedMap[Key, T]) Len() int {
	return len(m.pairs)
}

// Equal tests whether two OrderedMaps are equal.
// Equality is based on the keys and values all being the same.
// Order is not considered.
func Equal[Key, T comparable](m, m2 OrderedMap[Key, T]) bool {
	return list.EqualUnordered(m.pairs, m2.pairs)
}

// EqualBy tests whether two OrderedMaps are equal by applying a predicate to each value in both maps.
func EqualBy[Key comparable, T any](predicate func(T, T) bool, m, m2 OrderedMap[Key, T]) bool {
	vals1 := list.Map(func(v Pair[Key, T]) T { return v.Second }, m.pairs)
	vals2 := list.Map(func(v Pair[Key, T]) T { return v.Second }, m2.pairs)
	return list.ForAll2(predicate, vals1, vals2)
}

// Contains tests whether the given key is present in the map.
func (m OrderedMap[Key, T]) Contains(key Key) bool {
	return m.indexOf(key) >= 0
}

func (m OrderedMap[Key, T]) indexOf(key Key) int {
	return list.IndexBy(func(p Pair[Key, T]) bool { return p.First == key }, m.pairs)
}

// Exists tests whether a key, value pair that matches the predicate is present in the map.
func (m OrderedMap[Key, T]) Exists(predicate func(Key, T) bool) bool {
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if predicate(key, value) {
			return true
		}
	}
	return false
}

// Set either adds the key and value to the map or
// updates the value of the key that is already present.
func (m *OrderedMap[Key, T]) Set(key Key, value T) {
	if index := m.indexOf(key); index >= 0 {
		p := m.pairs[index]
		p.Second = value
		m.pairs[index] = p
		if m.less != nil {
			m.pairs = list.SortWith(m.less, m.pairs)
		}
		return
	}
	m.pairs = append(m.pairs, PairOf(key, value))
	if m.less != nil {
		m.pairs = list.SortWith(m.less, m.pairs)
	}
}

// Get either gets the value associated with the key
// or returns the zero value of the value type if the
// key is not present.
func (m OrderedMap[Key, T]) Get(key Key) T {
	if index := m.indexOf(key); index >= 0 {
		return m.pairs[index].Second
	}
	return *(new(T))
}

// Remove removes a key from the map.
func (m *OrderedMap[Key, T]) Remove(key Key) {
	if index := m.indexOf(key); index >= 0 {
		var err error
		m.pairs, err = list.RemoveAt(index, m.pairs)
		if err != nil {
			panic(err)
		}
	}
}

// TryGet returns an optional value where if the key exists, it will be Some(value),
// otherwise it will be None.
func (m OrderedMap[Key, T]) TryGet(key Key) option.Option[T] {
	if index := m.indexOf(key); index >= 0 {
		return option.Some(m.pairs[index].Second)
	}
	return option.None[T]()
}

// Filter returns a new OrderedMap with only the values that match the predicate.
func (m OrderedMap[Key, T]) Filter(predicate func(Key, T) bool) OrderedMap[Key, T] {
	output := NewOrderedMap[Key, T]()
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if predicate(key, value) {
			output.Set(key, value)
		}
	}
	return output
}

// Find is the same as Get except that it returns an error if the key is not found.
func (m OrderedMap[Key, T]) Find(key Key) (T, error) {
	if index := m.indexOf(key); index >= 0 {
		return m.pairs[index].Second, nil
	}
	return *(new(T)), fmt.Errorf("OrderedMap.Find(%v): %w", key, errors.KeyNotFoundErr)
}

// FindKey returns a key that matches the predicate or a KeyNotFound error.
func (m OrderedMap[Key, T]) FindKey(predicate func(Key, T) bool) (Key, error) {
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if predicate(key, value) {
			return key, nil
		}
	}
	return *(new(Key)), fmt.Errorf("OrderedMap.FindKey: %w", errors.KeyNotFoundErr)
}

// TryFindKey is the same as FindKey but it returns an Option with None if no key matches instead of returning an error.
func (m OrderedMap[Key, T]) TryFindKey(predicate func(Key, T) bool) option.Option[Key] {
	if value, err := m.FindKey(predicate); err != nil {
		return option.None[Key]()
	} else {
		return option.Some(value)
	}
}

// ForAll tests whether all key, value pairs match the predicate.
func (m OrderedMap[Key, T]) ForAll(predicate func(Key, T) bool) bool {
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if !predicate(key, value) {
			return false
		}
	}
	return true
}

// IsEmpty tests whether the map is empty.
func (m OrderedMap[Key, T]) IsEmpty() bool {
	return len(m.pairs) == 0
}

// Iter applies the action to each key, value pair in the map.
func (m OrderedMap[Key, T]) Iter(action func(Key, T)) {
	for _, p := range m.pairs {
		key, value := FromPair(p)
		action(key, value)
	}
}

// Iteri applies the action to each key, value pair in the map, with the index as the first parameter to the action.
func (m OrderedMap[Key, T]) Iteri(action func(int, Key, T)) {
	for i, p := range m.pairs {
		key, value := FromPair(p)
		action(i, key, value)
	}
}

// Keys returns all the keys in the map.
func (m OrderedMap[Key, T]) Keys() []Key {
	output := make([]Key, len(m.pairs))
	m.Iteri(func(i int, key Key, _ T) { output[i] = key })
	return output
}

// Values returns all the values in the map.
func (m OrderedMap[Key, T]) Values() []T {
	output := make([]T, len(m.pairs))
	m.Iteri(func(i int, _ Key, value T) { output[i] = value })
	return output
}

// Partition returns two OrderedMaps where the first contains all key, value pairs that match the predicate
// and the second contains all key, value pairs that do not match the predicate.
func (m OrderedMap[Key, T]) Partition(predicate func(Key, T) bool) (trueMap OrderedMap[Key, T], falseMap OrderedMap[Key, T]) {
	trueMap = NewOrderedMap[Key, T]()
	falseMap = NewOrderedMap[Key, T]()
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if predicate(key, value) {
			trueMap.Set(key, value)
		} else {
			falseMap.Set(key, value)
		}
	}
	return trueMap, falseMap
}

// Fold applies the folder function to each key, value pair until arriving at the final state.
func Fold[Key comparable, T, State any](folder func(State, Key, T) State, initial State, table OrderedMap[Key, T]) State {
	return list.Fold(func(s State, t Pair[Key, T]) State {
		return folder(s, t.First, t.Second)
	}, initial, table.pairs)
}

// FoldBack applies the folder function in reverse order to each key, value pair until arriving at the final state.
func FoldBack[Key comparable, T, State any](folder func(Key, T, State) State, table OrderedMap[Key, T], initial State) State {
	return list.FoldBack(func(t Pair[Key, T], s State) State {
		return folder(t.First, t.Second, s)
	}, table.pairs, initial)
}

// MapTo creates a new map with the keys and values changed through the mapping function.
func MapTo[Key, KeyR comparable, T, ValueR any](mapping func(Key, T) (KeyR, ValueR), table OrderedMap[Key, T]) OrderedMap[KeyR, ValueR] {
	output := NewOrderedMap[KeyR, ValueR]()
	for _, p := range table.pairs {
		key, value := FromPair(p)
		output.Set(mapping(key, value))
	}
	return output
}

// MapValuesTo creates a new map with the same keys, but the values are changed through the mapping function.
func MapValuesTo[Key comparable, T, R any](mapping func(Key, T) R, table OrderedMap[Key, T]) OrderedMap[Key, R] {
	output := NewOrderedMap[Key, R]()
	for _, p := range table.pairs {
		key, value := FromPair(p)
		output.Set(key, mapping(key, value))
	}
	return output
}

// Pick searches the map looking for the first element where the given function returns a Some value. Returns a KeyNotFoundErr if no such element exists.
func Pick[Key comparable, T, R any](chooser func(Key, T) option.Option[R], table OrderedMap[Key, T]) (R, error) {
	for _, p := range table.pairs {
		key, value := FromPair(p)
		v := chooser(key, value)
		if v.IsSome() {
			return v.Value(), nil
		}
	}
	return *(new(R)), fmt.Errorf("Pick: %w", errors.KeyNotFoundErr)
}

// TryPick searches the map looking for the first element where the given function returns a Some value and returns the Some value. Returns None if no such element exists.
func TryPick[Key comparable, T, R any](chooser func(Key, T) option.Option[R], table OrderedMap[Key, T]) option.Option[R] {
	for _, p := range table.pairs {
		key, value := FromPair(p)
		v := chooser(key, value)
		if v.IsSome() {
			return v
		}
	}
	return option.None[R]()
}

// Set returns a copy of table with the key set the value.
func Set[Key comparable, T any](table OrderedMap[Key, T], key Key, value T) OrderedMap[Key, T] {
	m := table.clone()
	m.Set(key, value)
	return m
}

// Remove returns a copy of table with the key removed.
func Remove[Key comparable, T any](table OrderedMap[Key, T], key Key) OrderedMap[Key, T] {
	m := table.clone()
	m.Remove(key)
	return m
}
