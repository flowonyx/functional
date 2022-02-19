package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleCountBy() {
	type foo struct {
		Bar string
	}

	input := []foo{{Bar: "a"}, {Bar: "b"}, {Bar: "a"}}

	c := functional.CountBy(func(f foo) string { return f.Bar }, input)
	fmt.Println(c)
	// Output: map[a:2 b:1]
}
