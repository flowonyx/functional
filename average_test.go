package functional_test

import (
	"fmt"

	. "github.com/flowonyx/functional"
)

func ExampleAverage() {
	fmt.Println(Average([]int{1, 2, 3, 4, 5}))
	// Output: 3
}

func ExampleAverageBy() {
	fmt.Println(AverageBy([]uint{1, 2, 3, 4, 5}, func(t uint) float64 {
		return float64(t) * 0.5
	}))
	// Output: 1.5
}
