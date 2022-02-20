package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleIndexed() {
	input := []string{"a", "b", "c"}
	r := list.Indexed(input)
	fmt.Println(r)
	// Output: [{0 a} {1 b} {2 c}]
}

func ExampleIndexOf() {
	input := []string{"a", "b", "c"}
	r := list.IndexOf("b", input)
	fmt.Println(r)
	// Output: 1
}

func Example2IndexOf() {
	input := []string{"a", "b", "c"}
	r := list.IndexOf("d", input)
	fmt.Println(r)
	// Output: -1
}

func ExampleIndexBy() {
	input := []string{"a", "bb", "ccc"}
	r := list.IndexBy(func(s string) bool { return len(s) <= 2 }, input)
	fmt.Println(r)
	// Output: 0
}

func ExampleIndexByBack() {
	input := []string{"a", "bb", "ccc"}
	r := list.IndexByBack(func(s string) bool { return len(s) <= 2 }, input)
	fmt.Println(r)
	// Output: 1
}

func ExampleIndexOfBack() {
	input := []int{1, 2, 3, 1}
	r := list.IndexOfBack(1, input)
	fmt.Println(r)
	// Output: 3
}
