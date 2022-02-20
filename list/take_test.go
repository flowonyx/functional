package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleTake() {
	r := list.Take(3, []int{0, 1, 2, 3, 4})
	fmt.Println(r.Value())
	// Output: [0 1 2]
}

func ExampleTakeWhile() {
	r := list.TakeWhile(func(i int) bool { return i == 0 || i%2 != 0 }, []int{0, 1, 2, 3, 4})
	fmt.Println(r)
	// Output: [0 1]
}

func ExampleTruncate() {
	r := list.Truncate(3, []int{0, 1, 2, 3, 4})
	fmt.Println(r)
	// Output: [0 1 2]
}
