package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExamplePartition() {
	input := []int{1, 2, 3, 4}
	t, f := functional.Partition(func(i int) bool { return i%2 == 0 }, input)
	fmt.Println(t, f)
	// Output: [2 4] [1 3]
}
