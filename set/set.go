package set

import (
	"github.com/flowonyx/functional"
	"github.com/flowonyx/functional/orderedMap"
	"golang.org/x/exp/constraints"
)

type Set[T comparable] struct {
	m orderedMap.OrderedMap[T, struct{}]
}

func NewSet[T comparable](lessFunc ...func(T, T) bool) Set[T] {
	return Set[T]{m: orderedMap.NewOrderedMap[T, struct{}](lessFunc...)}
}

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

func FromSlice[T comparable](input []T) Set[T] {
	s := NewSet[T]()
	for _, i := range input {
		s.Add(i)
	}
	return s
}

func (s Set[T]) ToSlice() []T {
	return s.Items()
}

func (s Set[T]) Add(item T) {
	s.m.Set(item, struct{}{})
}

func (s Set[T]) Remove(item T) {
	s.m.Remove(item)
}

func (s Set[T]) Equal(s2 Set[T]) bool {
	return orderedMap.Equal(s.m, s2.m)
}

func (s Set[T]) Contains(item T) bool {
	return s.m.Contains(item)
}

func (s Set[T]) Exists(predicate functional.Predicate[T]) bool {
	return s.m.Exists(func(item T, _ struct{}) bool { return predicate(item) })
}

func (s Set[T]) Count() int {
	return s.m.Len()
}

func (s Set[T]) IsEmpty() bool {
	return s.Count() == 0
}

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

func (s Set[T]) IsProperSubsetOf(potentialSuperset Set[T]) bool {
	if potentialSuperset.Count() == 0 || s.Count() == 0 {
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

func (s Set[T]) IsProperSupersetOf(potentialSubset Set[T]) bool {
	return potentialSubset.IsProperSubsetOf(s)
}

func (s Set[T]) IsSupersetOf(potentialSubset Set[T]) bool {
	return potentialSubset.IsSubsetOf(s)
}

func (s Set[T]) IndexOf(item T) int {
	index := -1
	s.m.Iteri(func(i int, t T, _ struct{}) {
		if t == item {
			index = i
		}
	})
	return index
}

func (s Set[T]) Items() []T {
	return s.m.Keys()
}

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

func (s Set[T]) Filter(predicate functional.Predicate[T]) Set[T] {
	output := NewSet[T]()
	for _, i := range s.Items() {
		if predicate(i) {
			output.Add(i)
		}
	}
	return output
}

func (s Set[T]) ForAll(predicate functional.Predicate[T]) bool {
	return s.m.ForAll(func(k T, _ struct{}) bool { return predicate(k) })
}

func (s Set[T]) Iter(action func(T)) {
	for _, i := range s.Items() {
		action(i)
	}
}

func (s Set[T]) Iteri(action func(int, T)) {
	for i, v := range s.Items() {
		action(i, v)
	}
}

func (s Set[T]) Partition(predicate functional.Predicate[T]) (Set[T], Set[T]) {
	tset := NewSet[T]()
	fset := NewSet[T]()
	s.Iter(func(t T) {
		if predicate(t) {
			tset.Add(t)
		} else {
			fset.Add(t)
		}
	})
	return tset, fset
}

func Fold[T comparable, State any](folder func(State, T) State, initial State, s Set[T]) State {
	return orderedMap.Fold(func(s State, k T, _ struct{}) State {
		return folder(s, k)
	}, initial, s.m)
}

func FoldBack[T comparable, State any](folder func(T, State) State, s Set[T], initial State) State {
	return orderedMap.FoldBack(func(k T, _ struct{}, s State) State {
		return folder(k, s)
	}, s.m, initial)
}

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

func Map[T comparable, R comparable](mapping functional.Projection[T, R], s Set[T]) Set[R] {
	output := NewSet[R]()
	s.Iter(func(t T) { output.Add(mapping(t)) })
	return output
}

func MaxElement[T constraints.Ordered](s Set[T]) T {
	return functional.Max(s.Items()...)
}

func MinElement[T constraints.Ordered](s Set[T]) T {
	return functional.Min(s.Items()...)
}

func MaxElementBy[T comparable, R constraints.Ordered](projection functional.Projection[T, R], s Set[T]) T {
	return functional.MaxBy(projection, s.Items())
}

func MinElementBy[T comparable, R constraints.Ordered](projection functional.Projection[T, R], s Set[T]) T {
	return functional.MinBy(projection, s.Items())
}
