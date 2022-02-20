package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleInitSlice() {
	r := list.InitSlice(4, func(i int) int { return i + 5 })
	fmt.Println(r)
	// Output: [5 6 7 8]
}

func ExampleInitSlice2D() {
	r := list.InitSlice2D(func(i, j int) int { return i + j }, 2, 3)
	fmt.Println(r)
	// Output: [[0 1 2] [1 2 3]]
}

func ExampleInitSlice3D() {
	r := list.InitSlice3D(func(i, j, k int) int { return 100*i + 10*j + k }, 2, 2, 3)
	fmt.Println(r)
	// Output: [[[0 1 2] [10 11 12]] [[100 101 102] [110 111 112]]]
}

func ExampleInitSlice4D() {
	r := list.InitSlice4D(func(i, j, k, l int) int { return 1000*i + 100*j + 10*k + l }, 2, 2, 2, 2)
	fmt.Println(r)
	// Output: [[[[0 1] [10 11]] [[100 101] [110 111]]] [[[1000 1001] [1010 1011]] [[1100 1101] [1110 1111]]]]
}
