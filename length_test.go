package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleLen2() {
	input := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3, 4}, {1, 2, 3}}
	r := functional.Len2(input)
	fmt.Println(r)
	// Output: 3
}

func ExampleLen3() {
	input := [][][]int{{{1, 2, 3}, {1, 2, 3}}, {{1, 2, 3, 4}}}
	r := functional.Len3(input)
	fmt.Println(r)
	// Output: 3
}

func ExampleLen4() {
	input := [][][][]int{{{{1, 2, 3}, {1, 2, 3}}}, {{{1, 2, 3, 4}}}}
	r := functional.Len4(input)
	fmt.Println(r)
	// Output: 3
}
