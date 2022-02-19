package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleTranspose() {
	r := functional.Transpose([][]int{{10, 20, 30}, {11, 21, 31}})
	fmt.Println(r)
	// Output: [[10 11] [20 21] [30 31]]
}
