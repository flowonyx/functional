package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleSplitAt() {
	input := []int{0, 1, 2, 3, 4, 5}
	a, b, err := list.SplitAt(3, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(a, b)
	// Output: [0 1 2] [3 4 5]
}

func ExampleSplitInto() {
	input1 := []int{1}
	input2 := []int{1, 2}
	input3 := []int{1, 2, 3}
	input4 := []int{1, 2, 3, 4}
	input5 := []int{1, 2, 3, 4, 5}
	r := list.SplitInto(3, input1)
	r2 := list.SplitInto(3, input2)
	r3 := list.SplitInto(3, input3)
	r4 := list.SplitInto(3, input4)
	r5 := list.SplitInto(3, input5)
	fmt.Println(r, r2, r3, r4, r5)
	// Output: [[1]] [[1] [2]] [[1] [2] [3]] [[1] [2] [3 4]] [[1 2] [3 4] [5]]
}
