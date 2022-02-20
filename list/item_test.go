package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleItem() {
	input := []int{0, 1, 2, 3, 4}
	r, err := list.Item(2, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: 2
}

func ExampleItem2D() {
	input := [][]int{{0, 1}, {2, 3, 4}}
	r, err := list.Item2D(input, 1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: 4
}

func ExampleItem3D() {
	input := [][][]int{{}, {{0, 1}, {2, 3, 4}}}
	r, err := list.Item3D(input, 1, 1, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: 3
}

func ExampleItem4D() {
	input := [][][][]int{{}, {{}, {{}, {0, 1}, {2, 3, 4}}}}
	r, err := list.Item4D(input, 1, 1, 2, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: 3
}

func ExampleTryItem() {
	input := []int{0, 1, 2, 3, 4}
	r := list.TryItem(2, input)
	fmt.Println(r.Value())
	// Output: 2
}

func Example2TryItem() {
	input := []int{0, 1, 2, 3, 4}
	r := list.TryItem(5, input)
	fmt.Println(r.IsNone())
	// Output: true
}
