package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func mod2(i int) bool { return i%2 == 0 }
func mod6(i int) bool { return i%6 == 0 }

func ExampleFind() {
	r, err := functional.Find(mod2, []int{1, 2, 3, 4})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: 2
}

func Example2Find() {
	_, err := functional.Find(mod6, []int{1, 2, 3})
	fmt.Println(err)
	// Output: key not found
}

func ExampleFindBack() {
	r, err := functional.FindBack(mod2, []int{1, 2, 3, 4})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: 4
}

func ExampleTryFind() {
	r := functional.TryFind(mod2, []int{1, 2, 3, 4})
	if r.IsNone() {
		panic("TryFind should have found 2")
	}
	fmt.Println(r.Value())
	// Output: 2
}

func ExampleTryFindBack() {
	r := functional.TryFindBack(mod2, []int{1, 2, 3, 4})
	if r.IsNone() {
		panic("TryFind should have found 4")
	}
	fmt.Println(r.Value())
	// Output: 4
}
