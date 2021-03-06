package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleRange() {
	r := list.Range(0, 10)
	fmt.Println(r)
	// Output: [0 1 2 3 4 5 6 7 8 9 10]
}

func ExampleRange_second() {
	r := list.Range(10, 0)
	fmt.Println(r)
	// Output: [10 9 8 7 6 5 4 3 2 1 0]
}

func ExampleRange_third() {
	r := list.Range(10, 0, 2)
	fmt.Println(r)
	// Output: [10 8 6 4 2 0]
}

func ExampleRangeTo() {
	r := list.RangeTo(8)
	fmt.Println(r)
	// Output: [0 1 2 3 4 5 6 7 8]
}

func ExampleRangeChan() {
	r := list.RangeChan(0, 10)
	var s []int
	for i := range r {
		s = append(s, i)
	}
	fmt.Println(s)
	// Output: [0 1 2 3 4 5 6 7 8 9 10]
}

func ExampleDoRange() {
	list.DoRange(func(i int) { fmt.Print(i) }, 2, 5)
	// Output: 2345
}

func ExampleDoRangeTo() {
	list.DoRangeTo(func(i int) { fmt.Print(i) }, 4)
	// Output: 01234
}

func ExampleDoRangeToRev() {
	list.DoRangeToRev(func(i int) { fmt.Print(i) }, 4)
	// Output: 43210
}
