package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleWindowed() {
	input := []int{1, 2, 3, 4, 5}
	r := list.Windowed(3, input)
	fmt.Println(r)
	// Output: [[1 2 3] [2 3 4] [3 4 5]]
}
