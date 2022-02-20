package simpleMap

import (
	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/option"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

type SimpleMap[K comparable, T any] map[K]T

func withClone[K comparable, T any](m SimpleMap[K, T], f func(SimpleMap[K, T])) SimpleMap[K, T] {
	m2 := maps.Clone(m)
	f(m2)
	return m2
}

func FromSlice[K comparable, T any](s []Pair[K, T]) SimpleMap[K, T] {
	m := SimpleMap[K, T]{}
	Iter(func(t Pair[K, T]) {
		key, value := FromPair(t)
		m[key] = value
	}, s)
	return m
}

func FromSlices[K comparable, T any](keys []K, values []T) SimpleMap[K, T] {
	return FromSlice(Zip(keys, values))
}

func (m SimpleMap[K, T]) Iter(action func(key K, value T)) {
	for k, v := range m {
		action(k, v)
	}
}

// Iteri iterates over the map with a count but no order can assumed.
func (m SimpleMap[K, T]) Iteri(action func(i int, key K, value T)) {
	var i int
	for k, v := range m {
		action(i, k, v)
		i++
	}
}

func (m SimpleMap[K, T]) ToSlice() []Pair[K, T] {
	p := []Pair[K, T]{}
	m.Iter(func(key K, value T) {
		p = append(p, PairOf(key, value))
	})
	return p
}

func (m SimpleMap[K, T]) Set(key K, value T) SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) { m[key] = value })
}

func (m SimpleMap[K, T]) Get(key K) T {
	return m[key]
}

func (m SimpleMap[K, T]) TryGet(key K) option.Option[T] {
	if v, ok := m[key]; ok {
		return option.Some(v)
	}
	return option.None[T]()
}

func (m SimpleMap[K, T]) Clone() SimpleMap[K, T] {
	return maps.Clone(m)
}

func (m SimpleMap[K, T]) Clear() SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) { maps.Clear(m) })
}

func (m SimpleMap[K, T]) CopyFrom(src SimpleMap[K, T]) SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) { maps.Copy(m, src) })
}

func (m SimpleMap[K, T]) Remove(key K) SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) { delete(m, key) })
}

func (m SimpleMap[K, T]) RemoveBy(del func(K, T) bool) SimpleMap[K, T] {
	return withClone(m, func(m SimpleMap[K, T]) {
		for k, v := range m {
			if del(k, v) {
				delete(m, k)
			}
		}
	})
}

func (m SimpleMap[K, T]) Filter(predicate Predicate2[K, T]) SimpleMap[K, T] {
	m2 := make(SimpleMap[K, T])
	for k, v := range m {
		if predicate(k, v) {
			m2[k] = v
		}
	}
	return m2
}

func (m SimpleMap[K, T]) Find(key K) (T, error) {
	if v := m.TryGet(key); v.IsSome() {
		return v.Value(), nil
	}
	return *(new(T)), KeyNotFoundErr
}

func (m SimpleMap[K, T]) FindKey(predicate Predicate2[K, T]) (K, error) {
	for k, v := range m {
		if predicate(k, v) {
			return k, nil
		}
	}
	return *(new(K)), KeyNotFoundErr
}

func (m SimpleMap[K, T]) TryFindKey(predicate Predicate2[K, T]) option.Option[K] {
	key, err := m.FindKey(predicate)
	if err != nil {
		return option.None[K]()
	}
	return option.Some(key)
}

func (m SimpleMap[K, T]) Keys() []K {
	return maps.Keys(m)
}

func (m SimpleMap[K, T]) Values() []T {
	return maps.Values(m)
}

func (m SimpleMap[K, T]) Len() int {
	return len(m)
}

func (m SimpleMap[K, T]) Contains(key K) bool {
	_, ok := m[key]
	return ok
}

func (m SimpleMap[K, T]) Exists(predicate Predicate2[K, T]) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

func (m SimpleMap[K, T]) ForAll(predicate Predicate2[K, T]) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

func (m SimpleMap[K, T]) IsEmpty() bool {
	return len(m) == 0
}

func (m SimpleMap[Key, T]) Partition(predicate Predicate2[Key, T]) (SimpleMap[Key, T], SimpleMap[Key, T]) {
	tmap := SimpleMap[Key, T]{}
	fmap := SimpleMap[Key, T]{}
	for key, value := range m {
		if predicate(key, value) {
			tmap.Set(key, value)
		} else {
			fmap.Set(key, value)
		}
	}
	return tmap, fmap
}

func Fold[K constraints.Ordered, T, State any](folder func(State, K, T) State, initial State, table SimpleMap[K, T]) State {
	return Fold(func(s State, t Pair[K, T]) State {
		return folder(s, t.First, t.Second)
	}, initial, SortBy(func(p Pair[K, T]) K { return p.First }, table.ToSlice()))
}

func FoldBack[K constraints.Ordered, T, State any](folder func(K, T, State) State, initial State, table SimpleMap[K, T]) State {
	return FoldBack(func(t Pair[K, T], s State) State {
		return folder(t.First, t.Second, s)
	}, SortBy(func(p Pair[K, T]) K { return p.First }, table.ToSlice()), initial)
}

func MapTo[K comparable, T, R any](mapping func(K, T) R, table SimpleMap[K, T]) SimpleMap[K, R] {
	output := SimpleMap[K, R]{}
	for key, value := range table {
		output.Set(key, mapping(key, value))
	}
	return output
}
