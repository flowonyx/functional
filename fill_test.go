package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleFill() {
	input := []int{0, 1, 2, 3, 4, 5}
	functional.Fill(input, 3, 2, 100)
	fmt.Println(input)
	// Output: [0 1 2 100 100 5]
}
