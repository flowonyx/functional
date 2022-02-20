package list_test

import (
	"fmt"
	"strconv"

	"github.com/flowonyx/functional/list"
)

func ExampleIter() {
	input := []string{"0", "1", "2", "3", "4", "5"}
	var r string
	list.Iter(func(s string) { r += s }, input)
	fmt.Println(r)
	// Output: 012345
}

func ExampleIteri() {
	input := []string{"5", "4", "3", "2", "1", "0"}
	var r string
	list.Iteri(func(i int, s string) { r += strconv.Itoa(i) + s }, input)
	fmt.Println(r)
	// Output: 051423324150
}

func ExampleIter2() {
	input1 := []string{"hello", "hi", "howdy"}
	input2 := []string{"john", "mary", "tim", "larry"}
	var r string
	list.Iter2(func(t1, t2 string) { r += t1 + " " + t2 + ", " }, input1, input2)
	fmt.Println(r)
	// Output: hello john, hi mary, howdy tim,
}

func ExampleIteri2() {
	input1 := []string{"hello", "hi", "howdy"}
	input2 := []string{"john", "mary", "tim", "larry"}
	var r string
	list.Iteri2(func(i int, t1, t2 string) { r += t1 + " " + t2 + list.IfV(i == len(input1)-1, "").Else(", ") }, input1, input2)
	fmt.Println(r)
	// Output: hello john, hi mary, howdy tim
}

func ExampleIter3() {
	input1 := []string{"hello", "hi", "howdy"}
	input2 := []string{"john", "mary", "tim", "larry"}
	input3 := []string{"dunn", "contrary", "tam"}
	var r string
	list.Iter3(func(t1, t2, t3 string) { r += t1 + " " + t2 + " " + t3 + ", " }, input1, input2, input3)
	fmt.Println(r)
	// Output: hello john dunn, hi mary contrary, howdy tim tam,
}

func ExampleIteri3() {
	input1 := []string{"hello", "hi", "howdy"}
	input2 := []string{"john", "mary", "tim", "larry"}
	input3 := []string{"dunn", "contrary", "tam"}
	var r string
	list.Iteri3(func(i int, t1, t2, t3 string) {
		r += t1 + " " + t2 + " " + t3 + list.IfV(i == list.LastIndexOf(input1), "").Else(", ")
	}, input1, input2, input3)
	fmt.Println(r)
	// Output: hello john dunn, hi mary contrary, howdy tim tam
}

func ExampleIter2D() {
	input := [][]int{{1, 2}, {3, 4}, {5}}
	var r string
	list.Iter2D(func(t int) { r += strconv.Itoa(t) }, input)
	fmt.Println(r)
	// Output: 12345
}

func ExampleIteri2D() {
	input := [][]int{{1, 2}, {3, 4}, {5}}
	var r string
	list.Iteri2D(func(i int, j int, t int) { r += fmt.Sprintf("(%d, %d):%d ", i, j, t) }, input)
	fmt.Println(r)
	// Output: (0, 0):1 (0, 1):2 (1, 0):3 (1, 1):4 (2, 0):5
}

func ExampleIter3D() {
	input := [][][]int{{{1, 2}, {3, 4}}, {{5}}}
	var r string
	list.Iter3D(func(t int) { r += fmt.Sprintf("%d", t) }, input)
	fmt.Println(r)
	// Output: 12345
}

func ExampleIteri3D() {
	input := [][][]int{{{1, 2}, {3, 4}}, {{5}}}
	var r string
	list.Iteri3D(func(i int, j int, k int, t int) { r += fmt.Sprintf("(%d, %d, %d):%d ", i, j, k, t) }, input)
	fmt.Println(r)
	// Output: (0, 0, 0):1 (0, 0, 1):2 (0, 1, 0):3 (0, 1, 1):4 (1, 0, 0):5
}

func ExampleIterRev() {
	input := []int{0, 1, 2, 3, 4}
	var s string
	list.IterRev(func(t int) { s += strconv.Itoa(t) }, input)
	fmt.Println(s)
	// Output: 43210
}

func ExampleIteriRev() {
	input := []int{0, 1, 2, 3, 4}
	var s string
	list.IteriRev(func(i, t int) { s += fmt.Sprintf("%d:%d,", i, t) }, input)
	fmt.Println(s)
	// Output: 4:4,3:3,2:2,1:1,0:0,
}
