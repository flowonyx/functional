package list_test

import (
	"fmt"
	"strconv"

	"github.com/flowonyx/functional/list"
)

func ExampleDistinct() {
	d := list.Distinct([]int{1, 2, 3, 2, 3, 4})
	fmt.Println(d)
	// Output: [1 2 3 4]
}

func ExampleDistinctBy() {
	d := list.DistinctBy(func(i int) string { return "s:" + strconv.Itoa(i) }, []int{1, 2, 3, 2, 3, 4})
	fmt.Println(d)
	// Output: [1 2 3 4]
}
