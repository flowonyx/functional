package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleMax() {
	max := functional.Max(1, 2, 3, -1)
	fmt.Println(max)
	// Output: 3
}

func ExampleMin() {
	min := functional.Min(1, 2, 3, -1)
	fmt.Println(min)
	// Output: -1
}

func ExampleMaxBy() {
	max := functional.MaxBy(func(s string) int { return len(s) }, []string{"short", "longest"})
	fmt.Println(max)
	// Output: longest
}

func ExampleMinBy() {
	min := functional.MinBy(func(s string) int { return len(s) }, []string{"short", "longest"})
	fmt.Println(min)
	// Output: short
}
