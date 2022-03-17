package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExamplePairOf() {
	p := functional.PairOf(1, "2")
	fmt.Println(p.String())
	// Output: (1, "2")
}

func ExampleTripleOf() {
	t := functional.TripleOf(1, "2", 3)
	fmt.Println(t.String())
	// Output: (1, "2", 3)
}

func ExampleFromPair() {
	p := functional.PairOf(1, "2")
	f, s := functional.FromPair(p)
	fmt.Println(f, s)
	// Output: 1 2
}

func ExampleFromTriple() {
	p := functional.TripleOf(1, "2", 3)
	f, s, t := functional.FromTriple(p)
	fmt.Println(f, s, t)
	// Output: 1 2 3
}
