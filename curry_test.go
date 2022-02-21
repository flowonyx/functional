package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleCurry() {
	f := func(i int) int { return i * 2 }
	f2 := functional.Curry(f, 2)
	r := f2()
	fmt.Println(r)
	// Output: 4
}

func ExampleCurry2() {
	f := func(mul, i int) int { return i * mul }
	f2 := functional.Curry2(f, 2, 2)
	r := f2()
	fmt.Println(r)
	// Output: 4
}

func ExampleCurry2To1() {
	f := func(mul, i int) int { return i * mul }
	f2 := functional.Curry2To1(f, 2)
	r := f2(2)
	fmt.Println(r)
	// Output: 4
}
