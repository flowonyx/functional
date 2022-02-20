package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleIfThenElse() {
	t := 1
	r := list.If(t == 1, func() int { return 1 }).Elif(t <= 2, func() int { return 2 }).Else(func() int { return 3 })
	fmt.Println(r)
	// Output: 1
}

func Example2IfThenElse() {
	t := 2
	r := list.If(t == 1, func() int { return 1 }).Elif(t <= 2, func() int { return 2 }).Else(func() int { return 3 })
	fmt.Println(r)
	// Output: 2
}

func Example3IfThenElse() {
	t := 3
	r := list.If(t == 1, func() int { return 1 }).Elif(t <= 2, func() int { return 2 }).Else(func() int { return 3 })
	fmt.Println(r)
	// Output: 3
}

func ExampleIfV() {
	t := 1
	r := list.IfV(t == 1, 1).Elif(t <= 2, 2).Else(3)
	fmt.Println(r)
	// Output: 1
}

func Example2IfV() {
	t := 2
	r := list.IfV(t == 1, 1).Elif(t <= 2, 2).Else(3)
	fmt.Println(r)
	// Output: 2
}

func Example3IfV() {
	t := 3
	r := list.IfV(t == 1, 1).Elif(t <= 2, 2).Else(3)
	fmt.Println(r)
	// Output: 3
}
