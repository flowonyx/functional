package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleExcept() {
	original := []int{1, 2, 3, 4, 5}
	itemsToExclude := []int{1, 3, 5}
	e := functional.Except(itemsToExclude, original)
	fmt.Println(e)
	// Output: [2 4]
}
