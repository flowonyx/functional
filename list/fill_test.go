package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleFill() {
	input := []int{0, 1, 2, 3, 4, 5}
	if err := list.Fill(input, 3, 2, 100); err != nil {
		panic(err)
	}
	fmt.Println(input)
	// Output: [0 1 2 100 100 5]
}
