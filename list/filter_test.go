package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleFilter() {
	input := []int{0, 1, 2, 3, 4}
	r := list.Filter(func(i int) bool { return i%2 == 0 }, input)
	fmt.Println(r)
	// Output: [0 2 4]
}
