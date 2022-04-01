// functional is a package to assist with functional style programming in go.
// At least go 1.18 is required as generics are heavily used throughout.
// Most functionality is found in the packages under this one. This top level package
// contains some generally applicable types and functions that are used in the sub packages.
package functional

import (
	"fmt"
	"strconv"
)

// Pair is a tuple with 2 items.
type Pair[T any, T2 any] struct {
	First  T
	Second T2
}

// PairOf is a simple method of creating a Pair.
func PairOf[T, T2 any](a T, b T2) Pair[T, T2] {
	return Pair[T, T2]{First: a, Second: b}
}

// FromPair pulls the items out of a Pair without needing to
// refer to the First and Second fields.
func FromPair[T1, T2 any](p Pair[T1, T2]) (T1, T2) {
	return p.First, p.Second
}

// Triple is a tuple with 3 items.
type Triple[T, T2, T3 any] struct {
	First  T
	Second T2
	Third  T3
}

// TripleOf is a simple method of creating a Triple.
func TripleOf[T, T2, T3 any](a T, b T2, c T3) Triple[T, T2, T3] {
	return Triple[T, T2, T3]{a, b, c}
}

// FromTriple pulls the items out of a Triple without needing to
// refer to the First, Second, and Third fields.
func FromTriple[T1, T2, T3 any](t Triple[T1, T2, T3]) (a T1, b T2, c T3) {
	return t.First, t.Second, t.Third
}

func (p Pair[T, T2]) String() string {
	var f, s any = p.First, p.Second
	format := func(item any) string {
		if s, ok := item.(fmt.Stringer); ok {
			return s.String()
		}
		switch r := item.(type) {
		case rune:
			return strconv.QuoteRune(r)
		case string:
			return strconv.Quote(r)
		default:
			return fmt.Sprint(r)
		}
	}

	first, second := format(f), format(s)

	return fmt.Sprintf("(%s, %s)", first, second)
}

func (p Triple[T, T2, T3]) String() string {
	var f, s, t any = p.First, p.Second, p.Third
	format := func(item any) string {
		switch r := item.(type) {
		case rune:
			return strconv.QuoteRune(r)
		case string:
			return strconv.Quote(r)
		default:
			return fmt.Sprint(r)
		}
	}

	first, second, third := format(f), format(s), format(t)

	return fmt.Sprintf("(%s, %s, %s)", first, second, third)
}
