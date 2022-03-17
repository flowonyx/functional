package functional_test

import (
	"fmt"

	. "github.com/flowonyx/functional"
)

func ExampleIf() {
	t := 1
	r := If(t == 1, func() int { return 1 }).Elif(t <= 2, func() int { return 2 }).Else(func() int { return 3 })
	fmt.Println(r)
	// Output: 1
}

func ExampleIf_second() {
	t := 2
	r := If(t == 1, func() int { return 1 }).Elif(t <= 2, func() int { return 2 }).Else(func() int { return 3 })
	fmt.Println(r)
	// Output: 2
}

func ExampleIf_third() {
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

func ExampleIfV_second() {
	t := 2
	r := IfV(t == 1, 1).Elif(t <= 2, 2).Else(3)
	fmt.Println(r)
	// Output: 2
}

func ExampleIfV_third() {
	t := 3
	r := IfV(t == 1, 1).Elif(t <= 2, 2).Else(3)
	fmt.Println(r)
	// Output: 3
}
