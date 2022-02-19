package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleIndexed() {
	input := []string{"a", "b", "c"}
	r := functional.Indexed(input)
	fmt.Println(r)
	// Output: [{0 a} {1 b} {2 c}]
}

func ExampleIndexOf() {
	input := []string{"a", "b", "c"}
	r := functional.IndexOf("b", input)
	fmt.Println(r)
	// Output: 1
}

func Example2IndexOf() {
	input := []string{"a", "b", "c"}
	r := functional.IndexOf("d", input)
	fmt.Println(r)
	// Output: -1
}

func ExampleIndexBy() {
	input := []string{"a", "bb", "ccc"}
	r := functional.IndexBy(func(s string) bool { return len(s) <= 2 }, input)
	fmt.Println(r)
	// Output: 0
}

func ExampleIndexByBack() {
	input := []string{"a", "bb", "ccc"}
	r := functional.IndexByBack(func(s string) bool { return len(s) <= 2 }, input)
	fmt.Println(r)
	// Output: 1
}

func ExampleIndexOfBack() {
	input := []int{1, 2, 3, 1}
	r := functional.IndexOfBack(1, input)
	fmt.Println(r)
	// Output: 3
}
