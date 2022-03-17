package functional_test

import (
	"errors"
	"fmt"
	"strings"

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

func ExampleCurry2To1F() {
	f := func(mul, i int) int { return i * mul }
	f2 := functional.Curry2To1F(f) // func(mul int) func(int) int { return func(i int) int { return i * mul }}
	f1 := f2(2)                    // func(i int) int { return i * 2 }
	r := f1(3)
	fmt.Println(r)
	// Output: 6
}

func ExampleCurry3() {
	f := func(a, b, c int) int { return a + b + c }
	f0 := functional.Curry3(f, 1, 2, 3)
	r := f0()
	fmt.Println(r)
	// Output: 6
}

func ExampleCurry3To2() {
	f := func(a, b, c int) int { return a + b + c }
	f2 := functional.Curry3To2(f, 1)
	r := f2(2, 3)
	fmt.Println(r)
	// Output: 6
}

func ExampleCurry3To1() {
	f := func(a, b, c int) int { return a + b + c }
	f1 := functional.Curry3To1(f, 1, 2)
	r := f1(3)
	fmt.Println(r)
	// Output: 6
}

func ExampleSwapParams0() {
	f := func(i int, s string) { fmt.Print(i, s) }
	fs := functional.SwapParams0(f)
	fs("string", 1)
	// Output: 1string
}

func ExampleSwapParams1() {
	fs := functional.SwapParams1(strings.Repeat)
	fc := functional.Curry2To1(fs, 5)
	fmt.Print(fc("*"), fc("-"))
	// Output: *****-----
}

func ExampleSwapParams2() {
	f := func(a, b int) (int, error) {
		if a < b {
			return -1, errors.New("a < b")
		}
		return b, nil
	}
	fs := functional.SwapParams2(f)
	_, err := fs(2, 1)
	if err == nil {
		panic("should have errored")
	}
	i, err := fs(1, 2)
	if err != nil {
		panic(err)
	}

	fmt.Print(i)
	// Output: 1
}
