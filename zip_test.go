package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleZip() {
	input1 := []int{1, 2}
	input2 := []string{"one", "two"}

	r := functional.Zip(input1, input2)
	fmt.Println(r)
	// Output: [{1 one} {2 two}]
}

func ExampleZip3() {
	input1 := []int{1, 2}
	input2 := []string{"one", "two"}
	input3 := []string{"I", "II"}

	r := functional.Zip3(input1, input2, input3)
	fmt.Println(r)
	// Output: [{1 one I} {2 two II}]
}

func ExampleUnzip() {
	input := []functional.Pair[int, string]{{1, "one"}, {2, "two"}}
	a, b := functional.Unzip(input)
	fmt.Println(a, b)
	// Output: [1 2] [one two]
}

func ExampleUnzip3() {
	input := []functional.Triple[int, string, string]{{1, "one", "I"}, {2, "two", "II"}}
	a, b, c := functional.Unzip3(input)
	fmt.Println(a, b, c)
	// Output: [1 2] [one two] [I II]
}
