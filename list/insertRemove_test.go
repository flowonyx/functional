package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleInsertAt() {
	input := []int{0, 1, 2, 3}
	r, err := list.InsertAt(2, 4, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: [0 1 4 2 3]
}

func ExampleInsertManyAt() {
	input := []int{0, 1, 2, 3}
	r, err := list.InsertManyAt(2, []int{4, 5, 6}, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: [0 1 4 5 6 2 3]
}

func ExampleRemoveAt() {
	input := []int{0, 1, 2, 3}
	r, err := list.RemoveAt(2, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: [0 1 3]

}

func ExampleRemoveAt_second() {
	input := []int{0, 1, 2, 3}
	r, err := list.RemoveAt(3, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: [0 1 2]

}

func ExampleRemoveManyAt() {
	input := []int{0, 1, 2, 3}
	r, err := list.RemoveManyAt(2, 5, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: [0 1]

}
