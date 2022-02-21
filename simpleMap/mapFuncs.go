package simpleMap

import (
	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/option"
)

func FilterMap[Key comparable, T any](predicate Predicate2[Key, T], table map[Key]T) map[Key]T {
	output := make(map[Key]T)
	for k, v := range table {
		if predicate(k, v) {
			output[k] = v
		}
	}
	return output
}

func TryFindMap[Key comparable, T any](key Key, table map[Key]T) option.Option[T] {
	if v, ok := table[key]; ok {
		return option.Some(v)
	}
	return option.None[T]()
}

func FindKeyMap[Key comparable, T any](predicate Predicate2[Key, T], table map[Key]T) (Key, error) {
	for k, v := range table {
		if predicate(k, v) {
			return k, nil
		}
	}
	return *(new(Key)), errors.KeyNotFoundErr
}

func TryFindKeyMap[Key comparable, T any](predicate Predicate2[Key, T], table map[Key]T) option.Option[Key] {
	if k, err := FindKeyMap(predicate, table); err != nil {
		return option.Some(k)
	}
	return option.None[Key]()
}

func FoldMap[Key comparable, T, State any](folder func(State, Key, T) State, initial State, table map[Key]T) State {
	output := initial
	for k, v := range table {
		output = folder(output, k, v)
	}
	return output
}
