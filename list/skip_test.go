package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleSkip() {
	input := []int{1, 2, 3, 4, 5}
	r := list.Skip(2, input)
	fmt.Println(r)
	// Output: [3 4 5]
}

func ExampleSkipWhile() {
	input := []int{1, 2, 3, 4, 5}
	r := list.SkipWhile(list.LessThan(3), input)
	fmt.Println(r)
	// Output: [3 4 5]
}
