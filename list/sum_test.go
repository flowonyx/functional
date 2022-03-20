package list_test

import (
	"fmt"
	"strconv"

	"github.com/flowonyx/functional/list"
)

func ExampleSum() {
	r := list.Sum([]int{1, 2, 3})
	fmt.Println(r)
	// Output: 6
}

func ExampleSum_second() {
	r := list.Sum([]string{"hello ", "world", "!"})
	fmt.Println(r)
	// Output: hello world!
}

func ExampleSumBy() {
	type summ struct {
		val float64
	}
	r := list.SumBy(func(s summ) float64 { return s.val }, []summ{{val: 1.1}, {val: 2.2}})
	fmt.Printf("%s", strconv.FormatFloat(r, 'f', 1, 64))
	// Output: 3.3
}
