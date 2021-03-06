package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleScan() {
	input := []int{1, 2, 3}
	r := list.Scan(func(s int, t int) int {
		return s + t
	}, 0, input)
	fmt.Println(r)
	// Output: [0 1 3 6]
}

func ExampleScanBack() {
	input := []int{1, 2, 3}
	r := list.ScanBack(func(s int, t int) int {
		return s + t
	}, 0, input)
	fmt.Println(r)
	// Output: [0 3 5 6]
}
