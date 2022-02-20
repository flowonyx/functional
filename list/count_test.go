package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleCountBy() {
	type foo struct {
		Bar string
	}

	input := []foo{{Bar: "a"}, {Bar: "b"}, {Bar: "a"}}

	c := list.CountBy(func(f foo) string { return f.Bar }, input)
	fmt.Println(c)
	// Output: map[a:2 b:1]
}
