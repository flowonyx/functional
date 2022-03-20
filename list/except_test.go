package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleExcept() {
	original := []int{1, 2, 3, 4, 5}
	itemsToExclude := []int{1, 3, 5}
	e := list.Except(itemsToExclude, original...)
	fmt.Println(e)
	// Output: [2 4]
}
