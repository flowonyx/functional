package functional_test

import (
	"fmt"

	. "github.com/flowonyx/functional"
)

func ExampleIfThenElse() {
	t := 1
	r := If(t == 1, func() int { return 1 }).Elif(t <= 2, func() int { return 2 }).Else(func() int { return 3 })
	fmt.Println(r)
	// Output: 1
}

func Example2IfThenElse() {
	t := 2
	r := If(t == 1, func() int { return 1 }).Elif(t <= 2, func() int { return 2 }).Else(func() int { return 3 })
	fmt.Println(r)
	// Output: 2
}

func Example3IfThenElse() {
	t := 3
	r := If(t == 1, func() int { return 1 }).Elif(t <= 2, func() int { return 2 }).Else(func() int { return 3 })
	fmt.Println(r)
	// Output: 3
}

func ExampleIfV() {
	t := 1
	r := IfV(t == 1, 1).Elif(t <= 2, 2).Else(3)
	fmt.Println(r)
	// Output: 1
}

func Example2IfV() {
	t := 2
	r := IfV(t == 1, 1).Elif(t <= 2, 2).Else(3)
	fmt.Println(r)
	// Output: 2
}

func Example3IfV() {
	t := 3
	r := IfV(t == 1, 1).Elif(t <= 2, 2).Else(3)
	fmt.Println(r)
	// Output: 3
}
