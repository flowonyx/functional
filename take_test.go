package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleTake() {
	r := functional.Take(3, []int{0, 1, 2, 3, 4})
	fmt.Println(r.Value())
	// Output: [0 1 2]
}

func ExampleTakeWhile() {
	r := functional.TakeWhile(func(i int) bool { return i == 0 || i%2 != 0 }, []int{0, 1, 2, 3, 4})
	fmt.Println(r)
	// Output: [0 1]
}

func ExampleTruncate() {
	r := functional.Truncate(3, []int{0, 1, 2, 3, 4})
	fmt.Println(r)
	// Output: [0 1 2]
}
