package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleMax() {
	max, err := list.Max(1, 2, 3, -1)
	if err != nil {
		panic(err)
	}
	fmt.Println(max)
	// Output: 3
}

func ExampleMin() {
	min, err := list.Min(1, 2, 3, -1)
	if err != nil {
		panic(err)
	}
	fmt.Println(min)
	// Output: -1
}

func ExampleMaxBy() {
	max, err := list.MaxBy(func(s string) int { return len(s) }, "short", "longest")
	if err != nil {
		panic(err)
	}
	fmt.Println(max)
	// Output: longest
}

func ExampleMinBy() {
	min, err := list.MinBy(func(s string) int { return len(s) }, "short", "longest")
	if err != nil {
		panic(err)
	}
	fmt.Println(min)
	// Output: short
}
