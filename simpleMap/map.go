// Package simpleMap provides a map[Key]Value that has some convenient methods.
// It is not safe for concurrent use any more than the standard map.
package simpleMap

import (
	"fmt"

	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/errors"
	. "github.com/flowonyx/functional/list"
	"github.com/flowonyx/functional/option"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

// SimpleMap provides some convenient methods on top of map.
type SimpleMap[K comparable, T any] map[K]T

func withClone[K comparable, T any](m SimpleMap[K, T], f func(SimpleMap[K, T])) SimpleMap[K, T] {
	m2 := maps.Clone(m)
	f(m2)
	return m2
}

// FromSlice creates a SimpleMap from a slice of key, value pairs.
func FromSlice[K comparable, T any](s []Pair[K, T]) SimpleMap[K, T] {
	m := SimpleMap[K, T]{}
	Iter(func(t Pair[K, T]) {
		key, value := FromPair(t)
		m[key] = value
	}, s)
	return m
}

// FromSlices creates a SimpleMap by combining a slice of keys with a slice of values.
func FromSlices[K comparable, T any](keys []K, values []T) SimpleMap[K, T] {
	return FromSlice(Zip(keys, values))
}

// Iter performs the given action for each key, value pair in the map.
func (m SimpleMap[K, T]) Iter(action func(key K, value T)) {
	for k, v := range m {
		action(k, v)
	}
}

// ToSlice converts the map into a slice of key, value pairs.
func (m SimpleMap[K, T]) ToSlice() []Pair[K, T] {
	p := []Pair[K, T]{}
	m.Iter(func(key K, value T) {
		p = append(p, PairOf(key, value))
	})
	return p
}

// Set returns a copy of the map with the key set to the new value.
func (m SimpleMap[K, T]) Set(key K, value T) SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) { m[key] = value })
}

// Get returns the value of the given key.
// If the key does not exist, it will be the zero value of the value type.
func (m SimpleMap[K, T]) Get(key K) T {
	return m[key]
}

// TryGet returns an optional value of the given key.
// If the key does not exist, the returned value will be None.
func (m SimpleMap[K, T]) TryGet(key K) option.Option[T] {
	if v, ok := m[key]; ok {
		return option.Some(v)
	}
	return option.None[T]()
}

// Clone creates a copy of the map.
func (m SimpleMap[K, T]) Clone() SimpleMap[K, T] {
	return maps.Clone(m)
}

// Clear returns a copy of the map with all keys cleared.
func (m SimpleMap[K, T]) Clear() SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) { maps.Clear(m) })
}

// CopyFrom returns a copy of the map with values copied from src.
func (m SimpleMap[K, T]) CopyFrom(src SimpleMap[K, T]) SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) { maps.Copy(m, src) })
}

// Remove returns a copy of the map with the given key deleted.
func (m SimpleMap[K, T]) Remove(key K) SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) { delete(m, key) })
}

// RemoveBy returns a copy of the map with all keys matched by the del predicate removed.
func (m SimpleMap[K, T]) RemoveBy(del func(K, T) bool) SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) {
		for k, v := range m {
			if del(k, v) {
				delete(m, k)
			}
		}
	})
}

// Filter returns a copy of the map with only key, value pairs matching the given predicate.
func (m SimpleMap[K, T]) Filter(predicate func(K, T) bool) SimpleMap[K, T] {
	m2 := make(SimpleMap[K, T])
	for k, v := range m {
		if predicate(k, v) {
			m2[k] = v
		}
	}
	return m2
}

// Find either returns the value belonging to the key or returns a KeyNotFound error if the key is not present.
func (m SimpleMap[K, T]) Find(key K) (T, error) {
	if v := m.TryGet(key); v.IsSome() {
		return v.Value(), nil
	}
	return *(new(T)), fmt.Errorf("SimpleMap.Find(%v): %w", key, errors.KeyNotFoundErr)
}

// FindKey finds the first key in the map that is matched by the predicate. Remember that no order can be assumed.
// If no key is matched by the predicate, it returns a KeyNotFound error.
func (m SimpleMap[K, T]) FindKey(predicate func(K, T) bool) (K, error) {
	for k, v := range m {
		if predicate(k, v) {
			return k, nil
		}
	}
	return *(new(K)), fmt.Errorf("SimpleMap.FindKey: %w", errors.KeyNotFoundErr)
}

// TryFindKey is just like FindKey but it returns an option with the value of None if the key is not found.
func (m SimpleMap[K, T]) TryFindKey(predicate func(K, T) bool) option.Option[K] {
	key, err := m.FindKey(predicate)
	if err != nil {
		return option.None[K]()
	}
	return option.Some(key)
}

// Keys returns a slice of all the keys in the map.
func (m SimpleMap[K, T]) Keys() []K {
	return maps.Keys(m)
}

// Values returns a slice of all the values in the map.
func (m SimpleMap[K, T]) Values() []T {
	return maps.Values(m)
}

// Len returns the count of items in the map.
func (m SimpleMap[K, T]) Len() int {
	return len(m)
}

// Contains test if the mapt contains the given key.
func (m SimpleMap[K, T]) Contains(key K) bool {
	_, ok := m[key]
	return ok
}

// Exists tests if the map contains a key, value pair that matches the predicate.
func (m SimpleMap[K, T]) Exists(predicate func(K, T) bool) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

// ForAll tests if all key, value pairs in the mapt match the predicate.
func (m SimpleMap[K, T]) ForAll(predicate func(K, T) bool) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

// IsEmpty tests if the map contains nothing.
func (m SimpleMap[K, T]) IsEmpty() bool {
	return len(m) == 0
}

// Partition returns two maps:
// The first contains key, value pairs from this map that match the predicate.
// The second contains key, value pairs from this map that do not match the predicate.
func (m SimpleMap[Key, T]) Partition(predicate func(Key, T) bool) (trueMap SimpleMap[Key, T], falseMap SimpleMap[Key, T]) {
	trueMap = SimpleMap[Key, T]{}
	falseMap = SimpleMap[Key, T]{}
	for key, value := range m {
		if predicate(key, value) {
			trueMap[key] = value
		} else {
			falseMap[key] = value
		}
	}
	return trueMap, falseMap
}

// FoldMap applies the folder function to each key, value pair in table until it reaches the finished state.
// The key, value pairs are sorted by key.
func FoldMap[K constraints.Ordered, T, State any](folder func(State, K, T) State, initial State, table SimpleMap[K, T]) State {
	return Fold(func(s State, t Pair[K, T]) State {
		return folder(s, t.First, t.Second)
	}, initial, SortBy(func(p Pair[K, T]) K { return p.First }, table.ToSlice()))
}

// FoldBackMap applies the folder function to each key, value pair in table in reverse order until it reaches the finished state.
// The key, value pairs are sorted by key.
func FoldBackMap[K constraints.Ordered, T, State any](folder func(K, T, State) State, initial State, table SimpleMap[K, T]) State {
	return FoldBack(func(t Pair[K, T], s State) State {
		return folder(t.First, t.Second, s)
	}, SortBy(func(p Pair[K, T]) K { return p.First }, table.ToSlice()), initial)
}

// MapTo creates a new map from mapping each key, value pair in table with the mapping function.
func MapTo[K comparable, T, R any](mapping func(K, T) R, table SimpleMap[K, T]) SimpleMap[K, R] {
	output := SimpleMap[K, R]{}
	for key, value := range table {
		output[key] = mapping(key, value)
	}
	return output
}
