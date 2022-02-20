package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleForAll() {
	r := list.ForAll(list.LessThan(4), []int{1, 2, 3})
	fmt.Println(r)
	// Output: true
}

func Example2ForAll() {
	r := list.ForAll(list.Not(list.LessThan(4)), []int{1, 2, 3})
	fmt.Println(r)
	// Output: false
}

func ExampleForAll2() {
	r := list.ForAll2(list.Equal2[int], []int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(r)
	// Output: true
}

func Example2ForAll2() {
	r := list.ForAll2(list.Equal2[int], []int{4, 2, 3}, []int{1, 2, 3})
	fmt.Println(r)
	// Output: false
}

func Example3ForAll2() {
	r := list.ForAll2(list.LessThan2[int], []int{1, 2, 3}, []int{2, 3, 4})
	fmt.Println(r)
	// Output: true
}
