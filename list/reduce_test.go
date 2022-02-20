package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleReduce() {
	input := []int{1, 3, 4, 2}
	r := list.Reduce(0, func(a, b int) int { return a*10 + b }, input)
	fmt.Println(r)
	// Output: 1342
}

func ExampleReduceBack() {
	input := []int{1, 3, 4, 2}
	r := list.ReduceBack(0, func(a, b int) int { return a*10 + b }, input)
	fmt.Println(r)
	// Output: 2431
}
