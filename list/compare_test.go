package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleMax() {
	max := list.Max(1, 2, 3, -1)
	fmt.Println(max)
	// Output: 3
}

func ExampleMin() {
	min := list.Min(1, 2, 3, -1)
	fmt.Println(min)
	// Output: -1
}

func ExampleMaxBy() {
	max := list.MaxBy(func(s string) int { return len(s) }, []string{"short", "longest"})
	fmt.Println(max)
	// Output: longest
}

func ExampleMinBy() {
	min := list.MinBy(func(s string) int { return len(s) }, []string{"short", "longest"})
	fmt.Println(min)
	// Output: short
}
