package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleInsertAt() {
	input := []int{0, 1, 2, 3}
	r, err := functional.InsertAt(2, 4, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: [0 1 4 2 3]
}

func ExampleInsertManyAt() {
	input := []int{0, 1, 2, 3}
	r, err := functional.InsertManyAt(2, []int{4, 5, 6}, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: [0 1 4 5 6 2 3]
}

func ExampleRemoveAt() {
	input := []int{0, 1, 2, 3}
	r, err := functional.RemoveAt(2, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: [0 1 3]

}

func ExampleRemoveManyAt() {
	input := []int{0, 1, 2, 3}
	r, err := functional.RemoveManyAt(2, 2, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: [0 1]

}
