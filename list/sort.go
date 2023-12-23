package list

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Sort returns a clone of a slice of any ordered type sorted in ascending order.
func Sort[T constraints.Ordered](values []T) []T {
	c := slices.Clone(values)
	slices.Sort(c)
	return c
}

// SortDescending returns a clone of a slice of any ordered type sorted in descending order.
func SortDescending[T constraints.Ordered](values []T) []T {
	c := slices.Clone(values)
	slices.SortFunc(c, cmpDesc)
	return c
}

// SortBy returns a clone of a slice of any type sorted in ascending order based on the key returned from projection.
func SortBy[T any, Key constraints.Ordered](projection func(T) Key, values []T) []T {
	return SortWith(func(t1, t2 T) int {
		p1 := projection(t1)
		p2 := projection(t2)
		return cmp(p1, p2)
	}, values)
}

// SortByDescending returns a clone of a slice of any type sorted in descending order based on the key returned from projection.
func SortByDescending[T any, Key constraints.Ordered](projection func(T) Key, values []T) []T {
	return SortWith(func(t1, t2 T) int {
		p1 := projection(t1)
		p2 := projection(t2)
		return cmpDesc(p1, p2)
	}, values)

}

// SortWith returns a clone of a slice of any type sorted in order as determined by the less function. This sort is not guaranteed to be stable.
func SortWith[T any](less func(T, T) int, values []T) []T {
	output := slices.Clone(values)
	slices.SortFunc(output, less)
	return output
}

// Reverse returns a clone of a slice with values in reverse order.
func Reverse[T any](values []T) []T {
	output := make([]T, len(values))
	j := LastIndexOf(output)
	for i := range values {
		output[j] = values[i]
		j--
	}
	return output
}

func cmp[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if b < a {
		return 1
	}
	return 0
}

func cmpDesc[T constraints.Ordered](a, b T) int {
	if a > b {
		return -1
	}
	if b > a {
		return 1
	}
	return 0
}
