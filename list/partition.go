package list

import (
	"golang.org/x/exp/slices"
)

func Partition[T any](predicate func(T) bool, input []T) ([]T, []T) {
	tlist := Empty[T](len(input))
	flist := Empty[T](len(input))
	for i := range input {
		if predicate(input[i]) {
			tlist = append(tlist, input[i])
		} else {
			flist = append(flist, input[i])
		}
	}
	slices.Clip(tlist)
	slices.Clip(flist)
	return tlist, flist
}
