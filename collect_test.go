package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleCollect() {
	type foo struct {
		Bar []int
	}
	input := []foo{
		{Bar: []int{1, 2}},
		{Bar: []int{3, 4}},
	}
	r := functional.Collect(func(f foo) []int { return f.Bar }, input)
	fmt.Println(r)
	// Output: [1 2 3 4]
}
