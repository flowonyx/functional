package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleCreate() {
	a := functional.Create(4, "a")
	fmt.Println(a)
	// Output: [a a a a]
}

func ExampleCreate2D() {
	a := functional.Create2D(4, 4, "a")
	fmt.Println(a)
	// Output: [[a a a a] [a a a a] [a a a a] [a a a a]]
}

func ExampleCreate3D() {
	a := functional.Create3D(2, 2, 2, "a")
	fmt.Println(a)
	// Output: [[[a a] [a a]] [[a a] [a a]]]
}

func ExampleCreate4D() {
	a := functional.Create4D(2, 2, 2, 2, "a")
	fmt.Println(a)
	// Output: [[[[a a] [a a]] [[a a] [a a]]] [[[a a] [a a]] [[a a] [a a]]]]
}

func ExampleZeroCreate() {
	a := functional.ZeroCreate[int](4)
	fmt.Println(a)
	// Output: [0 0 0 0]
}

func ExampleZeroCreate2D() {
	a := functional.ZeroCreate2D[int](4, 4)
	fmt.Println(a)
	// Output: [[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]
}

func ExampleZeroCreate3D() {
	a := functional.ZeroCreate3D[int](1, 1, 2)
	fmt.Println(a)
	// Output: [[[0 0]]]
}

func ExampleZeroCreate4D() {
	a := functional.ZeroCreate4D[int](1, 1, 2, 2)
	fmt.Println(a)
	// Output: [[[[0 0] [0 0]]]]
}
