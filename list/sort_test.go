package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleSort() {
	input := []int{3, 2, 5, 8, 1}
	r := list.Sort(input)
	fmt.Println(r)
	// Output: [1 2 3 5 8]
}

func ExampleSortDescending() {
	input := []int{3, 2, 5, 8, 1}
	r := list.SortDescending(input)
	fmt.Println(r)
	// Output: [8 5 3 2 1]
}

func ExampleSortBy() {
	input := []int{3, 2, 5, 8, 1}
	r := list.SortBy(func(i int) int {
		return list.IfV(i < 5, i*10).Else(i)
	}, input)
	fmt.Println(r)
	// Output: [5 8 1 2 3]
}

func ExampleSortByDescending() {
	input := []int{3, 2, 5, 8, 1}
	r := list.SortByDescending(func(i int) int {
		return list.IfV(i < 5, i*10).Else(i)
	}, input)
	fmt.Println(r)
	// Output: [3 2 1 8 5]
}

func ExampleSortWith() {
	input := []int{3, 2, 5, 8, 1}
	r := list.SortWith(func(t1, t2 int) bool {
		return t1 > t2
	}, input)
	fmt.Println(r)
	// Output: [8 5 3 2 1]
}

func ExampleReverse() {
	input := []int{3, 2, 5, 8, 1}
	r := list.Reverse(input)
	fmt.Println(r)
	// Output: [1 8 5 2 3]
}
