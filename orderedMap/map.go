package orderedMap

import (
	"sync"

	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/option"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type OrderedMap[Key comparable, T any] struct {
	pairs []Pair[Key, T]
	keys  map[Key]int
	mu    *sync.Mutex
	less  func(Key, Key) bool
}

func NewOrderedMap[Key comparable, T any](lessFunc ...func(Key, Key) bool) OrderedMap[Key, T] {
	var lf func(Key, Key) bool
	if len(lessFunc) > 0 {
		lf = lessFunc[0]
	}
	return OrderedMap[Key, T]{
		pairs: make([]Pair[Key, T], 0),
		keys:  make(map[Key]int),
		mu:    new(sync.Mutex),
		less:  lf,
	}

}

func FromSlice[Key comparable, T any](s []Pair[Key, T], lessFunc ...func(Key, Key) bool) OrderedMap[Key, T] {
	var lf func(Key, Key) bool
	if len(lessFunc) > 0 {
		lf = lessFunc[0]
	}
	m := OrderedMap[Key, T]{
		pairs: slices.Clone(s),
		keys:  make(map[Key]int),
		mu:    new(sync.Mutex),
		less:  lf,
	}
	if !m.sort() {
		for i := range m.pairs {
			key, _ := FromPair(m.pairs[i])
			m.keys[key] = i
		}
	}
	return m
}

func (m OrderedMap[Key, T]) sort() bool {
	if m.less != nil {
		pairs := make([]Pair[Key, T], len(m.pairs))
		keys := SortWith(m.less, m.Keys())
		for i, k := range keys {
			pairs[i] = m.pairs[m.keys[k]]
			m.keys[pairs[i].First] = i
		}
		m.pairs = pairs
		return true
	}
	return false
}

func (m OrderedMap[Key, T]) ToSlice() []Pair[Key, T] {
	m.lock()
	defer m.unlock()
	return slices.Clone(m.pairs)
}

func (m OrderedMap[Key, T]) lock() {
	m.mu.Lock()
}

func (m OrderedMap[Key, T]) unlock() {
	m.mu.Unlock()
}

func (m OrderedMap[Key, T]) clone() OrderedMap[Key, T] {
	m.lock()
	defer m.unlock()
	keys := make(map[Key]int)
	for k, v := range m.keys {
		keys[k] = v
	}
	return OrderedMap[Key, T]{pairs: slices.Clone(m.pairs), keys: keys, mu: new(sync.Mutex)}
}

func (m OrderedMap[Key, T]) Len() int {
	m.lock()
	defer m.unlock()
	return len(m.pairs)
}

func Equal[Key, T comparable](m, m2 OrderedMap[Key, T]) bool {
	if !maps.Equal(m.keys, m2.keys) {
		return false
	}
	for k := range m.keys {
		if m2.Get(k) != m.Get(k) {
			return false
		}
	}
	return true
}

func EqualBy[Key comparable, T any](predicate Predicate2[T, T], m, m2 OrderedMap[Key, T]) bool {
	if !maps.Equal(m.keys, m2.keys) {
		return false
	}
	for k := range m.keys {
		if !predicate(m.Get(k), m2.Get(k)) {
			return false
		}
	}
	return true
}

func (m OrderedMap[Key, T]) Contains(key Key) bool {
	m.lock()
	defer m.unlock()
	_, ok := m.keys[key]
	return ok
}

func (m OrderedMap[Key, T]) Exists(predicate Predicate2[Key, T]) bool {
	m.lock()
	defer m.unlock()
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if predicate(key, value) {
			return true
		}
	}
	return false
}

func (m OrderedMap[Key, T]) Set(key Key, value T) {
	m.lock()
	defer m.unlock()
	if index, ok := m.keys[key]; ok {
		p := m.pairs[index]
		p.Second = value
		m.pairs[index] = p
	}
	m.pairs = append(m.pairs, PairOf(key, value))
	m.keys[key] = len(m.pairs) - 1
	m.sort()
}

func (m OrderedMap[Key, T]) Get(key Key) T {
	m.lock()
	defer m.unlock()
	if index, ok := m.keys[key]; ok {
		return m.pairs[index].Second
	}
	return *(new(T))
}

func (m OrderedMap[Key, T]) Remove(key Key) {
	m.lock()
	defer m.unlock()
	if index, ok := m.keys[key]; ok {
		var err error
		m.pairs, err = RemoveAt(index, m.pairs)
		if err != nil {
			panic(err)
		}
		delete(m.keys, key)
		m.sort()
	}
}

func (m OrderedMap[Key, T]) TryGet(key Key) option.Option[T] {
	m.lock()
	defer m.unlock()
	if index, ok := m.keys[key]; ok {
		return option.Some(m.pairs[index].Second)
	}
	return option.None[T]()
}

func (m OrderedMap[Key, T]) Filter(predicate Predicate2[Key, T]) OrderedMap[Key, T] {
	m.lock()
	defer m.unlock()
	output := NewOrderedMap[Key, T]()
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if predicate(key, value) {
			output.Set(key, value)
		}
	}
	return output
}

func (m OrderedMap[Key, T]) Find(key Key) (T, error) {
	m.lock()
	defer m.unlock()
	if index, ok := m.keys[key]; ok {
		return m.pairs[index].Second, nil
	}
	return *(new(T)), KeyNotFoundErr
}

func (m OrderedMap[Key, T]) FindKey(predicate Predicate2[Key, T]) (Key, error) {
	m.lock()
	defer m.unlock()
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if predicate(key, value) {
			return key, nil
		}
	}
	return *(new(Key)), KeyNotFoundErr
}

func (m OrderedMap[Key, T]) TryFindKey(predicate Predicate2[Key, T]) option.Option[Key] {
	if value, err := m.FindKey(predicate); err != nil {
		return option.None[Key]()
	} else {
		return option.Some(value)
	}
}

func (m OrderedMap[Key, T]) ForAll(predicate Predicate2[Key, T]) bool {
	m.lock()
	defer m.unlock()
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if !predicate(key, value) {
			return false
		}
	}
	return true
}

func (m OrderedMap[Key, T]) IsEmpty() bool {
	return len(m.pairs) == 0
}

func (m OrderedMap[Key, T]) Iter(action func(Key, T)) {
	m.lock()
	defer m.unlock()
	for _, p := range m.pairs {
		key, value := FromPair(p)
		action(key, value)
	}
}

func (m OrderedMap[Key, T]) Iteri(action func(int, Key, T)) {
	m.lock()
	defer m.unlock()
	for i, p := range m.pairs {
		key, value := FromPair(p)
		action(i, key, value)
	}
}

func (m OrderedMap[Key, T]) Keys() []Key {
	output := make([]Key, len(m.pairs))
	m.Iteri(func(i int, key Key, _ T) { output[i] = key })
	return output
}

func (m OrderedMap[Key, T]) Values() []T {
	output := make([]T, len(m.pairs))
	m.Iteri(func(i int, _ Key, value T) { output[i] = value })
	return output
}

func (m OrderedMap[Key, T]) Partition(predicate Predicate2[Key, T]) (OrderedMap[Key, T], OrderedMap[Key, T]) {
	tmap := NewOrderedMap[Key, T]()
	fmap := NewOrderedMap[Key, T]()
	m.lock()
	defer m.unlock()
	for _, p := range m.pairs {
		key, value := FromPair(p)
		if predicate(key, value) {
			tmap.Set(key, value)
		} else {
			fmap.Set(key, value)
		}
	}
	return tmap, fmap
}

func Fold[Key comparable, T, State any](folder func(State, Key, T) State, initial State, table OrderedMap[Key, T]) State {
	table.lock()
	defer table.unlock()
	return Fold(func(s State, t Pair[Key, T]) State {
		return folder(s, t.First, t.Second)
	}, initial, table.pairs)
}

func FoldBack[Key comparable, T, State any](folder func(Key, T, State) State, table OrderedMap[Key, T], initial State) State {
	table.lock()
	defer table.unlock()
	return FoldBack(func(t Pair[Key, T], s State) State {
		return folder(t.First, t.Second, s)
	}, table.pairs, initial)
}

func MapTo[Key comparable, T, R any](mapping func(Key, T) R, table OrderedMap[Key, T]) OrderedMap[Key, R] {
	output := NewOrderedMap[Key, R]()
	table.lock()
	defer table.unlock()
	for _, p := range table.pairs {
		key, value := FromPair(p)
		output.Set(key, mapping(key, value))
	}
	return output
}

func Pick[Key comparable, T, R any](chooser func(Key, T) option.Option[R], table OrderedMap[Key, T]) (R, error) {
	for _, p := range table.pairs {
		key, value := FromPair(p)
		v := chooser(key, value)
		if v.IsSome() {
			return v.Value(), nil
		}
	}
	return *(new(R)), KeyNotFoundErr
}

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

func Set[Key comparable, T any](table OrderedMap[Key, T], key Key, value T) OrderedMap[Key, T] {
	m := table.clone()
	m.Set(key, value)
	return m
}

func Remove[Key comparable, T any](table OrderedMap[Key, T], key Key) OrderedMap[Key, T] {
	m := table.clone()
	m.Remove(key)
	return m
}
