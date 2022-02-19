package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleRange_1() {
	r := functional.Range(0, 10)
	fmt.Println(r)
	// Output: [0 1 2 3 4 5 6 7 8 9 10]
}

func ExampleRange_2() {
	r := functional.Range(10, 0)
	fmt.Println(r)
	// Output: [10 9 8 7 6 5 4 3 2 1 0]
}

func ExampleRange_3() {
	r := functional.Range(10, 0, 2)
	fmt.Println(r)
	// Output: [10 8 6 4 2 0]
}

func ExampleRangeTo() {
	r := functional.RangeTo(8)
	fmt.Println(r)
	// Output: [0 1 2 3 4 5 6 7 8]
}

func ExampleRangeSeq() {
	r := functional.RangeChan(0, 10)
	var s []int
	for i := range r {
		s = append(s, i)
	}
	fmt.Println(s)
	// Output: [0 1 2 3 4 5 6 7 8 9 10]
}

func ExampleDoRange() {
	functional.DoRange(func(i int) { fmt.Print(i) }, 2, 5)
	// Output: 2345
}

func ExampleDoRangeTo() {
	functional.DoRangeTo(func(i int) { fmt.Print(i) }, 4)
	// Output: 01234
}

func ExampleDoRangeToRev() {
	functional.DoRangeToRev(func(i int) { fmt.Print(i) }, 4)
	// Output: 43210
}
