package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleSplitAt() {
	input := []int{0, 1, 2, 3, 4, 5}
	a, b := functional.SplitAt(3, input)
	fmt.Println(a, b)
	// Output: [0 1 2] [3 4 5]
}

func ExampleSplitInto() {
	input := []int{1, 2, 3, 4, 5}
	r := functional.SplitInto(3, input)
	fmt.Println(r)
	// Output: [[1 2] [3 4] [5]]
}
