package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleMap() {
	input := []string{"a", "bb", "ccc"}
	r := functional.Map(func(s string) int { return len(s) }, input)
	fmt.Println(r)
	// Output: [1 2 3]
}

func ExampleMap2() {
	input1 := []string{"a", "bb", "ccc"}
	input2 := []string{"dddd", "eeeee", "ffffff"}
	r := functional.Map2(func(a, b string) int { return len(a) + len(b) }, input1, input2)
	fmt.Println(r)
	// Output: [5 7 9]
}

func ExampleMap3() {
	input1 := []string{"a", "bb", "ccc"}
	input2 := []string{"dddd", "eeeee", "ffffff"}
	input3 := []string{"A", "B", "C"}
	r := functional.Map3(func(a, b, c string) string { return fmt.Sprintf("%s: %d", c, len(a)+len(b)) }, input1, input2, input3)
	fmt.Println(r)
	// Output: [A: 5 B: 7 C: 9]
}

func ExampleMapi() {
	input := []string{"a", "bb", "ccc"}
	r := functional.Mapi(func(i int, s string) int { return i * len(s) }, input)
	fmt.Println(r)
	// Output: [0 2 6]
}

func ExampleMapi2() {
	input1 := []string{"a", "bb", "ccc"}
	input2 := []string{"dddd", "eeeee", "ffffff"}
	r := functional.Mapi2(func(i int, a, b string) int { return i * (len(a) + len(b)) }, input1, input2)
	fmt.Println(r)
	// Output: [0 7 18]
}

func ExampleMapi3() {
	input1 := []string{"a", "bb", "ccc"}
	input2 := []string{"dddd", "eeeee", "ffffff"}
	input3 := []string{"A", "B", "C"}
	r := functional.Mapi3(func(i int, a, b, c string) string { return fmt.Sprintf("%d-%s: %d", i, c, len(a)+len(b)) }, input1, input2, input3)
	fmt.Println(r)
	// Output: [0-A: 5 1-B: 7 2-C: 9]
}

func ExampleMap2D() {
	input := [][]string{{"a", "bb"}, {"ccc"}}
	r := functional.Map2D(func(s string) int { return len(s) }, input)
	fmt.Println(r)
	// Output: [[1 2] [3]]
}

func ExampleMap3D() {
	input := [][][]string{{{"a", "bb"}, {"ccc"}}, {{"dddd"}}}
	r := functional.Map3D(func(s string) int { return len(s) }, input)
	fmt.Println(r)
	// Output: [[[1 2] [3]] [[4]]]
}

func ExampleMapi2D() {
	input := [][]string{{"a", "bb"}, {"ccc"}}
	r := functional.Mapi2D(func(i, j int, s string) int { return i + j + len(s) }, input)
	fmt.Println(r)
	// Output: [[1 3] [4]]
}

func ExampleMapi3D() {
	input := [][][]string{{{"a", "bb"}, {"ccc"}}, {{"dddd"}}}
	r := functional.Mapi3D(func(i, j, k int, s string) int { return i + j + k + len(s) }, input)
	fmt.Println(r)
	// Output: [[[1 3] [4]] [[5]]]
}
