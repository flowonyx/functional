package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleConcat() {
	inputs := [][]int{{1, 2}, {3}, {4, 5}}
	r := functional.Concat(inputs...)
	fmt.Println(r)
	// Output: [1 2 3 4 5]
}
