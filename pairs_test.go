package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleAllPairs() {
	input1 := []int{1, 2}
	input2 := []int{3, 4}
	r := functional.AllPairs(input1, input2)
	fmt.Println(r)
	// Output: [{1 3} {1 4} {2 3} {2 4}]
}

func ExamplePairwise() {
	input := []int{1, 2, 3, 4, 5}
	r := functional.Pairwise(input)
	fmt.Println(r)
	// Output: [{1 2} {2 3} {3 4} {4 5}]
}
