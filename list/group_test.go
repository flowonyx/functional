package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleGroupBy() {
	r := list.GroupBy(func(i int) int { return i % 2 }, []int{1, 2, 3, 4, 5})
	fmt.Println(r)
	// Output: [(1, [1 3 5]) (0, [2 4])]
}

func ExampleGroupByAsMap() {
	r := list.GroupByAsMap(func(i int) int { return i % 2 }, []int{1, 2, 3, 4, 5})
	fmt.Println(r)
	// Output: map[0:[2 4] 1:[1 3 5]]
}
