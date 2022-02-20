package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleCreate() {
	a := list.Create(4, "a")
	fmt.Println(a)
	// Output: [a a a a]
}

func ExampleCreate2D() {
	a := list.Create2D(4, 4, "a")
	fmt.Println(a)
	// Output: [[a a a a] [a a a a] [a a a a] [a a a a]]
}

func ExampleCreate3D() {
	a := list.Create3D(2, 2, 2, "a")
	fmt.Println(a)
	// Output: [[[a a] [a a]] [[a a] [a a]]]
}

func ExampleCreate4D() {
	a := list.Create4D(2, 2, 2, 2, "a")
	fmt.Println(a)
	// Output: [[[[a a] [a a]] [[a a] [a a]]] [[[a a] [a a]] [[a a] [a a]]]]
}

func ExampleZeroCreate() {
	a := list.ZeroCreate[int](4)
	fmt.Println(a)
	// Output: [0 0 0 0]
}

func ExampleZeroCreate2D() {
	a := list.ZeroCreate2D[int](4, 4)
	fmt.Println(a)
	// Output: [[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]
}

func ExampleZeroCreate3D() {
	a := list.ZeroCreate3D[int](1, 1, 2)
	fmt.Println(a)
	// Output: [[[0 0]]]
}

func ExampleZeroCreate4D() {
	a := list.ZeroCreate4D[int](1, 1, 2, 2)
	fmt.Println(a)
	// Output: [[[[0 0] [0 0]]]]
}
