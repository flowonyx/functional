// Package simpleMap provides a map[Key]Value that has some convenient methods.
package maps

import (
	"fmt"

	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/list"
	"github.com/flowonyx/functional/option"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

func withClone[K comparable, T any](m map[K]T, f func(map[K]T)) map[K]T {
	m2 := maps.Clone(m)
	f(m2)
	return m2
}

// FromSlice creates a map from a slice of key, value pairs.
func FromSlice[K comparable, T any](s []Pair[K, T]) map[K]T {
	m := map[K]T{}
	list.Iter(func(t Pair[K, T]) {
		key, value := FromPair(t)
		m[key] = value
	}, s)
	return m
}

// FromSlices creates a map by combining a slice of keys with a slice of values.
func FromSlices[K comparable, T any](keys []K, values []T) map[K]T {
	return FromSlice(list.Zip(keys, values))
}

// Iter performs the given action for each key, value pair in the map.
func Iter[K comparable, T any](action func(key K, value T), m map[K]T) {
	for k, v := range m {
		action(k, v)
	}
}

// ToSlice converts the map into a slice of key, value functional.Pairs.
func ToSlice[K comparable, T any](m map[K]T) []Pair[K, T] {
	p := []Pair[K, T]{}
	Iter(func(key K, value T) {
		p = append(p, PairOf(key, value))
	}, m)
	return p
}

// Set returns a copy of the map with the key set to the new value.
func Set[K comparable, T any](key K, value T, m map[K]T) map[K]T {
	return withClone(m, func(m map[K]T) { m[key] = value })
}

// Get returns the value of the given key.
// If the key does not exist, it will be the zero value of the value type.
func Get[K comparable, T any](key K, m map[K]T) T {
	return m[key]
}

// TryGet returns an optional value of the given key.
// If the key does not exist, the returned value will be None.
func TryGet[K comparable, T any](key K, m map[K]T) option.Option[T] {
	if v, ok := m[key]; ok {
		return option.Some(v)
	}
	return option.None[T]()
}

// Clone creates a copy of the map.
func Clone[K comparable, T any](m map[K]T) map[K]T {
	return maps.Clone(m)
}

// Clear clears all keys in a map.
func Clear[K comparable, T any](m map[K]T) {
	maps.Clear(m)
}

// CopyFrom returns a copy of the map with values copied from src.
func CopyFrom[K comparable, T any](src map[K]T, dest map[K]T) map[K]T {
	return withClone(dest, func(m map[K]T) { maps.Copy(m, src) })
}

// Remove returns a copy of the map with the given key deleted.
func Remove[K comparable, T any](key K, m map[K]T) map[K]T {
	return withClone(m, func(m map[K]T) { delete(m, key) })
}

// RemoveBy returns a copy of the map with all keys matched by the del predicate removed.
func RemoveBy[K comparable, T any](del func(K, T) bool, m map[K]T) map[K]T {
	return withClone(m, func(m map[K]T) {
		for k, v := range m {
			if del(k, v) {
				delete(m, k)
			}
		}
	})
}

// Filter returns a copy of the map with only key, value pairs matching the given predicate.
func Filter[K comparable, T any](predicate func(K, T) bool, m map[K]T) map[K]T {
	m2 := make(map[K]T)
	for k, v := range m {
		if predicate(k, v) {
			m2[k] = v
		}
	}
	return m2
}

// Find either returns the value belonging to the key or returns a KeyNotFoundErr error if the key is not present.
func Find[K comparable, T any](key K, m map[K]T) (T, error) {
	if v := TryGet(key, m); v.IsSome() {
		return v.Value(), nil
	}
	return *(new(T)), fmt.Errorf("maps.Find(%v): %w", key, errors.KeyNotFoundErr)
}

// FindKey finds the first key in the map that is matched by the predicate. Remember that no order can be assumed.
// If no key is matched by the predicate, it returns a KeyNotFoundErr error.
func FindKey[K comparable, T any](predicate func(K, T) bool, m map[K]T) (K, error) {
	for k, v := range m {
		if predicate(k, v) {
			return k, nil
		}
	}
	return *(new(K)), fmt.Errorf("maps.FindKey: %w", errors.KeyNotFoundErr)
}

// TryFindKey is just like FindKey but it returns an option with the value of None if the key is not found.
func TryFindKey[K comparable, T any](predicate func(K, T) bool, m map[K]T) option.Option[K] {
	key, err := FindKey(predicate, m)
	if err != nil {
		return option.None[K]()
	}
	return option.Some(key)
}

// Keys returns a slice of all the keys in the map.
func Keys[K comparable, T any](m map[K]T) []K {
	return maps.Keys(m)
}

// Values returns a slice of all the values in the map.
func Values[K comparable, T any](m map[K]T) []T {
	return maps.Values(m)
}

// Len returns the count of items in the map.
func Len[K comparable, T any](m map[K]T) int {
	return len(m)
}

// Contains test if the map contains the given key.
func Contains[K comparable, T any](key K, m map[K]T) bool {
	_, ok := m[key]
	return ok
}

// Exists tests if the map contains a key, value pair that matches the predicate.
func Exists[K comparable, T any](predicate func(K, T) bool, m map[K]T) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

// ForAll tests if all key, value pairs in the map match the predicate.
func ForAll[K comparable, T any](predicate func(K, T) bool, m map[K]T) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

// IsEmpty tests if the map contains nothing.
func IsEmpty[K comparable, T any](m map[K]T) bool {
	return len(m) == 0
}

// Partition returns two maps:
// The first contains key, value pairs from this map that match the predicate.
// The second contains key, value pairs from this map that do not match the predicate.
func Partition[K comparable, T any](predicate func(K, T) bool, m map[K]T) (trueMap map[K]T, falseMap map[K]T) {
	trueMap = map[K]T{}
	falseMap = map[K]T{}
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
func FoldMap[K constraints.Ordered, T, State any](folder func(State, K, T) State, initial State, table map[K]T) State {
	return list.Fold(func(s State, t Pair[K, T]) State {
		return folder(s, t.First, t.Second)
	}, initial, list.SortBy(func(p Pair[K, T]) K { return p.First }, ToSlice(table)))
}

// FoldBackMap applies the folder function to each key, value pair in table in reverse order until it reaches the finished state.
// The key, value pairs are sorted by key.
func FoldBackMap[K constraints.Ordered, T, State any](folder func(K, T, State) State, initial State, table map[K]T) State {
	return list.FoldBack(func(t Pair[K, T], s State) State {
		return folder(t.First, t.Second, s)
	}, list.SortBy(func(p Pair[K, T]) K { return p.First }, ToSlice(table)), initial)
}

// MapTo creates a new map from mapping each key, value pair in table with the mapping function.
func MapTo[K comparable, T, R any](mapping func(K, T) R, table map[K]T) map[K]R {
	output := map[K]R{}
	for key, value := range table {
		output[key] = mapping(key, value)
	}
	return output
}
