package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleExists() {
	e := list.Exists(func(i int) bool { return i%4 == 0 }, 1, 2, 3, 4, 5)
	fmt.Println(e)
	// Output: true
}

func ExampleExists_second() {
	e := list.Exists(func(i int) bool { return i%6 == 0 }, 1, 2, 3, 4, 5)
	fmt.Println(e)
	// Output: false
}

func ExampleExists2() {
	input1 := []int{1, 2}
	input2 := []int{1, 2, 0}
	e := list.Exists2(func(a, b int) bool { return a > b }, input1, input2)
	fmt.Println(e)
	// Output: false
}

func ExampleExists2_second() {
	input1 := []int{1, 4}
	input2 := []int{1, 3, 5}
	e := list.Exists2(func(a, b int) bool { return a > b }, input1, input2)
	fmt.Println(e)
	// Output: true
}
