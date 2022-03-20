package list_test

import (
	"fmt"

	. "github.com/flowonyx/functional/list"
)

func ExampleAverage() {
	fmt.Println(Average(1, 2, 3, 4, 5))
	// Output: 3
}

func ExampleAverageBy() {
	fmt.Println(AverageBy(func(t uint) float64 {
		return float64(t) * 0.5
	}, []uint{1, 2, 3, 4, 5}...))
	// Output: 1.5
}
