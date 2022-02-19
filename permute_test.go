package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExamplePermute() {
	input := []int{1, 2, 3, 4}
	r := functional.Permute(func(i int) int { return (i + 1) % 4 }, input)
	fmt.Println(r)
	// Output: [4 1 2 3]
}
