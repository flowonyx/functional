package list

import "fmt"

func ExampleContains() {
	input := []int{1, 2, 3, 4, 5, 6}
	r1 := Contains(3, input...)
	r2 := Contains(7, input...)
	fmt.Println(r1, r2)
	// Output: true false
}
