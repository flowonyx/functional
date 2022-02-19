package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleSkip() {
	input := []int{1, 2, 3, 4, 5}
	r := functional.Skip(2, input)
	fmt.Println(r)
	// Output: [3 4 5]
}

func ExampleSkipWhile() {
	input := []int{1, 2, 3, 4, 5}
	r := functional.SkipWhile(functional.LessThan(3), input)
	fmt.Println(r)
	// Output: [3 4 5]
}
