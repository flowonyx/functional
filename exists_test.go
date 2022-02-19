package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleExists() {
	e := functional.Exists(func(i int) bool { return i%4 == 0 }, []int{1, 2, 3, 4, 5})
	fmt.Println(e)
	// Output: true
}

func Example2Exists() {
	e := functional.Exists(func(i int) bool { return i%6 == 0 }, []int{1, 2, 3, 4, 5})
	fmt.Println(e)
	// Output: false
}

func ExampleExists2() {
	input1 := []int{1, 2}
	input2 := []int{1, 2, 0}
	e := functional.Exists2(func(a, b int) bool { return a > b }, input1, input2)
	fmt.Println(e)
	// Output: false
}

func Example2Exists2() {
	input1 := []int{1, 4}
	input2 := []int{1, 3, 5}
	e := functional.Exists2(func(a, b int) bool { return a > b }, input1, input2)
	fmt.Println(e)
	// Output: true
}
