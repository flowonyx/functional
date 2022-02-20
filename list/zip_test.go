package list_test

import (
	"fmt"

	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/list"
)

func ExampleZip() {
	input1 := []int{1, 2}
	input2 := []string{"one", "two"}

	r := list.Zip(input1, input2)
	fmt.Println(r)
	// Output: [{1 one} {2 two}]
}

func ExampleZip3() {
	input1 := []int{1, 2}
	input2 := []string{"one", "two"}
	input3 := []string{"I", "II"}

	r := list.Zip3(input1, input2, input3)
	fmt.Println(r)
	// Output: [{1 one I} {2 two II}]
}

func ExampleUnzip() {
	input := []Pair[int, string]{{1, "one"}, {2, "two"}}
	a, b := list.Unzip(input)
	fmt.Println(a, b)
	// Output: [1 2] [one two]
}

func ExampleUnzip3() {
	input := []Triple[int, string, string]{{1, "one", "I"}, {2, "two", "II"}}
	a, b, c := list.Unzip3(input)
	fmt.Println(a, b, c)
	// Output: [1 2] [one two] [I II]
}
