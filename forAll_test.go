package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleForAll() {
	r := functional.ForAll(functional.LessThan(4), []int{1, 2, 3})
	fmt.Println(r)
	// Output: true
}

func Example2ForAll() {
	r := functional.ForAll(functional.Not(functional.LessThan(4)), []int{1, 2, 3})
	fmt.Println(r)
	// Output: false
}

func ExampleForAll2() {
	r := functional.ForAll2(functional.Equal2[int], []int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(r)
	// Output: true
}

func Example2ForAll2() {
	r := functional.ForAll2(functional.Equal2[int], []int{4, 2, 3}, []int{1, 2, 3})
	fmt.Println(r)
	// Output: false
}

func Example3ForAll2() {
	r := functional.ForAll2(functional.LessThan2[int], []int{1, 2, 3}, []int{2, 3, 4})
	fmt.Println(r)
	// Output: true
}
