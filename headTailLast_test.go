package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleHead() {
	h, err := functional.Head([]int{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(h)
	// Output: 1
}

func ExampleTryHead() {
	h := functional.TryHead([]int{1, 2, 3})
	fmt.Println(h.Value())
	// Output: 1
}

func ExampleTail() {
	t := functional.Tail([]int{1, 2, 3})
	fmt.Println(t)
	// Output: [2 3]
}

func ExampleLast() {
	l, err := functional.Last([]int{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(l)
	// Output: 3
}

func ExampleTryLast() {
	l := functional.TryLast([]int{1, 2, 3})
	fmt.Println(l.Value())
	// Output: 3
}
