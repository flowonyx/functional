package list

import "fmt"

func ExampleEqual() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{1, 2, 3, 4}
	s4 := []int{1, 2, 3, 5, 4}

	fmt.Println(Equal(s1, s2), Equal(s1, s3), Equal(s1, s4))
	// Output: true false false
}

func ExampleEqualUnordered() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{1, 2, 3, 4}
	s4 := []int{1, 2, 3, 5, 4}

	fmt.Println(EqualUnordered(s1, s2), EqualUnordered(s1, s3), EqualUnordered(s1, s4))
	// Output: true false true
}
