package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleTranspose() {
	r := list.Transpose([][]int{{10, 20, 30}, {11, 21, 31}})
	fmt.Println(r)
	// Output: [[10 11] [20 21] [30 31]]
}
