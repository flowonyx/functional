package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleConcat() {
	inputs := [][]int{{1, 2}, {3}, {4, 5}}
	r := list.Concat(inputs...)
	fmt.Println(r)
	// Output: [1 2 3 4 5]
}
