package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleHead() {
	h, err := list.Head([]int{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(h)
	// Output: 1
}

func ExampleTryHead() {
	h := list.TryHead([]int{1, 2, 3})
	fmt.Println(h.Value())
	// Output: 1
}

func ExampleTail() {
	t := list.Tail([]int{1, 2, 3})
	fmt.Println(t)
	// Output: [2 3]
}

func ExampleLast() {
	l, err := list.Last([]int{1, 2, 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(l)
	// Output: 3
}

func ExampleTryLast() {
	l := list.TryLast([]int{1, 2, 3})
	fmt.Println(l.Value())
	// Output: 3
}
